package pml

import (
	"encoding/json"
	"fmt"
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
		fmt.Println("found ", lit, ", expected ", token)
		os.Exit(1)
	}
	return lit
}

func (p *Parser) Parse() *Process {
	actions := []Action{}

	p.ensureNextTokenType(PROCESS)
	processName := p.ensureNextTokenType(IDENT)
	p.ensureNextTokenType(LBRACE)
	p.ensureNextTokenType(SEQUENCE)
	sequenceName := p.ensureNextTokenType(IDENT)
	p.ensureNextTokenType(LBRACE)

	for {
		if tok, _ := p.scanIgnoreWhitespace(); tok == ACTION {
			actionName := p.ensureNextTokenType(IDENT)
			p.ensureNextTokenType(LBRACE)
			p.ensureNextTokenType(SCRIPT)
			p.ensureNextTokenType(LBRACE)
			stringifiedJSON := p.ensureNextTokenType(IDENT)
			p.ensureNextTokenType(RBRACE)
			p.ensureNextTokenType(RBRACE)

			action := decodeActionJSON(stringifiedJSON)
			action.Name = actionName
			actions = append(actions, action)
		} else {
			p.unscan() // Put final one back so it's used below.
			break
		}
	}
	p.ensureNextTokenType(RBRACE)
	p.ensureNextTokenType(RBRACE)
	sequence := Sequence{Name: sequenceName, Actions: actions}
	return &Process{Name: processName, Sequences: []Sequence{sequence}}
}

func decodeActionJSON(str string) Action {
	fmt.Println(str)
	action := Action{}
	if err := json.Unmarshal([]byte(str), &action); err != nil {
		panic(err)
	}
	return action
}
