package main

import (
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/dsogari/barber-shop/graph"
	"github.com/dsogari/barber-shop/orm"
	"github.com/dsogari/barber-shop/rest"
)

func main() {
	dbFilename := flag.String("database", "test.db", "Path to the database file")
	restPort := flag.String("rest_port", "8080", "Port to serve the REST API")
	graphqlPort := flag.String("graphql_port", "8081", "Port to serve the GraphQL API")
	flag.Parse()

	orm.SetupDatabase(*dbFilename)
	restServer := rest.SetupServer()
	/*graphServer := */ graph.SetupServer()

	var wg sync.WaitGroup

	// Listen and serve REST
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(restServer.Run(":" + *restPort))
	}()

	// Listen and serve GraphQL
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", *graphqlPort)
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(http.ListenAndServe(":"+*graphqlPort, nil))
	}()

	wg.Wait()
}
