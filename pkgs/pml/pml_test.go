package pml

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var resDir = os.Getenv("RES_DIR")

func TestMissingFile(t *testing.T) {
	fmt.Println("* Testing loading PML file that does not exist")
	_, err := processFromFile("nonexistent.pml")
	if assert.NotNil(t, err, "Process should not exist") {
		fmt.Println("PASSED!")
	}
}

func TestNoProcesses(t *testing.T) {
	fmt.Println("* Testing loading PML file with no processes")
	_, err := processFromFile("no_processes.pml")
	if assert.NotNil(t, err, "Process should not exist") {
		fmt.Println("PASSED!")
	}
}

func TestMultipleProcesses(t *testing.T) {
	fmt.Println("* Testing loading PML file with multiple processes")
	_, err := processFromFile("multiple_processes.pml")
	if assert.NotNil(t, err, "Process should not exist") {
		fmt.Println("PASSED!")
	}
}

func TestMissingPMLConstruct(t *testing.T) {
	fmt.Println("* Testing Missing PML Construct")
	_, err := processFromFile("missing_pml_construct.pml")
	if assert.NotNil(t, err, "Process should fail due to missing PML construct") {
		fmt.Println("PASSED!")
	}
}

func TestSubtasksExist(t *testing.T) {
	fmt.Println("* Testing process with sequences")
	process, _ := processFromFile("subtasks.pml")
	tasks := process.AllTasks()
	if assert.Equal(t, len(tasks), 2, "Process should have two subtasks") {
		fmt.Println("PASSED!")
	}
}

func TestNoSubtasks(t *testing.T) {
	fmt.Println("* Testing analysing process with no sequences")
	process, _ := processFromFile("no_subtasks.pml")
	tasks := process.AllTasks()
	if assert.Equal(t, len(tasks), 0, "Process should have no subtasks") {
		fmt.Println("PASSED!")
	}
}

func TestNoDrugs(t *testing.T) {
	fmt.Println("* Testing analysing process with a script with no drugs")
	process, _ := processFromFile("no_drugs.pml")
	drugs := process.AllDrugs()
	if assert.Equal(t, len(drugs), 0, "Process should have no drugs") {
		fmt.Println("PASSED!")
	}
}

func TestMultipleDrugs(t *testing.T) {
	fmt.Println("* Testing analysing process with multiple sequences and actions containing drugs")
	process, _ := processFromFile("multi_drugs.pml")
	drugs := process.AllDrugs()
	expected := []string{"coke", "7up", "pepsi", "fanta", "dr pepper"}
	if assert.Equal(t, drugs, expected, "Drugs should include exactly coke, 7up, pepsi, fanta and dr pepper") {
		fmt.Println("PASSED!")
	}
}

func TestValidateSameType(t *testing.T) {
	fmt.Println("* Testing that PML files with name clashes among items of the same type are rejected")
	_, err := processFromFile("sequence_clashes.pml")
	if assert.NotNil(t, err, "There should be name clashes that are detected") {
		fmt.Println("PASSED!")
	}
}

func TestValidateDifferentTypes(t *testing.T) {
	fmt.Println("* Testing that PML files with name clashes in different namespaces are accepted")
	_, err := processFromFile("multilevel_clashes.pml")
	if assert.NotNil(t, err, "There should be name clashes that are detected") {
		fmt.Println("PASSED!")
	}
}

func TestValidateNoClashes(t *testing.T) {
	fmt.Println("* Testing that PML files with no name clashes are not rejected")
	_, err := processFromFile("no_clashes.pml")
	if assert.Nil(t, err, "There should be no name clashes detected") {
		fmt.Println("PASSED!")
	}
}

func TestBrokenFile(t *testing.T) {
	fmt.Println("* Testing that malformed PML files are rejected")
	_, err := processFromFile("errortest.pml")
	if assert.NotNil(t, err, "An issue should be raised to syntax errors in the PML") {
		fmt.Println("PASSED!")
	}
}

func TestDelayExistence(t *testing.T) {
	fmt.Println("* Testing that malformed PML files are rejected")
	process, err := processFromFile("delays.pml")
	success, message := delayHelper(process)
	if assert.Equal(t, err, nil, "There should no errors detected processing the file") && assert.Equal(t, success, true, message) {
		fmt.Println("PASSED!")
	}
}

func TestActionDelayFail(t *testing.T) {
	fmt.Println("* Testing that PML files with delays inside actions are rejected")
	_, err := processFromFile("action_delay.pml")
	if assert.NotEqual(t, err, nil, "There should be an error raised due to delay inside action") {
		fmt.Println("PASSED!")
	}
}

