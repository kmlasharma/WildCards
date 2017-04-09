package pml

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDDIPairsListsEqual(t *testing.T) {
	fmt.Println("Comparing two DrugPair lists")
	var listOne = []DrugPair{
		DrugPair{
			DrugA:      "coke",
			DrugB:      "pepsi",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
		DrugPair{
			DrugA:      "7up",
			DrugB:      "club",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
		DrugPair{
			DrugA:      "fizz",
			DrugB:      "lemsip",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
	}
	var listTwo = []DrugPair{
		DrugPair{
			DrugA:      "fizz",
			DrugB:      "lemsip",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
		DrugPair{
			DrugA:      "coke",
			DrugB:      "pepsi",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
		DrugPair{
			DrugA:      "7up",
			DrugB:      "club",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
	}
	if assert.True(t, drugPairListsEqual(listOne, listTwo), "DrugPair lists should be the same") {
		fmt.Println("PASSED!")
	}
}

func TestDrugPairListContains(t *testing.T) {
	fmt.Println("Checking for DrugPair in DrugPair list")
	var listOne = []DrugPair{
		DrugPair{
			DrugA:      "coke",
			DrugB:      "pepsi",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
		DrugPair{
			DrugA:      "7up",
			DrugB:      "club",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
		DrugPair{
			DrugA:      "fizz",
			DrugB:      "lemsip",
			Delay:      Delay(0),
			DDIType:    ParallelType,
			ParentName: "parent",
		},
	}
	var drugPair = DrugPair{
		DrugA:      "coke",
		DrugB:      "pepsi",
		Delay:      Delay(0),
		DDIType:    ParallelType,
		ParentName: "parent",
	}
	if assert.True(t, drugPairListContains(listOne, drugPair), "Drug should be in list") {
		fmt.Println("PASSED!")
	}
}

func TestModifyAllChildrenNames(t *testing.T) {
	fmt.Println("Testing adding string to all children")
	baseProcess, _ := processFromFile("multi_drugs.pml")
	expectedProcess, _ := processFromFile("modified_multi_drugs.pml")
	baseProcess.ChangeNames("_1")
	if assert.Equal(t, baseProcess.Encode("  "), expectedProcess.Encode("  "), "Processes should be equal") {
		fmt.Println("PASSED!")
	}
}

func TestJoinPMLProcesses(t *testing.T) {
	fmt.Println("Testing joining processes")
	processOne, _ := processFromFile("valid_delay.pml")
	processTwo, _ := processFromFile("delays.pml")
	processThree, _ := processFromFile("no_drugs.pml")
	processFour, _ := processFromFile("subtasks.pml")
	joinedProcess := JoinPMLProcesses(processOne, processTwo, processThree, processFour)
	expectedProcess, _ := processFromFile("joined_processes.pml")
	if assert.Equal(t, joinedProcess.Encode("  "), expectedProcess.Encode("  "), "Processes should be equal") {
		fmt.Println("PASSED!")
	}
}

func TestWriteToFile(t *testing.T) {
	fmt.Println("Testing writing process to file")
	expectedProcess, _ := processFromFile("test.pml")
	WriteProcessToFile(expectedProcess, "written_process.pml")
	actualProcess, _ := processFromFile("written_process.pml")
	if assert.Equal(t, expectedProcess.Encode("  "), actualProcess.Encode("  "), "Process should be equal") {
		fmt.Println("PASSED!")
	}
}

func TestWriteMergedProcessToFile(t *testing.T) {
	fmt.Println("Testing writing merged processes to file")
	processOne, _ := processFromFile("valid_delay.pml")
	processTwo, _ := processFromFile("delays.pml")
	processThree, _ := processFromFile("no_drugs.pml")
	joinedProcess := JoinPMLProcesses(processOne, processTwo, processThree)
	WriteProcessToFile(joinedProcess, "written_joined_process.pml")
	writtenJoinedProcess, _ := processFromFile("written_joined_process.pml")
	if assert.Equal(t, joinedProcess.Encode("  "), writtenJoinedProcess.Encode("  "), "Processes should be equal") {
		fmt.Println("PASSED!")
	}
}
