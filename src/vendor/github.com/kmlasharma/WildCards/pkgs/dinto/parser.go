package dinto

import (
	"fmt"
	"encoding/csv"
	"os"
	"io"
    "strconv"
)

const SECONDS_TO_MIN = 60
const SECONDS_TO_HOUR = 3600
const SECONDS_TO_DAY = 86400
const SECONDS_TO_WEEK = 604800

func ReadInteractionsFromFile(filepath string) (interactions []Interaction, err error) {
	f, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    csvr := csv.NewReader(f)
    _, err = csvr.Read() // read in header
    if err != nil {
    	return nil, err
    }
    for {
        row, err := csvr.Read()
        if err != nil {
        	if err == io.EOF {
                err = nil
                return interactions, err
            }
            return nil, err
        }
        drugA := row[0]
        drugB := row[1]
        goodBad := row[2]
        adverse := goodBad == "bad"
        time, err := strconv.Atoi(row[3])
        if err != nil {
            fmt.Println("Error occurred parsing integer from csv.")
            return nil, err
        }
        unit := row[4]
        if unit != "sec" {
            time = convertToSeconds(time, unit)
        }
		interaction := Interaction{DrugA: drugA, DrugB: drugB, Adverse: adverse, Time: time}
		interactions = append(interactions, interaction)
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
    }
    return timeInSeconds
}