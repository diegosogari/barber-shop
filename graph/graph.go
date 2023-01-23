package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dsogari/barber-shop/graph/generated"
	"github.com/dsogari/barber-shop/graph/resolvers"
)

func SetupServer() (srv *handler.Server) {
	config := generated.Config{Resolvers: &resolvers.Resolver{}}
	srv = handler.NewDefaultServer(generated.NewExecutableSchema(config))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	return
}