func TestPeriodicDrugUse(t *testing.T) {
	fmt.Println("* Testing that periodic drug use is registered")
	process, err := processFromFile("periodic_use.pml")
	success, message := periodicDrugUseHelper(process)
	if assert.NotEqual(t, err, nil, "There should not be an error") && assert.Equal(t, success, true, message) {
		fmt.Println("PASSED!")
	}
}

func TestSequentialDDI(t *testing.T) {
	fmt.Println("* Testing that periodic drug use is registered")
	process, err := processFromFile("sequential_ddi.pml")
	success, message := sequentialDDIHelper(process)
	if assert.NotEqual(t, err, nil, "There should not be an error") && assert.Equal(t, success, true, message) {
		fmt.Println("PASSED!")
	}
}

func sequentialDDIHelper(process *Element) (success bool, message string) {
	drugPair := process.FindDrugPairs()[0]

	expectedDrugA := "coke"
	expectedDrugB := "pepsi"
	expectedDelay := Delay(0)
	expectedDDIType := SequentialType
	expectedParentName := "seq1"

	actualDrugA := drugPair.DrugA
	actualDrugB := drugPair.DrugB
	actualDelay := drugPair.delay
	actualDDIType := drugPair.ddiType
	actualParentName := drugPair.parentName

	if expectedDrugA != actualDrugA {
		return false, fmt.Sprintf("Expected DrugA { %s } does not match actual DrugA { %s }", expectedDrugA, actualDrugA)
	}
	if expectedDrugB != actualDrugB {
		return false, fmt.Sprintf("Expected DrugB { %s } does not match actual DrugB { %s }", expectedDrugB, actualDrugB)
	}
	if expectedDelay != actualDelay {
		return false, fmt.Sprintf("Expected delay { %i } does not match actual delay { %i }", expectedDelay, actualDelay)
	}
	if expectedDDIType != actualDDIType {
		return false, fmt.Sprintf("Expected DDIType { %s } does not match actual type { %s }", expectedDDIType, actualDDIType)
	}
	if expectedParentName != actualParentName {
		return false, fmt.Sprintf("Expected parent name { %s } does not match actual name { %s }", expectedParentName, actualParentName)
	}
	return true, "success"
}

func periodicDrugUseHelper(process *Element) (success bool, message string) {
	iter1 := process.Children[0].(Element)

	expectedLoops := 5
	expectedDelay := NewDelay("3 days")

	actualLoops := (iter1.Children[0]).(Element).Loops
	actualDelay := iter1.Children[2].(Delay)

	if expectedLoops != actualLoops {
		return false, fmt.Sprintf("Expected loops { %i } does not match actual loops { %i }", expectedLoops, actualLoops)
	}
	if expectedDelay != actualDelay {
		return false, fmt.Sprintf("Expected delay { %i } does not match actual delay { %i }", expectedDelay, actualDelay)
	}
	return true, "success"
}

func processFromFile(filepath string) (*Element, error) {
	reader, err := os.Open(resDir + "/" + filepath)
	if err != nil {
		return &Element{}, err
	}
	parser := NewParser(reader)
	process, err := parser.Parse()
	return process, err
}

func delayHelper(process *Element) (success bool, message string) {
	procChildren := process.Children

	seq1 := procChildren[0].(Element)
	task1 := procChildren[1].(Element)
	iter1 := procChildren[4].(Element)

	seq1Children := seq1.Children
	seq1ExpectedDelay := Delay("30 sec")
	seq1ActualDelay := seq1Children[1].(Delay)
	if seq1ExpectedDelay != seq1ActualDelay {
		return false, "seq1 expected does not match actual"
	}
	task1Children := task1.Children
	task1ExpectedDelay := Delay("20 min")
	task1ActualDelay := task1Children[1].(Delay)
	if task1ExpectedDelay != task1ActualDelay {
		return false, "task1 expected does not match actual"
	}
	procDelayOneActual := procChildren[2].(Delay)
	procDelayOneExpected := Delay("5 hr")
	if procDelayOneExpected != procDelayOneActual {
		return false, "proc delay one expected does not match actual"
	}
	procDelayTwoActual := procChildren[3].(Element).(Delay)
	procDelayTwoExpected := Delay("4 day")
	if procDelayTwoExpected != procDelayTwoActual {
		return false, "proc delay two expected does not match actual"
	}
	iter1Children := iter1.Children
	iter1ExpectedDelay := Delay("3 week")
	iter1ActualDelay := iter1Children[1].(Delay)
	if iter1ExpectedDelay != iter1ActualDelay {
		return false, "iter1 expected does not match actual"
	}
	return true, ""
}

// TODO: tests for: 	Merging clinical pathways
//			PML-TX save to file
// 			Time interval offset
