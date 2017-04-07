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

func main() {
	var csvFilePath string


	startMessage()
	process := selectPML()
	csvFilePath = selectCSV()

	var selection string
	var continueApp bool = true
	fmt.Println("LOOP NOT CURRENTLY WORKING. ENTER 5 TO CONTINUE")

	for(continueApp == true){
		fmt.Println("\n")

		selection= getOptionSelection()
		switch selection{
			case "1":
				fmt.Println("showing all interactions")
				break
			case "2":
				fmt.Println("showing adverse drug interactions")
				break
			case "3":
				fmt.Println("Pml file saved to file")
				break
			case "4":
				fmt.Println("merging pml files")
				break
			case "5":
				continueApp = false
				break
			default:
				break
		}
	}


	fmt.Println("Drugs in this PML Process:\n", strings.Join(process.AllDrugs(), ", "))

	taskNames := []string{}
	for _, task := range process.AllTasks() {
		taskNames = append(taskNames, task.Name)
	}

	fmt.Println("Tasks in this PML Process:\n", strings.Join(taskNames, ", "))

	fmt.Println("Sequential Drug Pairs:")
	fmt.Println("======================")
	for _, pair := range process.FindSequentialDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Sequence:", pair.ParentName())
	}
	fmt.Println("\n")

	fmt.Println("\nParallel Drug Pairs:")
	fmt.Println("====================")
	for _, pair := range process.FindParallelDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Parallel:", pair.ParentName())
	}
	fmt.Println("\n")

	fmt.Println("Alternative Non-DDI Drug Pairs:")
	fmt.Println("================================")
	for _, pair := range process.FindAlternativeNonDDIDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Selection:", pair.ParentName())
	}
	fmt.Println("\n")

	fmt.Println("Alternative Repeated Drug Pairs:")
	fmt.Println("================================")
	for _, pair := range process.FindRepeatedAlternativeDrugPairs() {
		fmt.Println("DrugA:", pair.DrugA, ", Drug B:", pair.DrugB, ", Parent Selection:", pair.ParentName())
	}

	fmt.Println("\nEncoded:\n", process.Encode(""))


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

func selectCSV() (selectedCSVPath string){
	fmt.Print("Enter path to CSV File: [default is ddi.csv] ")
	fmt.Scanln(&selectedCSVPath)
	selectedCSVPath = strings.TrimRight(selectedCSVPath, "\n")

	if selectedCSVPath == "" {
		selectedCSVPath = testDDIFile
	}

	checkExtension(selectedCSVPath, "csv")
	return selectedCSVPath
}

func selectPML() (*pml.Element){
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
	process, err := parser.Parse()
	if err != nil {
		logger.Fatal("Error: Could not parse process.", err)
	}
	return process
}

func getOptionSelection() string{
	var selectedOperation string
	fmt.Println("What operation would you like to complete?")
	fmt.Println("\n 1) Show All Interactions")
	fmt.Println("\n 2) Show Adverse Drug Interactions (With Closest Approach)")
	fmt.Println("\n 3) Save PML to File ")
	fmt.Println("\n 4) Merge PML Files ")
	fmt.Println("\n 5) Quit App")

	fmt.Println("Please enter number responding to operation")

	fmt.Scanln(&selectedOperation)
	selectedOperation = strings.TrimRight(selectedOperation, "\n")
	return selectedOperation
}
