package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var resDir = os.Getenv("RES_DIR")

var db = NewDatabase()

/*
func TestTimeIntervalOffset(t *testing.T) {
 	fmt.Println("* Testing that time interval offsets are processed")
 	process, err := processFromFile("time_interval_offset.pml")
	drugPair := process.FindDrugPairs()[0]
	actualDDIType := Pull drug pair interaction from DDI database
		expectedDDIType := "NOT ADVERSE"
	---
	if assert.Nil(t, err, "Error with PML file") && assert.Equal(t, expectedDDIType, actualDDIType, "Time interval offset not correctly registered -resulting in wrong DDI type for drug pair")
		fmt.Println("PASSED!")
	}
}
*/

func TestDDIClosestApproach(t *testing.T) {
	fmt.Println("* Testing DDI closest approach")
	assert := setup(t)
	process, procErr := processFromFile("closest_approach.pml")
	actualInteractions, interactionErr := ddi.FindActiveInteractionsForPairs(process.FindDrugPairs())
	expectedInteraction := ddi.Interaction{"caffeine","alcohol", false, 604800}
	if assert.Nil(procErr, "There should not be an error with the process") && assert.Nil(interactionErr, "There should not be an error finding interactions") && assertEqual(1, len(actualInteractions), "Too many interactions found") && assert.Equal(actualInteractions[0], expectedInteraction, fmt.Sprintf("Wrong interation found, expected { %s } found { %s }", expectedInteraction, actualInteractions[0])) {
		fmt.Println("PASSED!")
	}
}

func processFromFile(filepath string) (*Element, error) {
	reader, err := os.Open(resDir + "/" + filepath)
	if err != nil {
		return &Element{}, err
	}
	parser := NewParser(reader)
	process, err := parser.Parse()
	return process, err
}


