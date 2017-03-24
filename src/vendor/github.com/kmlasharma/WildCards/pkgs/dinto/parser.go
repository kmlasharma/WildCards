package dinto

import (
	"fmt"
	"encoding/csv"
	"os"
	"io"
    "strconv"
    "strings"
)

const SECONDS_TO_MIN = 60
const SECONDS_TO_HOUR = 3600
const SECONDS_TO_DAY = 86400
const SECONDS_TO_WEEK = 604800

const DRUG_1 = "Drug A"
const DRUG_2 = "Drug B"
const DDI_TYPE = "DDI Type"
const TIME = "Time"
const UNIT = "Unit"

func ReadInteractionsFromFile(filepath string) (interactions []Interaction, err error) {
	f, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    csvr := csv.NewReader(f)
    header, err := csvr.Read() // read in header
    var headerToIndex = analyseHeader(header)
    fmt.Println(headerToIndex)
    if err != nil {
    	return nil, err
    }
    for {
        row, err := csvr.Read()
        if err != nil {
        	if err == io.EOF { //reached end of file
                err = nil
                return interactions, err
            }
            return nil, err
        }
        interaction, err := rowToInteractionObject(row, headerToIndex)
        if (err == nil) {
            interactions = append(interactions, interaction)
        }
    }
    f.Close()
    fmt.Println(interactions)
	return interactions, nil
}

func convertToSeconds(time int, unit string) (timeInSeconds int) {

    switch unit {
    case "min":
        timeInSeconds = time * SECONDS_TO_MIN
    case "hr":
        timeInSeconds = time * SECONDS_TO_HOUR
    case "day":
        timeInSeconds = time * SECONDS_TO_DAY
    case "week":
        timeInSeconds = time * SECONDS_TO_WEEK
    default:
        timeInSeconds = time //assume it is seconds
    }
    return timeInSeconds
}

func analyseHeader(header []string) (headerToIndex map[string]int) {
    headerToIndex = make(map[string]int)
    for i, _ := range header {
        if strings.Contains(header[i], "1") || strings.Contains(header[i], "A") {
            headerToIndex[DRUG_1] = i
        } else if strings.Contains(header[i], "2") || strings.Contains(header[i], "B") {
            headerToIndex[DRUG_2] = i
        } else if strings.Contains(header[i], "Type") {
            headerToIndex[DDI_TYPE] = i
        } else if strings.Contains(header[i], "Time") {
            headerToIndex[TIME] = i
        } else if strings.Contains(header[i], "Unit") {
            headerToIndex[UNIT] = i
        }
    }
    return headerToIndex
}

func rowToInteractionObject(row []string, headerToIndex map[string]int) (interaction Interaction, err error) {
    drugA := row[headerToIndex[DRUG_1]]
    drugB := row[headerToIndex[DRUG_2]]
    goodBad := row[headerToIndex[DDI_TYPE]]
    adverse := goodBad == "bad"
    time, err := strconv.Atoi(row[headerToIndex[TIME]])
    if err != nil {
        fmt.Println("Error occurred parsing integer from csv.")
        return Interaction{}, err
    }
    unit := row[headerToIndex[UNIT]]
    time = convertToSeconds(time, unit)
    interaction = Interaction{DrugA: drugA, DrugB: drugB, Adverse: adverse, Time: time}
    return interaction, nil
}