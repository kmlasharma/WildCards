package pml

// Token represents a lexical token.
type Token string

const (
	// Special tokens
	ILLEGAL Token = "Illegal"
	EOF           = "EOF"
	WS            = "Whitespace"

	// Literals
	IDENT = "Ident" // names

	// Misc characters
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	PROCESS  = "process"
	SEQUENCE = "sequence"
	ACTION   = "action"
	SCRIPT   = "script"
)
