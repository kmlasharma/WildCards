package pml

type ElementType int

const (
	ProcessType ElementType = iota
	IterationType
	TaskType
	BranchType
	SelectionType
	SequenceType
	ActionType
	DelayType
)

type DrugPair struct {
	DrugA      string
	DrugB      string
	delay      Delay
	ddiType    DDIType //parallel, sequential etc
	parentName string
}

// temp getter. TODO remove
func (dp DrugPair) ParentName() string {
	return dp.parentName
}

type DDIType int

const (
	SequentialType DDIType = iota
	ParallelType
	RepeatedAlternativeDDIType
	AlternativeNonDDIType
)

func (d DDIType) String() string {
	switch d {
	case SequentialType:
		return "Sequential Type"
	case ParallelType:
		return "Parallel Type"
	case RepeatedAlternativeDDIType:
		return "Repeated Alternative DDI Type"
	case AlternativeNonDDIType:
		return "Alternative Non-DDI Type"
	default:
		return ""
	}
}

type ElementInterface interface {
	Type() ElementType
	IsSubElementType() bool
	GetName() string
	ChangeNames(string)
	Encode(string) string
}

type Element struct {
	Name        string
	Children    []ElementInterface
	Loops       int /* Only applies to iterations */
	elementType ElementType
}

type Action struct {
	Name  string   `json:"-"`
	Drugs []string `json:"drugs"`
}

type Delay int

func NewDelay(delayStr string) Delay {
	seconds := convertToSeconds(delayStr)
	return Delay(seconds)
}

func (el Element) Type() ElementType {
	return el.elementType
}

func (act Action) Type() ElementType {
	return ActionType
}

func (delay Delay) Type() ElementType {
	return DelayType
}

func (el Element) GetName() string {
	return el.Name
}

func (act Action) GetName() string {
	return act.Name
}

func (dl Delay) GetName() string {
	return "delay"
}

func (el Element) IsSubElementType() bool {
	return true
}

func (act Action) IsSubElementType() bool {
	return false
}

func (dl Delay) IsSubElementType() bool {
	return false
}

func (el Element) AllDrugs() (drugs []string) {
	for _, child := range el.Children {
		if child.Type() == ActionType {
			action := child.(*Action)
			drugs = append(drugs, action.Drugs...)
		} else if child.IsSubElementType() {
			element := child.(*Element)
			drugs = append(drugs, element.AllDrugs()...)
		}
	}
	return
}

// Find all tasks for this element, and also find subtasks of subelements
func (el Element) AllTasks() (tasks []*Element) {
	for _, child := range el.Children {
		if child.Type() == TaskType {
			element := child.(*Element)
			tasks = append(tasks, element)
		}
		if child.IsSubElementType() {
			element := child.(*Element)
			tasks = append(tasks, element.AllTasks()...)
		}
	}
	return
}

func (d Delay) toHumanReadableDate() string {
	return "30 secs" // TODO
}
