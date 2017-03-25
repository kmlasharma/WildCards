package pml

import (
	"fmt"
	"os"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoProcesses(t *testing.T) {
	fmt.Println("* Testing loading PML file with no processes")
	reader, _ := os.Open("/root/no_processes.pml") // empty file
	parser := NewParser(reader)
	// we use os.exit so it's not really possible to test this at the moment
	process := parser.Parse()
	if(assert.Nil(t, process, "Process should not exist")) {
		fmt.Println("PASSED!")
	}
}

// our application doesn't currently handle multiple processes
// we should probably add a method for fetching all processes
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
	// TODO
	// Need clarification here: are we checking only for clashes in process names?
	// What about a process that contains tasks with overlapping names?
	// Or overlapping task names over different processes?
	fmt.Println("* Testing that PML files with process name clashes are rejected")


	fmt.Println("Unimplemented...")
}

func TestValidateNoClashes(t * testing.T) {
	// TODO
	// Need clarification here: are we checking only for clashes in process names?
	// What about a process that contains tasks with overlapping names?
	// Or overlapping task names over different processes?
	fmt.Println("* Testing that PML files with no process name clashes are not rejected")


	fmt.Println("Unimplemented...")
}

// TODO: tests for broken PML files
