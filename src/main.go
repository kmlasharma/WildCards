package main

import (
	"fmt"
	"github.com/kmlasharma/WildCards/pkgs/logger"
	"github.com/kmlasharma/WildCards/pkgs/pml"
	"github.com/kmlasharma/WildCards/pkgs/progressbar"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var pmlFilePath string
	var owlFilePath string

	fmt.Println("Welcome to the app")
	fmt.Println("\nHere is how it works:")
	fmt.Println("\t* You will choose a PML file")
	fmt.Println("\t* You will choose a OWL file")
	fmt.Println("\t* The app will generate the following:")
	fmt.Println("\t\t1) Analysis based on your files")
	fmt.Println("\t\t2) A log file for you to read called analysis.log")
	fmt.Println("\t\t3) An error output file called analysis.err")
	fmt.Println("\nYou will now be asked to choose the files to analyse\nIf you want to use the default files then just hit enter at the prompt")

	fmt.Print("\nEnter path to PML File: [default is test.pml] ")
	fmt.Scanln(&pmlFilePath)
	pmlFilePath = strings.TrimRight(pmlFilePath, "\n")

	fmt.Print("Enter path to OWL File: [default is test.owl] ")
	fmt.Scanln(&owlFilePath)
	owlFilePath = strings.TrimRight(owlFilePath, "\n")

	if pmlFilePath == "" {
		pmlFilePath = "test.pml"
	}
	if owlFilePath == "" {
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
		logger.Println("Invalid file type.")
		os.Exit(0)
	}
	_, err := os.Open(path)
	if err != nil {
		logger.Println("Cannot open file")
		logger.Println(err)
		os.Exit(0)
	}
}

func PEOS(path string) {
	reader, _ := os.Open(path)
	parser := pml.NewParser(reader)
	process := parser.Parse()
	logger.Println("Process: ", process.Name)
	logger.Println("Drugs in Process:\n", strings.Join(process.AllDrugs(), "\n"))
}

func Ontology(path string) {
	progressbar.DisplayProgressBarForOwlFile(path)
	cmd := "python3 /go/src/app/dinto/ontology.py " + path
	out, _ := exec.Command("sh", "-c", cmd).Output()
	logger.Println(string(out))
}
