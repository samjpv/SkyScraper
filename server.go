package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

func main() {

	db, err := sql.Open("mysql", "sa:[MSSQLServer!]@tcp(localhost:1433)/DEFAULT")

	if err != nil {
		log.Printf("Error")
		panic(err.Error())
	}
	log.Printf("DB opened without error!")
	defer db.Close()

	// perform a db.Query insert
	log.Printf("Moving on")
	insert, err := db.Query("INSERT INTO dbo.Persons VALUES ('firstname', 'lastname')")
	log.Printf("Query completed")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

	// original code
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("Database access", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for interactive access", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

}
