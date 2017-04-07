package pml

// Token represents a lexical token.
type Token string

const (
	// Special tokens
	ILLEGAL Token = "Illegal"
	EOF           = "EOF"
	WS            = "Whitespace"

	// Literals
	IDENT = "Ident"   // names
	LIT   = "Literal" // Literal

	// Misc characters
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	PROCESS   = "process"
	SEQUENCE  = "sequence"
	ITERATION = "iteration"
	TASK      = "task"
	BRANCH    = "branch"
	SELECTION = "selection"
	DELAY     = "delay"
	LOOPS     = "loops"
	WAIT      = "wait"
	ACTION    = "action"
	SCRIPT    = "script"
	REQUIRES  = "requires"
	PROVIDES  = "provides"
)