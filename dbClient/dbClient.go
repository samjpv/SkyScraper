package dbClient

import (
	"database/sql"
	"log"
)

func OpenDb() *sql.DB {
	db, err := sql.Open("sqlserver", "Data Source='localhost, 1433';Initial Catalog=DEFAULT;User ID=sa;Password=[MSSQLServer!]")
	if err != nil {
		log.Printf("Error")
		panic(err.Error())
	}
	return db
}
