package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/kmlasharma/WildCards/pkgs/ddi"
	"os"
)

func TestTimeIntervalOffset(t *testing.T) {
 	fmt.Println("* Testing that time interval offsets are processed")
	assert := setup(t)
 	process := processFromFile(resDir + "/" + "time_interval_offset.pml")
	drugPair := process.FindDrugPairs()[0]
	_, notFoundErr := db.FindActiveInteractionForPair(drugPair)
	if assert.NotNil(notFoundErr, "Time interval offset not correctly registered") {
		fmt.Println("PASSED!")
	}
}

func TestDDIClosestApproach(t *testing.T) {
	fmt.Println("* Testing DDI closest approach")
	assert := setup(t)
	process := processFromFile(resDir + "/" + "closest_approach.pml")	
	actualInteractions, interactionErr := db.FindActiveInteractionsForPairs(process.FindDrugPairs())
	expectedInteraction := ddi.Interaction{ "caffeine","alcohol", true, 604800 }
	if assert.Nil(interactionErr, "There should not be an error finding interactions") && assert.Equal(1, len(actualInteractions), "Wrong number of interactions found") && assert.Equal(actualInteractions[0], expectedInteraction, fmt.Sprintf("Wrong interation found, expected { %s } found { %s }", expectedInteraction, actualInteractions[0])) {
		fmt.Println("PASSED!")
	}
}


func setup(t *testing.T) *assert.Assertions {
	db = ddi.NewDatabase()
	db.PopulateFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	return assert.New(t)
}

