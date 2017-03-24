package pml

import (
	"encoding/json"
	"errors"
	"github.com/kmlasharma/WildCards/pkgs/logger"
	"io"
	"os"
)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() *Process {
	p.ensureNextTokenType(PROCESS)
	processName := p.ensureNextTokenType(IDENT)
	p.ensureNextTokenType(LBRACE)
	baseElement := p.parseBaseElement()
	p.ensureNextTokenType(RBRACE)
	return &Process{Name: processName, baseElement: baseElement}
}

func (p *Parser) parseSequence() Sequence {
	p.ensureNextTokenType(SEQUENCE)
	sequenceName := p.ensureNextTokenType(IDENT)
	p.ensureNextTokenType(LBRACE)
	baseElement, actions := p.parseBaseElementAndActions()
	p.ensureNextTokenType(RBRACE)
	return Sequence{Name: sequenceName, baseElement: baseElement, Actions: actions}
}

func (p *Parser) parseIteration() Iteration {
	p.ensureNextTokenType(ITERATION)
	iterationName := p.ensureNextTokenType(IDENT)
	p.ensureNextTokenType(LBRACE)
	baseElement, actions := p.parseBaseElementAndActions()
	p.ensureNextTokenType(RBRACE)
	return Iteration{Name: iterationName, baseElement: baseElement, Actions: actions}
}

func (p *Parser) parseTask() Task {
	p.ensureNextTokenType(TASK)
	taskName := p.ensureNextTokenType(IDENT)
	p.ensureNextTokenType(LBRACE)
	baseElement, actions := p.parseBaseElementAndActions()
	p.ensureNextTokenType(RBRACE)
	return Task{Name: taskName, baseElement: baseElement, Actions: actions}
}

func (p *Parser) parseAction() (Action, error) {
	p.ensureNextTokenType(ACTION)
	actionName := p.ensureNextTokenType(IDENT)
	p.ensureNextTokenType(LBRACE)

	var stringifiedJSON string

	for {
		tok, _ := p.scanIgnoreWhitespace()
		if tok == SCRIPT {
			p.ensureNextTokenType(LBRACE)
			stringifiedJSON = p.ensureNextTokenType(IDENT)
			p.ensureNextTokenType(RBRACE)
		} else if tok == REQUIRES || tok == PROVIDES {
			// Wait till we reach RBRACE.
			for tok != RBRACE {
				tok, _ = p.scanIgnoreWhitespace()
			}
		} else if tok == RBRACE {
			// Wait for final RBRACE
			break
		}
	}

	if stringifiedJSON != "" {
		return decodeActionJSON(stringifiedJSON, actionName)
	}
	return Action{}, errors.New("No Script tag")
}

func (p *Parser) parseBaseElement() (baseElement BaseElement) {
	for {
		tok, _ := p.scanIgnoreWhitespace()
		p.unscan() // Put it back for cleaniness.
		if tok == SEQUENCE {
			seq := p.parseSequence()
			baseElement.Sequences = append(baseElement.Sequences, seq)
		} else if tok == ITERATION {
			iter := p.parseIteration()
			baseElement.Itertions = append(baseElement.Itertions, iter)
		} else if tok == TASK {
			task := p.parseTask()
			baseElement.Tasks = append(baseElement.Tasks, task)
		} else {
			break
		}
	}
	return
}

func (p *Parser) parseBaseElementAndActions() (baseElement BaseElement, actions []Action) {
	for {
		tok, _ := p.scanIgnoreWhitespace()
		p.unscan() // Put it back for cleaniness.
		if tok == ACTION {
			action, err := p.parseAction()
			if err == nil { // Skip if non JSON script
				actions = append(actions, action)
			}
		} else if tok == SEQUENCE {
			seq := p.parseSequence()
			baseElement.Sequences = append(baseElement.Sequences, seq)
		} else if tok == ITERATION {
			iter := p.parseIteration()
			baseElement.Itertions = append(baseElement.Itertions, iter)
		} else if tok == TASK {
			task := p.parseTask()
			baseElement.Tasks = append(baseElement.Tasks, task)
		} else {
			break
		}
	}
	return
}

func decodeActionJSON(str string, name string) (Action, error) {
	action := Action{Name: name}
	if err := json.Unmarshal([]byte(str), &action); err != nil {
		return action, errors.New("Non JSON script")
	}
	return action, nil
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

func (p *Parser) ensureNextTokenType(tok Token) string {
	token, lit := p.scanIgnoreWhitespace()
	if tok != token {
		logger.Error("found '", lit, "', expected ", tok)
		os.Exit(1)
	}
	return lit
}
