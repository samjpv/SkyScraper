package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

const apiPort = "8080"
const mysqlPort = "1433"

func main() {

	// connect to the mysql database

	db, err := sql.Open("sqlserver", "Data Source='localhost, 1433';Initial Catalog=DEFAULT;User ID=sa;Password=[MSSQLServer!]")
	defer db.Close()
	if err != nil {
		log.Printf("Error")
		panic(err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Printf("Connection successfully pinged")
	}
	log.Printf("DB opened without error!")

	ins, err := db.Query("INSERT INTO People VALUES ('hello', 'there')")
	defer ins.Close()
	if err != nil {
		log.Printf("Error")
		panic(err.Error())
	}
	log.Printf("Successful query")

	log.Printf("Moving on")
	people, err := db.Query("SELECT * FROM People")
	defer people.Close()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Query completed without error!")

	var (
		firstName  string
		secondName string
	)
	for people.Next() {
		err := people.Scan(&firstName, &secondName)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("FirstName: %s", firstName)
		log.Printf("SecondName: %s", secondName)
	}

	// var (
	// 	sqlversion string
	// )
	// rows, err := db.Query("select @@version")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for rows.Next() {
	// 	err := rows.Scan(&sqlversion)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(sqlversion)
	// }

	// original code
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = apiPort
	// }

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("Database access", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for interactive access", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

}
