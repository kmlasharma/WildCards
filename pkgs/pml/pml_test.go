package pml

import (
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
	"testing"
)

var resDir = os.Getenv("RES_DIR")

func TestNoProcesses(t *testing.T) {
	// we use os.exit so we need to keep this commented until it is removed

	//fmt.Println("* Testing loading PML file with no processes")
	//reader, _ := os.Open("/root/no_processes.pml") // empty file
	//parser := NewParser(reader)
	//process := parser.Parse()
	//if(assert.Nil(t, process, "Process should not exist")) {
	//	fmt.Println("PASSED!")
	//}
}

// our application doesn't currently handle multiple processes
// need clarification on whether or not this is required
func TestMultipleProcesses(t * testing.T) {
	//TODO
	fmt.Println("* Testing loading PML file with multiple processes")


	fmt.Println("Unimplemented...")
}

func TestNoSequences(t *testing.T) {
	fmt.Println("* Testing analysing process with no sequences")
	path := os.Getenv("RES_DIR")
	reader , _ := os.Open(path + "/no_sequences.pml")
	parser := NewParser(reader)
	process := parser.Parse()
	drugs := process.AllDrugs()
	tasks := process.AllTasks()
	assert.Equal(t, len(drugs), 0, "Process should have no drugs")
	if (assert.Equal(t, len(tasks), 0, "Process should have no subtasks")) {
		fmt.Println("SUCCESS!")
	}
}

func TestNoDrugs(t * testing.T) {
	fmt.Println("* Testing analysing process with a script with no drugs")
	path := os.Getenv("RES_DIR")
	reader , _ := os.Open(path + "/no_drugs.pml")
	parser := NewParser(reader)
	process := parser.Parse()
	drugs := process.AllDrugs()
	if(assert.Equal(t, len(drugs), 0, "Process should have no drugs")) {
		fmt.Println("SUCCESS!")
	}
}

func TestMultipleDrugs(t *testing.T) {
	fmt.Println("* Testing analysing process with multiple sequences and actions containing drugs")
	path := os.Getenv("RES_DIR")
	reader , _ := os.Open(path + "/multi_drugs.pml")
	parser := NewParser(reader)
	process := parser.Parse()
	drugs := process.AllDrugs()
	expected := []string{"coke", "7up", "pepsi", "fanta", "dr pepper"}
	if(assert.Equal(t, drugs, expected, "Drugs should include exactly coke, 7up, pepsi, fanta and dr pepper")) {
		fmt.Println("SUCCESS!")
	}
}

func TestValidateSameType(t * testing.T) {
	fmt.Println("* Testing that PML files with name clashes among items of the same type are rejected")
	reader, _ := os.Open(resDir + "/sequence_clashes.pml")
	parser := NewParser(reader)
	process := parser.Parse()
	err := process.Validate()
	if(assert.NotNil(t, err, "Error should be raised due to mutiply defined identifiers")) {
		if(assert.Equal(t, err.Error(), "Multiply defined identifiers: Andy, Mary", "Andy and Mary should be defined multiple times")) {
			fmt.Println("PASSED!")
		}
	}
}

func TestValidateDifferentTypes(t * testing.T) {
	fmt.Println("* Testing that PML files with name clashes among items of different types are rejected")
	reader, _ := os.Open(resDir + "/multilevel_clashes.pml")
	parser := NewParser(reader)
	process := parser.Parse()
	err := process.Validate()
	if(assert.NotNil(t, err, "Error should be raised due to mutiply defined identifiers")) {
		if(assert.Equal(t, err.Error(), "Multiply defined identifiers: Mary, Andy, John", "Mary, Andy and John should be defined multiple times")) {
			fmt.Println("PASSED!")
		}
	}
}

func TestValidateNoClashes(t * testing.T) {
	fmt.Println("* Testing that PML files with no name clashes are not rejected")
	reader, _ := os.Open(resDir + "/no_clashes.pml") 
	parser := NewParser(reader)
	process := parser.Parse()
	if(assert.Nil(t, process.Validate(), "There should be no name clashes detected")) {
		fmt.Println("PASSED!")
	}
}

// TODO: tests for broken PML files
