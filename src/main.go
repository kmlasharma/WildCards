package main

import (
	"fmt"
	"github.com/kmlasharma/WildCards/pkgs/ddi"
	"github.com/kmlasharma/WildCards/pkgs/logger"
	"github.com/kmlasharma/WildCards/pkgs/pml"
	"os"
	"strings"
)

func main() {
	var pmlFilePath string
	var csvFilePath string

	resDir := os.Getenv("RES_DIR")
	testPMLFile := resDir + "/test.pml"
	testDDIFile := resDir + "/ddi.csv"

	fmt.Println("Welcome to the app")
	fmt.Println("\nHere is how it works:")
	fmt.Println("\t* You will choose a PML file")
	fmt.Println("\t* You will choose a CSV file for DDI's")
	fmt.Println("\t* The app will generate the following:")
	fmt.Println("\t\t1) Analysis based on your files")
	fmt.Println("\t\t2) A log file for you to read called analysis.log")
	fmt.Println("\t\t3) An error output file called analysis.err")
	fmt.Println("\nYou will now be asked to choose the files to analyse\nIf you want to use the default files then just hit enter at the prompt")

	fmt.Print("\nEnter path to PML File: [default is test.pml] ")
	fmt.Scanln(&pmlFilePath)
	pmlFilePath = strings.TrimRight(pmlFilePath, "\n")

	if pmlFilePath == "" {
		fmt.Println("### Using the default PML file " + testPMLFile + " ###")
		pmlFilePath = testPMLFile
	}

	checkExtension(pmlFilePath, "pml")
	process := processFromFile(pmlFilePath)

	fmt.Print("Enter path to CSV File: [default is ddi.csv] ")
	fmt.Scanln(&csvFilePath)
	csvFilePath = strings.TrimRight(csvFilePath, "\n")

	if csvFilePath == "" {
		fmt.Println("### Using the default DDI file " + testDDIFile + " ###")
		csvFilePath = testDDIFile
	}

	checkExtension(csvFilePath, "csv")

	db := ddi.NewDatabase()
	db.PopulateFromFile(csvFilePath)
	interactions, err := db.FindInteractions(process.AllDrugs())

	if err != nil {
		fmt.Println("Couldn't find any DDI's")
		os.Exit(1)
	}

	fmt.Println("DDI's for this PML File:", interactions)
}

func checkExtension(path string, extension string) {
	list := strings.Split(path, ".")
	if len(list) < 2 || list[len(list)-1] != extension { // If there's a dot and the last one is not equal to the extension
		logger.Error("Invalid file type.")
		os.Exit(0)
	}
	_, err := os.Open(path)
	if err != nil {
		logger.Error("ERROR: Cannot open file")
		logger.Error(err)
		os.Exit(0)
	}
}

func processFromFile(path string) *pml.Element {
	reader, _ := os.Open(path)
	parser := pml.NewParser(reader)
	process, err := parser.Parse()
	if err != nil {
		logger.Fatal("Error: Could not parse process. Error:", err)
	}
	return process
}
