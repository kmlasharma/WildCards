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

func (p Process) AllDrugs() (drugs []string) {
	for _, seq := range p.Sequences {
		for _, action := range seq.Actions {
			for _, drug := range action.Drugs {
				drugs = append(drugs, drug)
			}
		}
	}
	return
}