package main

import (
	"fmt"
	"github.com/kmlasharma/cs4098/src/pml"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var pmlFilePath string
	var owlFilePath string

	fmt.Print("Enter path to PML File: ")
	fmt.Scanln(&pmlFilePath)

	fmt.Print("Enter path to OWL File: ")
	fmt.Scanln(&owlFilePath)

	checkExtension(pmlFilePath, "pml")
	checkExtension(owlFilePath, "owl")

	PEOS(pmlFilePath)
	Ontology(owlFilePath)
}

func checkExtension(path string, extension string) {
	list := strings.Split(path, ".")
	if len(list) < 2 || list[len(list)-1] != extension { // If there is a dot and the last one is not equal to the extension
		fmt.Println("Invalid file type.")
		os.Exit(0)
	}
	_, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot open file")
		os.Exit(0)
	}
}

func PEOS(path string) {
	reader, _ := os.Open(path)
	parser := pml.NewParser(reader)
	process := parser.Parse()
	fmt.Println("Process: ", process)
}

func Ontology(path string) {
	cmd := "python3 /go/src/app/dinto/ontology.py " + path
	out, _ := exec.Command("sh", "-c", cmd).Output()
	fmt.Println(string(out))
}
