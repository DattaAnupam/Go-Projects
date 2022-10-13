package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dbtable struct {
	id        int
	firstname string
	lastname  string
}

func main() {
	db, err := sql.Open("mysql", "root:MyPass123!mYsql@tcp(127.0.0.1:3306)/mytestdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM table1")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var dbtable1 dbtable
		err = results.Scan(&dbtable1.id, &dbtable1.firstname, &dbtable1.lastname)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Id: %d, Firstname: %s, Lastname: %s\n", dbtable1.id, dbtable1.firstname, dbtable1.lastname)
	}

	fmt.Println("Success!!!Value Received")
	// insert, err := db.Query("INSERT INTO table1 VALUES(1, 'Anupam', 'Datta')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	// fmt.Println("Success!!!Value Added")
}
