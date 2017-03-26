package pml

import (
	"errors"
	"fmt"
)

type Element struct {
	Name       string
	Sequences  []*Element
	Iterations []*Element
	Tasks      []*Element
	Actions    []Action
}

type Action struct {
	Name  string   `json:"-"`
	Drugs []string `json:"drugs"`
}

/* Ensure there is no name clashes */
/*
	* Check:
		- p.Actions
		- p.baseElement.Sequences
		- p.baseElement.Iterations
		- p.baseElement.Tasks
		for seq in p.baseElement.Sequences
			- seq.Validate()
		etc.
*/
func (el Element) Validate() (errs []error) {
	errs = append(errs, validateActions(el.Actions)...)
	errs = append(errs, validateElements(el.Sequences)...)
	errs = append(errs, validateElements(el.Iterations)...)
	errs = append(errs, validateElements(el.Tasks)...)

	for _, iteration := range el.Iterations {
		errs = append(errs, iteration.Validate()...)
	}
	for _, seq := range el.Sequences {
		errs = append(errs, seq.Validate()...)
	}
	for _, task := range el.Tasks {
		errs = append(errs, task.Validate()...)
	}
	return
}

func (el Element) AllDrugs() (drugs []string) {
	for _, action := range el.Actions {
		drugs = append(drugs, action.Drugs...)
	}
	for _, iteration := range el.Iterations {
		drugs = append(drugs, iteration.AllDrugs()...)
	}
	for _, seq := range el.Sequences {
		drugs = append(drugs, seq.AllDrugs()...)
	}
	for _, task := range el.Tasks {
		drugs = append(drugs, task.AllDrugs()...)
	}
	return
}

func (el Element) AllTasks() (tasks []*Element) {
	tasks = append(tasks, el.Tasks...)
	for _, iteration := range el.Iterations {
		tasks = append(tasks, iteration.AllTasks()...)
	}
	for _, seq := range el.Sequences {
		tasks = append(tasks, seq.AllTasks()...)
	}
	for _, task := range el.Tasks {
		tasks = append(tasks, task.AllTasks()...)
	}
	return
}

func validateActions(actions []Action) (errs []error) {
	freq := make(map[string]bool)
	for _, action := range actions {
		if freq[action.Name] {
			str := fmt.Sprintf("Name clash found: %s", action.Name)
			errs = append(errs, errors.New(str))
		}
		freq[action.Name] = true
	}
	return
}

func validateElements(elements []*Element) (errs []error) {
	freq := make(map[string]bool)
	for _, element := range elements {
		if freq[element.Name] {
			str := fmt.Sprintf("Name clash found: %s", element.Name)
			errs = append(errs, errors.New(str))
		}
		freq[element.Name] = true
	}
	return
}
