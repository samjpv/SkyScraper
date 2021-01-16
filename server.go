package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"test/SkyScraper/graph"
	"test/SkyScraper/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/denisenkom/go-mssqldb"
)

const apiPort = "8080"
const mysqlPort = "1433"

func main() {

	// connect to the mysql database

	db, err := sql.Open("sqlserver", "Data Source='localhost, 1433';Initial Catalog=DEFAULT;User ID=sa;Password=[MSSQLServer!]")
	defer db.Close()

	if err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = apiPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("Database access", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for interactive access", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
