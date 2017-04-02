package pml

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
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

func TestDelayExistence(t * testing.T) {
        fmt.Println("* Testing that delays are processed")
        process, err := processFromFile("delays.pml")
        success, message := delayHelper(process)
        if (assert.Equal(t, err, nil, "There should no errors detected processing the file") && assert.Equal(t, success, true, message)) {
                fmt.Println("PASSED!")
        }
}

func TestActionDelayFail(t * testing.T) {
        fmt.Println("* Testing that PML files with delays inside actions are rejected")
        _, err := processFromFile("action_delay.pml")
        if assert.NotEqual(t, err, nil, "There should be an error raised due to delay inside action") {
                fmt.Println("PASSED!")
        }
}

func TestTimeIntervalOffset(t * testing.T) {
        fmt.Println("* Testing that time interval offsets are processed")
        process, err := processFromFile("time_interval_offset.pml")
	success, message := timeIntervalOffsetHelper(process)
        if (assert.Equal(t, err, nil, "There should no errors detected processing the file") && assert.Equal(t, success, true, message)) {
		fmt.Println("PASSED!")
	}
}


func TestPeriodicDrugUse(t * testing.T) {
	fmt.Println("* Testing that periodic drug use is registered")
	process, err := processFromFile("periodic_use.pml")
	success, message := periodicDrugUseHelper(process)
	if (assert.NotEqual(t, err, nil, "There should not be an error") && assert.Equal(t, success, true, message)) {
		fmt.Println("PASSED!")
	}
}

func TestSequentialDDI(t * testing.T) {
	fmt.Println("* Testing that periodic drug use is registered")
	process, err := processFromFile("sequential_ddi.pml")
	success, message := sequentialDDIHelper(process)
	if (assert.NotEqual(t, err, nil, "There should not be an error") && assert.Equal(t, success, true, message)) {
		fmt.Println("PASSED!")
	}
}

func sequentialDDIHelper(process *Element) (success bool, message string) {
	drugPair := process.FindDrugPairs[0]

	expectedDrugA := "coke"
	expectedDrugB := "pepsi"
	expectedDelay := Delay("0")
	expectedDDIType := SequentialType
	expectedParentName := "seq1"

	actualDrugA := drugPair.DrugA
	actualDrugB := drugPair.DrugB
	actualDelay := drugPair.Delay
	actualDDIType := drugPair.ddiType
	actualParentName := drugPair.parentName

	if expectedDrugA != actualDrugA {
		return false, "Expected DrugA {" + expectedDrugA + "} does not match actual DrugA {" + actualDrugA + "}"
	}
	if expectedDrugB != actualDrugB {
		return false, "Expected DrugB {" + expectedDrugB + "} does not match actual DrugB {" + actualDrugB + "}"
	}
	if expectedDelay != actualDelay {
		return false, "Expected delay {" + expectedDelay + "} does not match actual delay {" + actualDelay + "}"
	}
	if expectedDDIType != actualDDIType {
		return false, "Expected DDI type {" + expectedDDIType.String() + "} does not match actual DDI {" + actualDDIType.String() + "}"
	}
	if expectedParentName != actualParentName {
		return false, "Expected parent name {" + expectedParentName + "} does not match actual parent name {" + actualParentName + "}"
	}
	return true, "success"
}

func periodicDrugUseHelper(process *Element) (success bool, message string) {
	iter1 := process.Children[0].(Element)

	expected_loops := Loops(5)
	expected_delay := Delay("3 days")

	actual_loops := iter1.Children[0].(Loops)
	actual_delay := iter1.Children[2].(Delay)

	if expected_loops != actual_loops {
		return false, "Expected loops {" + expected_loops + "} does not equal actual loops {" + actual_loops + "}"
	}
	if expected_delay != actual_delay {
		return false, "Expected delay {" + expected_delay + "} does not equal actual delay {" + actual_delay + "}"
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
	proc_children := process.Children

	seq1 := proc_children[0].(Element)
	task1 := proc_children[1].(Element)
	iter1 := proc_children[4].(Element)

	seq1_children := seq1.Children
	seq1_expected_delay := Delay("30 sec")
	seq1_actual_delay := seq1_children[1].(Delay)
	if seq1_expected_delay != seq1_actual_delay {
		return false, "seq1 expected does not match actual"
	}
	task1_children := task1.Children
	task1_expected_delay := Delay("20 min")
	task1_actual_delay := task1_children[1].(Delay)
	if task1_expected_delay != task1_actual_delay {
		return false, "task1 expected does not match actual"
	}
	proc_delay_one_actual := proc_children[2].(Element).(Delay)
	proc_delay_one_expected := Delay("5 hr")
	if proc_delay_one_expected != proc_delay_one_actual {
		return false, "proc delay one expected does not match actual"
	}
	proc_delay_two_actual := proc_children[3].(Element).(Delay)
	proc_delay_two_expected := Delay("4 day")
	if proc_delay_two_expected != proc_delay_two_actual {
		return false, "proc delay two expected does not match actual"
	}
	iter1_children := iter1.Children
	iter1_expected_delay := Delay("3 week")
	iter1_actual_delay := iter1_children[1].(Delay)
	if iter1_expected_delay != iter1_actual_delay {
		return false, "iter1 expected does not match actual"
	}
	return true, ""
}


func timeIntervalOffsetHelper(process *Element) (success bool, message string) {
	proc_children := process.Children
	timeIntervalOffset := proc_children[0].(Element).(Wait)
	subsequentDelay := proc_children[1].(Element).(Delay)
	currentDateAndTime := time.Now().Format(time.UnixDate)
	today := strings.Split(currentDateAndTime, " ")[0]

	var waitLength = Delay("0 day")

	switch today {
	case "Mon":
		waitLength = Delay("0 day")
	case "Tue":
		waitLength = Delay("1 day")
	case "Wed":
		waitLength = Delay("2 day")
	case "Thu":
		waitLength = Delay("3 day")
	case "Fri":
		waitLength = Delay("4 day")
	case "Sat":
		waitLength = Delay("5 day")
	case "Sun":
		waitLength = Delay("6 day")
	}

	if subsequentDelay != waitLength {
		return false, "Did not delay for the right amount of time based on specified time interval offset"
	}
	return true, ""
}


// TODO: tests for: 	Merging clinical pathways
//			PML-TX save to file
// 			Time interval offset

