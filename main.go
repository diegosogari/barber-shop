package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dsogari/barber-shop/graph/generated"
	"github.com/dsogari/barber-shop/graph/resolvers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Path to the database file
var dbFilename string

// Port to serve the GraphQL API
var serverPort string

// Setup the command-line flags
func setupFlags() {
	flag.StringVar(&dbFilename, "db_filename", "data.db", "Path to the database file")
	flag.StringVar(&serverPort, "server_port", "8080", "Port to serve the GraphQL API")
	flag.Parse()
}

// Setup the database connection
func setupDatabase() {
	log.Printf("Opening database: %s", dbFilename)
	dialector := sqlite.Open(dbFilename + "?_foreign_keys=on")
	if db, err := gorm.Open(dialector, &gorm.Config{}); err != nil {
		log.Fatal(err)
	} else {
		resolvers.MigrateSchema(db)
	}
}

//go:embed react-app/build
var uiFS embed.FS

// Setup the HTTP server
func setupServer() {
	config := generated.Config{Resolvers: &resolvers.Resolver{}}
	http.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(config)))
	http.Handle("/play", playground.Handler("GraphQL playground", "/query"))

	if fs, err := fs.Sub(uiFS, "react-app/build"); err != nil {
		log.Fatal("failed to get ui fs", err)
	} else {
		http.Handle("/", http.FileServer(http.FS(fs)))
	}
}

func main() {
	setupFlags()
	setupDatabase()
	setupServer()
	log.Printf("Connect to http://localhost:%s", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
