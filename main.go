package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dsogari/barber-shop/graph/generated"
	"github.com/dsogari/barber-shop/graph/resolvers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbFilename string
var serverPort string

func SetupFlags() {
	flag.StringVar(&dbFilename, "db_filename", "data.db", "Path to the database file")
	flag.StringVar(&serverPort, "server_port", "8080", "Port to serve the GraphQL API")
	flag.Parse()
}

func SetupDatabase() {
	log.Printf("Opening database: %s", dbFilename)
	dialector := sqlite.Open(dbFilename + "?_foreign_keys=on")
	if db, err := gorm.Open(dialector, &gorm.Config{}); err != nil {
		log.Fatal(err)
	} else {
		resolvers.MigrateSchema(db)
	}
}

func SetupServer() {
	config := generated.Config{Resolvers: &resolvers.Resolver{}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
}

func main() {
	SetupFlags()
	SetupDatabase()
	SetupServer()
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
