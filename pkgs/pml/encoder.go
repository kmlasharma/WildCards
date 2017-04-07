package pml

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func (el Element) Encode(indent string) string {
	initialLine := fmt.Sprintf("%s%s %s {", indent, elementTypeToToken(el.elementType), el.Name)
	lastLine := fmt.Sprintf("%s}", indent)
	lines := []string{initialLine}
	for _, child := range el.Children {
		lines = append(lines, child.Encode(indent+"  "))
	}

	if el.Loops > 0 {
		line := fmt.Sprintf("%sloops { \"%d\" }", indent+"  ", el.Loops)
		lines = append(lines, line)
	}

	lines = append(lines, lastLine)
	return strings.Join(lines, "\n")
}

func (act Action) Encode(indent string) string {
	initialLine := fmt.Sprintf("%saction %s {", indent, act.Name)
	scriptLine := fmt.Sprintf("%s  script { %s }", indent, encodeDrugs(act.Drugs))
	lastLine := fmt.Sprintf("%s}", indent)
	lines := []string{initialLine, scriptLine, lastLine}
	return strings.Join(lines, "\n")
}

func (d Delay) Encode(indent string) string {
	return fmt.Sprintf("%sdelay { \"%s\" }", indent, d.toHumanReadableDate())
}

func (w Wait) Encode(indent string) string {
	return fmt.Sprintf("%swait { \"%s\" }", indent, w)
}

func elementTypeToToken(et ElementType) string {
	switch et {
	case ProcessType:
		return "process"
	case IterationType:
		return "iteration"
	case TaskType:
		return "task"
	case BranchType:
		return "branch"
	case SelectionType:
		return "selection"
	case SequenceType:
		return "sequence"
	case ActionType:
		return "action"
	case DelayType:
		return "delay"
	default:
		return ""
	}
}

func encodeDrugs(drugs []string) string {
	dict := map[string]interface{}{"drugs": drugs}
	bytes, _ := json.Marshal(dict)
	return strconv.Quote(string(bytes))
}
