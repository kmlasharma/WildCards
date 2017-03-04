package pml

type Process struct {
	Name      string
	Sequences []Sequence
}

type Sequence struct {
	Name    string
	Actions []Action
}

type Action struct {
	Name  string   `json:"-"`
	Drugs []string `json:"drugs"`
}
