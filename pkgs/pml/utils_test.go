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
	//for _, e := range listOne {
	//	var x = (e == drugPair)
	//	fmt.Println(x)
	//}
	if assert.True(t, drugPairListContains(listOne, drugPair), "Drug should be in list") {
		fmt.Println("PASSED!")
	}
}
