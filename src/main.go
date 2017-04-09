package main

import (
	"fmt"
	"github.com/kmlasharma/WildCards/pkgs/ddi"
	"github.com/kmlasharma/WildCards/pkgs/logger"
	"github.com/kmlasharma/WildCards/pkgs/pml"
	"os"
	"strings"
)

var resDir = os.Getenv("RES_DIR")
var testPMLFile = resDir + "/test.pml"
var testDDIFile = resDir + "/ddi.csv"
var process *pml.Element
var db *ddi.Database

func main() {
	startMessage()
	process = selectPML()

	reportWarnings()

	csvFilePath := selectCSV()
	db := ddi.NewDatabase()
	db.PopulateFromFile(csvFilePath)

	for {
		fmt.Println("\n")
		selection := getOptionSelection()
		clearScreen()
		switch selection {
		case "1":
			fmt.Println("Drugs in this PML Process:\n", strings.Join(process.AllDrugs(), ", "))
			showAllInteractions()
		case "2":
			fmt.Println("SHOWING ADVERSE DRUG INTERACTIONS, WITH CLOSEST APPROACH")
		case "3":
			showSequentialDrugPairs()
		case "4":
			showParallelDrugPairs()
		case "5":
			showAlternativeNonDDIDrugPairs()
		case "6":
			showAlternativeRepeatedDDIDrugPairs()
		case "7":
			fmt.Println("SAVE PML TO FILE")
		case "8":
			fmt.Println("MERGE PML FILES")
		case "9":
			process = selectPML()
		case "10":
			os.Exit(0)
		default:
			fmt.Println(fmt.Sprintf("\"%s\" is not a valid selection. Please try again", selection))
		}
	}
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

func selectCSV() (selectedCSVPath string) {
	fmt.Print("Enter path to CSV File: [default is ddi.csv] ")
	fmt.Scanln(&selectedCSVPath)
	selectedCSVPath = strings.TrimRight(selectedCSVPath, "\n")
	if selectedCSVPath == "" {
		selectedCSVPath = testDDIFile
	}
	checkExtension(selectedCSVPath, "csv")
	return
}

func selectPML() *pml.Element {
	var selectedPMLPath string
	fmt.Print("\nEnter path to PML File: [default is test.pml] ")
	fmt.Scanln(&selectedPMLPath)
	selectedPMLPath = strings.TrimRight(selectedPMLPath, "\n")
	if selectedPMLPath == "" {
		selectedPMLPath = testPMLFile
	}
	checkExtension(selectedPMLPath, "pml")
	return processFromFile(selectedPMLPath)
}

func startMessage() {
	fmt.Println("Welcome to the app")
	fmt.Println("\nHere is how it works:")
	fmt.Println("\t* You will choose a PML file")
	fmt.Println("\t* You will choose a CSV file for DDI's")
	fmt.Println("\t* The app will generate the following:")
	fmt.Println("\t\t1) Analysis based on your files")
	fmt.Println("\t\t2) A log file for you to read called analysis.log")
	fmt.Println("\t\t3) An error output file called analysis.err")
	fmt.Println("\nYou will now be asked to choose the files to analyse\nIf you want to use the default files then just hit enter at the prompt")
}

func processFromFile(path string) *pml.Element {
	reader, _ := os.Open(path)
	parser := pml.NewParser(reader)
	retProcess, err := parser.Parse()
	if err != nil {
		logger.Fatal("Error: Could not parse process.", err)
	}
	return retProcess
}

func getOptionSelection() string {
	var selectedOperation string
	fmt.Println("What operation would you like to complete?")
	fmt.Println(`
        1) Show All Interactions
        2) Show Adverse Drug Interactions (With Closest Approach)
        3) Show Sequential DDIs
        4) Show Parallel DDIs
        5) Show Alternative Non DDIs
        6) Show Alternative Repeated DDIs
        7) Save PML to File
        8) Merge PML Files
        9) Change to new pml file
        10) Quit Application
     `)

	fmt.Scanln(&selectedOperation)
	selectedOperation = strings.TrimRight(selectedOperation, "\n")
	return selectedOperation
}

func reportWarnings() {
	showTaskConstructs()
	// Report missing name constructs etc
}

func showTaskConstructs() {
	taskNames := []string{}
	for _, task := range process.AllTasks() {
		taskNames = append(taskNames, task.Name)
	}
	if len(taskNames) > 0 {
		fmt.Println("Tasks in this PML Process:\n", strings.Join(taskNames, ", "))
	}
}

func showSequentialDrugPairs() {
	fmt.Println("Sequential Drug Pairs:")
	fmt.Println("======================")
	findAndPrintInteractions(process.FindSequentialDrugPairs())
}

func showParallelDrugPairs() {
	fmt.Println("\nParallel Drug Pairs:")
	fmt.Println("====================")
	findAndPrintInteractions(process.FindParallelDrugPairs())
}

func showAlternativeNonDDIDrugPairs() {
	fmt.Println("Alternative Non-DDI Drug Pairs:")
	fmt.Println("================================")
	findAndPrintInteractions(process.FindAlternativeNonDDIDrugPairs())
}

func showAlternativeRepeatedDDIDrugPairs() {
	fmt.Println("Alternative Repeated Drug Pairs:")
	fmt.Println("================================")
	findAndPrintInteractions(process.FindRepeatedAlternativeDrugPairs())
}

func showAllInteractions() {
	fmt.Println("All Interactions:")
	fmt.Println("=================")
	findAndPrintInteractions(process.FindDrugPairs())
}

func findAndPrintInteractions(pairs []pml.DrugPair) {
	interactions, err := db.FindActiveInteractionsForPairs(pairs)
	if err != nil {
		fmt.Println("Couldn't find any DDI's")
		os.Exit(1)
	}
	fmt.Println(interactions)
	fmt.Println("\n")
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
