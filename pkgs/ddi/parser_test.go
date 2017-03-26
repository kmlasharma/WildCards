package ddi

import (
	"os"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadInteractionFromFileLineCount(t *testing.T) {
	fmt.Println("Beginning tests on parser:")
	fmt.Println("* Testing all lines of csv file parsed.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.Equal(t,6,len(interactions), "Testing Correct line count parsed") {
		fmt.Println("PASSED!")
	}
}

func TestCorrectDrugALoad(t *testing.T) {
	fmt.Println("* Testing Drug A parsed Correctly.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.Equal(t,"coke",interactions[0].DrugA, "Testing that correct DrugA, 'coke', is parsed") {
		fmt.Println("PASSED!")
	}
}

func TestCorrectDrugBLoad(t *testing.T) {
	fmt.Println("* Testing DrugB parsed Correctly.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.Equal(t,"pepsi", interactions[0].DrugB, "DrugB should be 'pepsi'") {
		fmt.Println("PASSED!")
	}
}

func TestCorrectAdverse(t *testing.T) {
	fmt.Println("* Testing Adverse parsed correctly.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.True(t,interactions[0].Adverse, "Coke and pepsi adverse should be true") {
		fmt.Println("PASSED!")
	}
}

func TestCorrectSecParse(t *testing.T) {
	fmt.Println("* Testing second unit parsed correctly.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.Equal(t,5,interactions[0].Time, "Time parsed should be 'sec'") {
		fmt.Println("PASSED!")
	}
}

func TestCorrectMinParse(t *testing.T) {
	fmt.Println("* Testing min unit parsed correctly.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.Equal(t,120, interactions[1].Time, "Time parsed should be 120 seconds") {
		fmt.Println("PASSED!")
	}
}

func TestCorrectHourParse(t *testing.T) {
	fmt.Println("* Testing hour unit parsed correctly.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.Equal(t,86400,interactions[3].Time, "Time parsed should be 5 seconds") {
		fmt.Println("PASSED!")
	}
}

func TestCorrectWeekParse(t *testing.T) {
	fmt.Println("* Testing week unit parsed correctly.")
	interactions, _ := readInteractionsFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	if assert.Equal(t,604800,interactions[4].Time, "Time parsed should be 604800 seconds") {
		fmt.Println("PASSED!")
	}
}





