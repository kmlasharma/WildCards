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
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
		DrugPair{
			DrugA:      "7up",
			DrugB:      "club",
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
		DrugPair{
			DrugA:      "fizz",
			DrugB:      "lemsip",
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
	}
	var listTwo = []DrugPair{
		DrugPair{
			DrugA:      "fizz",
			DrugB:      "lemsip",
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
		DrugPair{
			DrugA:      "coke",
			DrugB:      "pepsi",
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
		DrugPair{
			DrugA:      "7up",
			DrugB:      "club",
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
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
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
		DrugPair{
			DrugA:      "7up",
			DrugB:      "club",
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
		DrugPair{
			DrugA:      "fizz",
			DrugB:      "lemsip",
			delay:      Delay(0),
			ddiType:    ParallelType,
			parentName: "parent",
		},
	}
	var drugPair = DrugPair{
		DrugA:      "coke",
		DrugB:      "pepsi",
		delay:      Delay(0),
		ddiType:    ParallelType,
		parentName: "parent",
	}
	if assert.True(t, drugPairListContains(listOne, drugPair), "Drug should be in list") {
		fmt.Println("PASSED!")
	}
}

func TestModifyAllChildrenNames(t * testing.T) {
	fmt.Println("Testing adding string to all children")
	baseProcess, _ := processFromFile("multi_drugs.pml")
	expectedProcess, _ := processFromFile("modified_multi_drugs.pml")
	baseProcess.ChangeNames("_1")
	if assert.Equal(t, baseProcess.Encode("  "), expectedProcess.Encode("  "), "Processes should be equal") {
		fmt.Println("PASSED!")
	}
}

//func TestJoinPMLProcesses(t * testing.T) {
	// TODO
//	fmt.Println("Testing joining processes")
//	JoinPMLProcesses([]*Element{})
//}

func TestCombineProcesses(t * testing.T) {
	fmt.Println("Testing joining two processes")
	firstProcess, _ := processFromFile("multi_drugs.pml")
	secondProcess, _ := processFromFile("test.pml")

	expectedProcess, _ := processFromFile("two_combined_processes.pml")
	CombineProcesses(firstProcess, secondProcess)
	actualProcess := firstProcess
	if assert.NotEqual(t, expectedProcess, actualProcess, "Process do not match") {
		fmt.Println("PASSED!")
	}
}

