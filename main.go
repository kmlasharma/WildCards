package main

import (
	"fmt"
	// "io/ioutil"
	"os"
	"strings"
	"github.com/KmlaSharma/WildCards/dinto"
)

func main() {
	fmt.Print("Enter path to PML File: ")
    var pmlFilePath string
    fmt.Scanln(&pmlFilePath)
    fmt.Println(pmlFilePath)
    checkExtension(pmlFilePath, "pml")

    fmt.Print("Enter path to OWL File: ")
    var owlFilePath string
    fmt.Scanln(&owlFilePath)
    fmt.Println(owlFilePath)
    checkExtension(owlFilePath, "owl")

	var ontology = dinto.GenerateOntology(owlFilePath)
	fmt.Println("Number of Prefixes:", len(ontology.Prefixes))
	fmt.Println("Number of Imports:", len(ontology.Imports))
}

func checkExtension(path string, extension string) {
	list := strings.Split(path, ".")
	if len(list) < 2 || list[len(list)-1] != extension { // If there is a dot and the last one is not equal to the extension
		fmt.Println("Invalid file type.")
		os.Exit(0)
	}
} 