package main

import (
	"fmt"
	"github.com/kmlasharma/WildCards/src/pml"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var pmlFilePath string
	var owlFilePath string

	fmt.Print("Welcome to the app\n")
	fmt.Print("Here is how it works:")
	fmt.Print("\t* You will choose a PML file")
	fmt.Print("\t* You will choose a OWL file")
	fmt.Print("\t* The app will generate the following:")
	fmt.Print("\t\t1) Analysis based on your files")
	fmt.Print("\t\t2) A log file for you to read called analysis.log")
	fmt.Print("\t\t3) An error output file called analysis.err")
	fmt.Print("\nYou will now be asked to choose the files to analyse\nIf you want to use the default files then just hit enter at the prompt")

	fmt.Print("Enter path to PML File: [default is test.pml]")
	fmt.Scanln(&pmlFilePath)

	fmt.Print("Enter path to OWL File: [default is test.owl]")
	fmt.Scanln(&owlFilePath)

	if strings.Compare(pmlFilePath, "") == 0 {
		pmlFilePath = "test.pml"
	}
	if strings.Compare(owlFilePath, "") == 0 {
		owlFilePath = "test.owl"
	}

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
