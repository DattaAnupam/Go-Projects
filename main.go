package main

import (
	"database/sql"
	"go-with-sql/pkg/services"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Note
	// To build and run please replace <Password> with your MySql Root password
	dsn := "root:<Password>@tcp(127.0.0.1:3306)/mytestdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	services.AddValuesToDb(db)
	services.RetrieveValuesFromDb(db)
}
