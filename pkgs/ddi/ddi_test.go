package ddi

import (
	"os"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var db = NewDatabase()

func TestExistenceOfInteraction(t *testing.T) {
	fmt.Println("* Testing Existence of Interaction")
	assert := setup(t)

	_, err := db.FindInteraction("pepsi", "flat7up")
	if assert.Nil(err, "Pepsi + flat7up interaction should exist") {
		fmt.Println("PASSED!")
	}
}

func TestAbsenceOfInteraction(t *testing.T) {
	fmt.Println("* Testing Expected Absence of Interaction")
	assert := setup(t)

	_, err := db.FindInteraction("coke", "flat7up")
	if assert.NotNil(err, "Coke + flat7up interaction should not exist") {
		fmt.Println("PASSED!")
	}
}

func TestCountOfInteraction(t *testing.T) {
	fmt.Println("* Testing Existence of Interaction")
	assert := setup(t)

	interactions, _ := db.FindInteractions([]string{"pepsi", "flat7up"})
	if assert.Equal( 1,len(interactions), "Pepsi + flat7up interaction should exist") {
		fmt.Println("PASSED!")
	}
}

func TestFindAdverseInteraction(t *testing.T) {
	fmt.Println("* Testing Finding an Adverse Interaction..")
	assert := setup(t)

	interaction, _ := db.FindInteraction("oj", "7up")
	if assert.True(interaction.Adverse, "OJ + 7up interaction should be adverse") {
		fmt.Println("PASSED!")
	}
}

func TestFindNonAdverseInteraction(t *testing.T) {
	fmt.Println("* Testing Finding an Non Adverse Interaction..")
	assert := setup(t)

	interaction, _ := db.FindInteraction("coke", "alcohol")
	if assert.True(!interaction.Adverse, "Coke + Pepsi interaction should be non adverse") {
		fmt.Println("PASSED!")
	}
}

func TestEnsureCorrectTimingForInteraction(t *testing.T) {
	fmt.Println("* Testing timing for Interaction..")
	assert := setup(t)

	interaction, _ := db.FindInteraction("paracetamol", "alcohol")
	if assert.Equal(24*60*60,interaction.Time, "Paracetamol and Alcohol should have a gap of 1 day") {
		fmt.Println("PASSED!")
	}
}

func setup(t *testing.T) *assert.Assertions {
	db.Clear()
	db.PopulateFromFile(os.Getenv("RES_DIR") + "/ddi.csv")
	return assert.New(t)
}

// TODO: tests for:	Reporting alternate non-DDIs
//			Reporting repeated alternate DDIs

