package pml

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type errParser struct {
	p   *Parser
	err error
}

func (ep *errParser) expect(tok Token) string {
	if ep.err != nil {
		return ""
	}
	lit, err := ep.p.ensureNextTokenType(tok)

	if err != nil && tok == IDENT {
		str := fmt.Sprintf("Un-named PML construct on line %d", ep.p.currentLineNumber())
		ep.err = errors.New(str)
	} else {
		ep.err = err
	}

	// Check if there is a name clash
	if tok == IDENT {
		if ep.p.freq[lit] {
			str := fmt.Sprintf("Name clash found: %s on line %d", lit, ep.p.currentLineNumber())
			ep.err = errors.New(str)
		}
		ep.p.freq[lit] = true
	}

	return lit
}

type Parser struct {
	s    *Scanner
	freq map[string]bool
	buf  struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r), freq: make(map[string]bool)}
}

func (p *Parser) Parse() (*Element, error) {
	ep := &errParser{p: p}
	ep.expect(PROCESS)
	processName := ep.expect(IDENT)
	ep.expect(LBRACE)
	element := p.parseChildren(ep)
	ep.expect(RBRACE)
	ep.expect(EOF) // There should be nothing else in the file other than the process.

	if ep.err != nil {
		return &Element{}, ep.err
	}

	element.Name = processName
	element.elementType = ProcessType
	return &element, nil
}

func (p *Parser) parseChildren(ep *errParser) (element Element) {
	if ep.err != nil {
		return Element{} // If there's already an error, don't bother doing all this.
	}
	for {
		tok, _ := p.scanIgnoreWhitespace()
		p.unscan() // Put it back for cleaniness.
		if tok == ACTION {
			action, err := p.parseAction(ep)
			if err == nil { // Skip if non JSON script
				element.Children = append(element.Children, action)
			}
		} else if tok == SEQUENCE || tok == TASK || tok == ITERATION || tok == BRANCH || tok == SELECTION {
			ele := p.parseElement(tok, ep)
			ele.elementType = elementTypeForToken(tok)
			element.Children = append(element.Children, ele)
		} else if tok == DELAY {
			delay := p.parseDelay(ep)
			element.Children = append(element.Children, delay)
		} else if tok == LOOPS {
			element.Loops = p.parseLoops(ep)
		} else if tok == WAIT {
			offset := p.parseTimeIntervalOffset(ep)
			element.Children = append(element.Children, offset)
		} else {
			break
		}
	}
	return
}

func (p *Parser) parseElement(initialToken Token, ep *errParser) *Element {
	ep.expect(initialToken)
	name := ep.expect(IDENT)
	ep.expect(LBRACE)
	element := p.parseChildren(ep)
	ep.expect(RBRACE)
	element.Name = name
	return &element
}

func (p *Parser) parseAction(ep *errParser) (Action, error) {
	ep.expect(ACTION)
	actionName := ep.expect(IDENT)
	ep.expect(LBRACE)

	var stringifiedJSON string

	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == SCRIPT {
			ep.expect(LBRACE)
			stringifiedJSON = ep.expect(LIT)
			ep.expect(RBRACE)
		} else if tok == REQUIRES || tok == PROVIDES {
			// Wait till we reach RBRACE.
			for tok != RBRACE {
				tok, _ = p.scanIgnoreWhitespace()
			}
		} else if tok == RBRACE {
			break // Wait for final RBRACE
		}
	}

	if stringifiedJSON != "" {
		return decodeActionJSON(stringifiedJSON, actionName)
	}
	return Action{}, errors.New("No Script tag")
}

func (p *Parser) parseDelay(ep *errParser) Delay {
	ep.expect(DELAY)
	ep.expect(LBRACE)
	str := ep.expect(LIT)
	ep.expect(RBRACE)
	return NewDelay(str)
}

func (p *Parser) parseLoops(ep *errParser) int {
	ep.expect(LOOPS)
	ep.expect(LBRACE)
	str := ep.expect(LIT)
	ep.expect(RBRACE)

	i, _ := strconv.Atoi(str)
	return i
}

func (p *Parser) parseTimeIntervalOffset(ep *errParser) Wait {
	ep.expect(WAIT)
	ep.expect(LBRACE)
	timeIntervalOffset := ep.expect(LIT)
	ep.expect(RBRACE)
	return Wait(timeIntervalOffset)
}

func decodeActionJSON(str string, name string) (Action, error) {
	action := Action{Name: name}
	if err := json.Unmarshal([]byte(str), &action); err != nil {
		return action, errors.New("Non JSON script")
	}
	return action, nil
}

func elementTypeForToken(tok Token) ElementType {
	switch tok {
	case SEQUENCE:
		return SequenceType
	case ITERATION:
		return IterationType
	case TASK:
		return TaskType
	case BRANCH:
		return BranchType
	case SELECTION:
		return SelectionType
	default:
		return ProcessType
	}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer (by forcing it to return it again next time)
func (p *Parser) unscan() {
	p.buf.n = 1
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) ensureNextTokenType(tok Token) (string, error) {
	if tok == LIT {
		tok = IDENT // LIT and IDENT are treated the same here
	}
	token, lit := p.scanIgnoreWhitespace()
	if tok != token {
		str := fmt.Sprintf("found '%s', expected %s on line %d", lit, tok, p.currentLineNumber())
		return "", errors.New(str)
	}
	return lit, nil
}

func (p *Parser) currentLineNumber() int {
	return p.s.ln
}
