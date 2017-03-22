package dinto

import (
	"fmt"
	// "github.com/kmlasharma/WildCards/pkgs/logger"
	// "strings"
	"encoding/csv"
	"os"
	"io"
)


func ReadInteractionsFromFile(filepath string) (interactions []Interaction, err error) {
	// s := Interaction{DrugA: "Sean", DrugB: "5", Adverse: true}
	// fmt.Println(s)
	// return s
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
        // time := row[3]
        // unit := row[4]
		interaction := Interaction{DrugA: drugA, DrugB: drugB, Adverse: adverse}
		interactions = append(interactions, interaction)
    }
    f.Close()
    fmt.Println(interactions)
	return interactions, nil

}