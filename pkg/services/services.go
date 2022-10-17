package services

import (
	"database/sql"
	"fmt"
)

type Dbtable struct {
	Id        int
	Firstname string
	Lastname  string
}

func AddValuesToDb(db *sql.DB) {
	insert, err := db.Query("INSERT INTO table1 VALUES(1, 'Anupam', 'Datta')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Success!!!Value Added")
}

func RetrieveValuesFromDb(db *sql.DB) {
	results, err := db.Query("SELECT * FROM table1")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var dbtable1 Dbtable
		err = results.Scan(&dbtable1.Id, &dbtable1.Firstname, &dbtable1.Lastname)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Id: %d, Firstname: %s, Lastname: %s\n", dbtable1.Id, dbtable1.Firstname, dbtable1.Lastname)
	}

	fmt.Println("Success!!!Value Received")
}
