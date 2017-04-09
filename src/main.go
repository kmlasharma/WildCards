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

	reportInfo()

	csvFilePath := selectCSV()
	db = ddi.NewDatabase()
	db.PopulateFromFile(csvFilePath)

	for {
		fmt.Println("\n")
		selection := getOptionSelection()
		clearScreen()
		switch selection {
		case "1":
			showAllInteractions()
		case "2":
			showAdverseInteractions()
		case "3":
			showSequentialDrugPairs()
		case "4":
			showParallelDrugPairs()
		case "5":
			showAlternativeNonDDIDrugPairs()
		case "6":
			showAlternativeRepeatedDDIDrugPairs()
		case "7":
			savePMLToFile()
		case "8":
			mergePMLFile()
		case "9":
			process = selectPML()
			reportInfo()
		case "10":
			showClosestApproaches()
		case "11":
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
	fmt.Println("\t\t1) A log file for you to read called analysis.log")
	fmt.Println("\t\t2) An error output file called analysis.err")
	fmt.Println("\t* You will then be presented with a more detailed menu for analysing the data")
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
        5) Show Alternative Non-DDIs
        6) Show Alternative Repeated DDIs
        7) Save PML to File
        8) Merge PML Files
        9) Change to new PML file
	10) Show DDI closest approaches
        11) Quit Application
     `)

	fmt.Scanln(&selectedOperation)
	selectedOperation = strings.TrimRight(selectedOperation, "\n")
	return selectedOperation
}

func reportInfo() {
	fmt.Println("\n")
	showAllDrugsInProcess()
	showTaskConstructs()
	showPeriodicDrugUse()
	showDelays()
	fmt.Println("\n")
}

func showAllDrugsInProcess() {
	fmt.Println("Drugs in Process:", strings.Join(process.AllDrugs(), ", "))
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

func showPeriodicDrugUse() {
	for _, iter := range process.AllPeriodicIterations() {
		fmt.Println("Periodic Drug Use in iteration", iter.Name)
	}
}

func showDelays() {
	for _, res := range process.AllWithDelays() {
		str := fmt.Sprintf("Delay of %s within element '%s'", res.Delay.HumanReadableTime(), res.Element.Name)
		fmt.Println(str)
	}
}

func mergePMLFile() {
	fmt.Println("Provide file path of PML file to merge.")
	newProcess := selectPML()
	process = pml.JoinPMLProcesses(process, newProcess)
	fmt.Println("Merged PML file into current Process. Save the file to see the result.")
}

func savePMLToFile() {
	var savedPMLFileName string
	fmt.Print("\nEnter filename to save PML File to: ")
	fmt.Scanln(&savedPMLFileName)
	pml.WriteProcessToFile(process, savedPMLFileName)
	fmt.Println("Saved PML File to", savedPMLFileName)
}

func showSequentialDrugPairs() {
	fmt.Println("Sequential DDIs:")
	fmt.Println("======================")
	findAndPrintInteractions(process.FindSequentialDrugPairs(), false)
}

func showParallelDrugPairs() {
	fmt.Println("Parallel DDIs:")
	fmt.Println("====================")
	findAndPrintInteractions(process.FindParallelDrugPairs(), false)
}

func showAlternativeNonDDIDrugPairs() {
	fmt.Println("Alternative Non-DDIs:")
	fmt.Println("================================")
	findAndPrintNonDDIs(process.FindAlternativeNonDDIDrugPairs(), false)
}

func showAlternativeRepeatedDDIDrugPairs() {
	fmt.Println("Alternative Repeated DDIs:")
	fmt.Println("================================")
	findAndPrintInteractions(process.FindRepeatedAlternativeDrugPairs(), false)
}

func showAllInteractions() {
	fmt.Println("All Interactions:")
	fmt.Println("=================")
	findAndPrintInteractions(process.FindDrugPairs(), false)
}

func showAdverseInteractions() {
	fmt.Println("All Adverse Interactions:")
	fmt.Println("=================")
	findAndPrintInteractions(process.FindDrugPairs(), true)
}

func findAndPrintNonDDIs(pairs []pml.DrugPair, onlyAdverse bool) {
	for _, pair := range pairs {
		altered_pair := pair
		altered_pair.DDIType = pml.SequentialType
		altered_pair.Delay = pml.Delay(0)
		interaction, err := db.FindActiveInteractionForPair(altered_pair)
		if err == nil && (!onlyAdverse || interaction.Adverse) {
			var adverse = "Yes"
			if !interaction.Adverse {
				adverse = "No"
			}
			fmt.Println(fmt.Sprintf("Drug A: \"%s\", Drug B: \"%s\", Adverse Interaction: \"%s\", Parent Name: \"%s\"", interaction.DrugA, interaction.DrugB, adverse, pair.ParentName))
		}
	}
}

func findAndPrintInteractions(pairs []pml.DrugPair, onlyAdverse bool) {
	for _, pair := range pairs {
		fmt.Println(pair)
		interaction, err := db.FindActiveInteractionForPair(pair)
		if err == nil && (!onlyAdverse || interaction.Adverse) {
			var adverse = "Yes"
			if !interaction.Adverse {
				adverse = "No"
			}
			fmt.Println(fmt.Sprintf("Drug A: \"%s\", Drug B: \"%s\", Adverse Interaction: \"%s\", Parent Name: \"%s\", Closest Approach: \"%s\", Interaction Timeframe: \"%s\"", interaction.DrugA, interaction.DrugB, adverse, pair.ParentName, pair.Delay.HumanReadableTime(), interaction.HumanReadableTime()))
		}
	}
}

func showClosestApproaches() {
	fmt.Println("DDI Closest Approaches:")
	fmt.Println("======================")
	pairs := process.FindDrugPairs()
	var new_pairs []pml.DrugPair
	for _, pair := range pairs {
		altered_pair := pair
		altered_pair.Delay = pml.Delay(0)
		altered_pair.DDIType = pml.SequentialType
		_, err := db.FindActiveInteractionForPair(altered_pair)
		if err == nil {
			new_pairs = append(new_pairs, pair)
		}
	}
	for _, pair := range new_pairs {
		var time string
		if pair.DDIType == pml.AlternativeNonDDIType {
			time = "infinite"
		} else {
			time = pair.Delay.HumanReadableTime()
		}
		fmt.Println(fmt.Sprintf("Drug A: %s, Drug B: %s, Closest Approach: %s", pair.DrugA, pair.DrugB, time))
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
