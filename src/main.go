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

func main() {
	startMessage()
	process = selectPML()
	csvFilePath := selectCSV()
	continueApp := true

	for continueApp {
		fmt.Println("\n")
		selection := getOptionSelection()
		clearScreen()
		switch selection {
		case "1":
			fmt.Println("Drugs in this PML Process:\n", strings.Join(process.AllDrugs(), ", "))
			showAllInteractions(csvFilePath)
		case "2":
			fmt.Println("SHOWING ADVERSE DRUG INTERACTIONS, WITH CLOSEST APPROACH")
		case "3":
			showTaskConstructs()
		case "4":
			showSequentialDrugPairs()
		case "5":
			showParallelDrugPairs()
		case "6":
			showAlternativeNonDDIDrugPairs()
		case "7":
			showAlternativeRepeatedDDIDrugPairs()
		case "8":
			fmt.Println("SAVE PML TO FILE")
		case "9":
			fmt.Println("MERGE PML FILES")
		case "10":
			process = selectPML()
		case "11":
			continueApp = false
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
        3) Show Task Constructs
        4) Show Sequential DDIs
        5) Show Parallel DDIs
        6) Show Alternative Non DDIs
        7) Show Alternative Repeated DDIs
        8) Save PML to File
        9) Merge PML Files
        10) Change to new pml file
        11) Quit Application
     `)

	fmt.Scanln(&selectedOperation)
	selectedOperation = strings.TrimRight(selectedOperation, "\n")
	return selectedOperation
}

func showTaskConstructs() {
	taskNames := []string{}
	for _, task := range process.AllTasks() {
		taskNames = append(taskNames, task.Name)
	}
	fmt.Println("Tasks in this PML Process:\n", strings.Join(taskNames, ", "))
}

func showSequentialDrugPairs() {
	fmt.Println("Sequential Drug Pairs:")
	fmt.Println("======================")
	for _, pair := range process.FindSequentialDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Sequence:", pair.ParentName())
	}
	fmt.Println("\n")
}

func showParallelDrugPairs() {
	fmt.Println("\nParallel Drug Pairs:")
	fmt.Println("====================")
	for _, pair := range process.FindParallelDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Branch:", pair.ParentName())
	}
	fmt.Println("\n")
}

func showAlternativeNonDDIDrugPairs() {
	fmt.Println("Alternative Non-DDI Drug Pairs:")
	fmt.Println("================================")
	for _, pair := range process.FindAlternativeNonDDIDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Selection:", pair.ParentName())
	}
	fmt.Println("\n")
}

func showAlternativeRepeatedDDIDrugPairs() {
	fmt.Println("Alternative Repeated Drug Pairs:")
	fmt.Println("================================")
	for _, pair := range process.FindRepeatedAlternativeDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Selection:", pair.ParentName())
	}

	fmt.Println("\n")
}

func showAllInteractions(csvPath string) {
	db := ddi.NewDatabase()
	db.PopulateFromFile(csvPath)
	interactions, err := db.FindInteractions(process.AllDrugs())
	if err != nil {
		fmt.Println("Couldn't find any DDI's")
		os.Exit(1)
	}
	fmt.Println("DDI's for this PML File:", interactions)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
