package pml

type BaseElement struct {
	Sequences []Sequence
	Itertions []Iteration
	Tasks     []Task
}

type Process struct {
	Name        string
	baseElement BaseElement
	Actions     []Action
}

type Sequence struct {
	Name        string
	baseElement BaseElement
	Actions     []Action
}

type Iteration struct {
	Name        string
	baseElement BaseElement
	Actions     []Action
}

type Task struct {
	Name        string
	baseElement BaseElement
	Actions     []Action
}

type Action struct {
	Name  string   `json:"-"`
	Drugs []string `json:"drugs"`
}

func (p Process) AllTasks() []Task {
	return p.baseElement.allTasks()
}

func (p Process) AllDrugs() (drugs []string) {
	return p.baseElement.allDrugs()
}

/* Ensure there is no name clashes */
func (p Process) Validate() error {
	return nil
}

/* Helper Functions */

func (baseElement BaseElement) allTasks() (tasks []Task) {
	tasks = append(tasks, baseElement.Tasks...)
	for _, iteration := range baseElement.Itertions {
		tasks = append(tasks, iteration.baseElement.allTasks()...)
	}
	for _, seq := range baseElement.Sequences {
		tasks = append(tasks, seq.baseElement.allTasks()...)
	}
	for _, task := range baseElement.Tasks {
		tasks = append(tasks, task.baseElement.allTasks()...)
	}
	return
}

func (baseElement BaseElement) allDrugs() (drugs []string) {
	for _, iteration := range baseElement.Itertions {
		drugs = append(drugs, iteration.allDrugs()...)
	}
	for _, seq := range baseElement.Sequences {
		drugs = append(drugs, seq.allDrugs()...)
	}
	for _, task := range baseElement.Tasks {
		drugs = append(drugs, task.allDrugs()...)
	}
	return
}

func (iteration Iteration) allDrugs() (drugs []string) {
	for _, action := range iteration.Actions {
		drugs = append(drugs, action.Drugs...)
	}
	drugs = append(drugs, iteration.baseElement.allDrugs()...)
	return
}

func (seq Sequence) allDrugs() (drugs []string) {
	for _, action := range seq.Actions {
		drugs = append(drugs, action.Drugs...)
	}
	drugs = append(drugs, seq.baseElement.allDrugs()...)
	return
}

func (task Task) allDrugs() (drugs []string) {
	for _, action := range task.Actions {
		drugs = append(drugs, action.Drugs...)
	}
	drugs = append(drugs, task.baseElement.allDrugs()...)
	return
}
