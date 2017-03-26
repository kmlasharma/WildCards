package pml

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Scanner struct {
	r  *bufio.Reader
	ln int /* current line number */
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), ln: 0}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	if ch == '\n' {
		s.ln++
	}

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) || isQuote(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '{':
		return LBRACE, string(ch)
	case '}':
		return RBRACE, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer

	// We are in a quote initially if the first char is a quotation mark.
	currentChar := s.read()
	inQuote := currentChar == '"'
	lastCharIsEscape := false

	// If the first character isn't a quotation mark, add it to buffer.
	if !isQuote(currentChar) {
		buf.WriteRune(currentChar)
	}

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.

	// If We are in a quotation mark string, allow anything
	for {
		ch := s.read()
		if ch == eof {
			break
		} else if inQuote || isLetter(ch) || isDigit(ch) || isQuote(ch) || ch == '_' {
			if isQuote(ch) && !lastCharIsEscape {
				inQuote = !inQuote
			} else if ch != escapeChar {
				// add to string if not opening/closing quotes, or escape character
				buf.WriteRune(ch)
			}
		} else {
			s.unread()
			break
		}
		lastCharIsEscape = ch == escapeChar
	}

	// If the string matches a keyword then return that keyword. Ignore case.
	switch strings.ToUpper(buf.String()) {
	case "PROCESS":
		return PROCESS, buf.String()
	case "SEQUENCE":
		return SEQUENCE, buf.String()
	case "ITERATION":
		return ITERATION, buf.String()
	case "TASK":
		return TASK, buf.String()
	case "ACTION":
		return ACTION, buf.String()
	case "SCRIPT":
		return SCRIPT, buf.String()
	case "REQUIRES":
		return REQUIRES, buf.String()
	case "PROVIDES":
		return PROVIDES, buf.String()
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isQuote(ch rune) bool {
	return ch == '"'
}

// eof represents a marker rune for the end of the reader.
var eof = rune(0)

var escapeChar = rune('\\')
