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
	fmt.Println("* Testing that PML files with name clashes in different namespaces are also rejected")
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

/* Seperate helper function */
func processFromFile(filepath string) (*Element, error) {
	reader, err := os.Open(resDir + "/" + filepath)
	if err != nil {
		return &Element{}, err
	}
	parser := NewParser(reader)
	process, err := parser.Parse()
	return process, err
}

// TODO: tests for: 	Merging clinical pathways
//			PML-TX save to file
