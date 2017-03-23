package dinto

import (
	"database/sql"
	"fmt"
	"github.com/kmlasharma/WildCards/pkgs/logger"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

type Dinto struct {
	db *sql.DB
}

type Interaction struct {
	DrugA   string
	DrugB   string
	Adverse bool
}

func NewDinto() *Dinto {
	db, err := sql.Open("sqlite3", "./dinto.db")
	if err != nil {
		logger.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	dinto := &Dinto{db: db}
	dinto.createTableIfNotExists()
	return dinto
}

func (dinto *Dinto) Populate(interactions []Interaction) error {
	strs := []string{"INSERT INTO interactions ('DrugA', 'DrugB', 'Adverse') VALUES"}
	for _, interaction := range interactions {
		var integer = 0
		if interaction.Adverse {
			integer = 1
		}
		line := fmt.Sprintf("('%s', '%s', '%d'),", interaction.DrugA, interaction.DrugB, integer)
		strs = append(strs, line)
	}
	concatenatedString := strings.Join(strs, "\n")
	command := strings.TrimSuffix(concatenatedString, ",") + ";" // Tidy up by removing last comma, and add semi colon
	_, err := dinto.db.Exec(command)
	return err
}

func (dinto *Dinto) FindInteractions(drugs []string) (interactions []Interaction, err error) {
	template := "SELECT * FROM interactions WHERE DrugA IN ('%s') AND DrugB IN ('%s');"
	drugsString := strings.Join(drugs, "','")
	query := fmt.Sprintf(template, drugsString, drugsString)
	fmt.Println(query)
	rows, err := dinto.db.Query(query)

	//rows, err := dinto.db.Query("SELECT DrugA, DrugB, Adverse FROM interactions WHERE DrugA IN (?, ?) AND DrugB IN ('coke','7up');")
	//fmt.Println(drugsString)
	var drugA, drugB string
	var adverse bool
	if err == nil {
		for rows.Next() {
			err := rows.Scan(&drugA, &drugB, &adverse)
			if err == nil {
				interaction := Interaction{
					DrugA:   drugA,
					DrugB:   drugB,
					Adverse: adverse,
				}
				interactions = append(interactions, interaction)
			}
		}
		err = rows.Err()
	}
	rows.Close()
	return
}

func (dinto *Dinto) Clear() {
	dinto.db.Exec("DELETE from interactions")
}

func (dinto *Dinto) Close() {
	dinto.db.Close()
}

func (dinto *Dinto) createTableIfNotExists() {
	command := `
  CREATE TABLE IF NOT EXISTS interactions(
    DrugA TEXT,
    DrugB TEXT,
    Adverse INTEGER
  );`
	_, err := dinto.db.Exec(command)
	if err != nil {
		logger.Fatal(err)
	}
}
