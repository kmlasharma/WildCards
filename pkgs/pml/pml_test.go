package pml

import (
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
	"testing"
)

var resDir = os.Getenv("RES_DIR")

func TestNoProcesses(t *testing.T) {
	// we use os.exit so it's not really possible to test this at the moment

	fmt.Println("* Testing loading PML file with no processes")
	reader, _ := os.Open("/root/no_processes.pml") // empty file
	parser := NewParser(reader)
	process := parser.Parse()
	if(assert.Nil(t, process, "Process should not exist")) {
		fmt.Println("PASSED!")
	}
}

// our application doesn't currently handle multiple processes
// need clarification on whether or not this is required
func TestMultipleProcesses(t * testing.T) {
	//TODO
	fmt.Println("* Testing loading PML file with multiple processes")


	fmt.Println("Unimplemented...")
}

func TestNoSequences(t *testing.T) {
	//TODO
	fmt.Println("* Testing analysing process with no sequences")


	fmt.Println("Unimplemented...")
}

func TestNoDrugs(t * testing.T) {
	//TODO
	fmt.Println("* Testing analysing process with a sequence & action that contains no drugs")


	fmt.Println("Unimplemented...")
}

func TestMultipleDrugs(t *testing.T) {
	//TODO
	fmt.Println("* Testing analysing process with actions that contain drugs")


	fmt.Println("Unimplemented...")
}

func TestMultipleTasksDrugs(t *testing.T) {
	//TODO
	fmt.Println("* Testing loading PML file with multiple tasks with multiple drugs")


	fmt.Println("Unimplemented...")
}

func TestValidateClashes(t * testing.T) {
	// Need clarification here: are we checking only for clashes in sequence names?
	// What about a sequence that contains actions with overlapping names?
	// Or overlapping action names over different sequences?
	// Or processes with overlapping names?
	fmt.Println("* Testing that PML files with task name clashes are rejected")
	reader, _ := os.Open(resDir + "/sequence_clashes.pml") // empty file
	parser := NewParser(reader)
	process := parser.Parse()
	err := process.Validate()
	fmt.Println(err)
	//if(assert.Equal(t, err.Error(), "Multiply defined sequence: mySeq")) {
	//	fmt.Println("PASSED!")
	//}
}

func TestValidateNoClashes(t * testing.T) {
	// Need clarification here: are we checking only for clashes in process names?
	// What about a process that contains tasks with overlapping names?
	// Or overlapping task names over different processes?
	fmt.Println("* Testing that PML files with no task name clashes are not rejected")
	reader, _ := os.Open(resDir + "/no_sequence_clashes.pml") // empty file
	parser := NewParser(reader)
	process := parser.Parse()
	if(assert.Nil(t, process.Validate(), "There should be no task name clashes detected")) {
		fmt.Println("PASSED!")
	}
}

// TODO: tests for broken PML files
