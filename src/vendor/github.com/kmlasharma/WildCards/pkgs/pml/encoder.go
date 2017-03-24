package pml

// import (
// 	"encoding/json"
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func (p *Process) Encode() string {
// 	return fmt.Sprintf("process %s {\n\t%s\n}", p.Name, encodeSequences(p.Sequences))
// }

// func encodeSequences(sequences []Sequence) string {
// 	arr := []string{}
// 	for _, seq := range sequences {
// 		arr = append(arr, encodeSequence(seq))
// 	}
// 	return strings.Join(arr, "\n")
// }

// func encodeSequence(seq Sequence) string {
// 	return fmt.Sprintf("sequence %s {\n\t%s\n}", seq.Name, encodeActions(seq.Actions))
// }

// func encodeIterations(iterations []Iteration) string {
// 	arr := []string{}
// 	for _, iter := range iterations {
// 		arr = append(arr, encodeIteration(iter))
// 	}
// 	return strings.Join(arr, "\n")
// }

// func encodeIteration(iter Iteration) string {
// 	return fmt.Sprintf("iteration %s {\n\t%s\n}", iter.Name, encodeActions(iter.Actions))
// }

// func encodeActions(actions []Action) string {
// 	arr := []string{}
// 	for _, action := range actions {
// 		arr = append(arr, encodeAction(action))
// 	}
// 	return strings.Join(arr, "\n")
// }

// func encodeAction(action Action) string {
// 	return fmt.Sprintf("action %s {\n\t%s\n}", action.Name, encodeScript(action.Drugs))
// }

// func encodeScript(drugs []string) string {
// 	dict := map[string]interface{}{"drugs": drugs}
// 	str, _ := json.Marshal(dict)
// 	return fmt.Sprintf("script { %s }", strconv.Quote(string(str)))
// }
