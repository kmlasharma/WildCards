package ddi

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/kmlasharma/WildCards/pkgs/logger"
	"github.com/kmlasharma/WildCards/pkgs/pml"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type Database struct {
	conn *sql.DB
}

type Interaction struct {
	DrugA   string
	DrugB   string
	Adverse bool
	Time    int
}

func NewDatabase() *Database {
	conn, err := sql.Open("sqlite3", "./dinto.db")
	if err != nil {
		logger.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	db := &Database{conn: conn}
	db.createTableIfNotExists()
	db.Clear()
	return db
}

func (db *Database) Populate(interactions []Interaction) error {
	strs := []string{"INSERT INTO interactions ('DrugA', 'DrugB', 'Adverse', 'Time') VALUES"}
	for _, interaction := range interactions {
		var integer = 0
		if interaction.Adverse {
			integer = 1
		}
		line := fmt.Sprintf("('%s', '%s', '%d', '%d'),", interaction.DrugA, interaction.DrugB, integer, interaction.Time)
		strs = append(strs, line)
	}
	concatenatedString := strings.Join(strs, "\n")
	command := strings.TrimSuffix(concatenatedString, ",") + ";" // Tidy up by removing last comma, and add semi colon
	_, err := db.conn.Exec(command)
	return err
}

func (db *Database) PopulateFromFile(filepath string) error {
	interactions, err := readInteractionsFromFile(filepath)
	if err != nil {
		return err
	}

	err = db.Populate(interactions)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) FindInteractions(drugs []string) (interactions []Interaction, err error) {
	template := "SELECT * FROM interactions WHERE DrugA IN ('%s') AND DrugB IN ('%s');"
	drugsString := strings.Join(drugs, "','")
	query := fmt.Sprintf(template, drugsString, drugsString)
	fmt.Println("Query:", query)
	rows, err := db.conn.Query(query)
	var drugA, drugB string
	var adverse bool
	var time int
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&drugA, &drugB, &adverse, &time)
			if err == nil {
				interaction := Interaction{
					DrugA:   drugA,
					DrugB:   drugB,
					Adverse: adverse,
					Time:    time,
				}
				interactions = append(interactions, interaction)
			}
		}
		err = rows.Err()
	}
	rows.Close()
	return
}

func (db *Database) FindActiveInteractionsForPairs(pairs []pml.DrugPair) (interactions []Interaction, err error) {
	for _, pair := range pairs {
		interaction, err := db.FindInteraction(pair.DrugA, pair.DrugB)
		if err == nil && interaction.Time > int(pair.Delay) {
			interactions = append(interactions, interaction)
		}
	}
	return
}

func (db *Database) FindInteraction(drugA, drugB string) (Interaction, error) {
	interactions, err := db.FindInteractions([]string{drugA, drugB})
	if err == nil && len(interactions) > 0 {
		return interactions[0], nil
	} else {
		return Interaction{}, errors.New("Not Found")
	}
}

func (db *Database) Clear() {
	db.conn.Exec("DELETE from interactions")
}

func (db *Database) Close() {
	db.conn.Close()
}

func (db *Database) createTableIfNotExists() {
	command := `
  CREATE TABLE IF NOT EXISTS interactions(
    DrugA TEXT,
    DrugB TEXT,
    Adverse INTEGER,
    Time INTEGER
  );`
	_, err := db.conn.Exec(command)
	if err != nil {
		logger.Fatal(err)
	}
}
