package main

import (
	"github.com/norbertruff/go-graphql/graphql"
	"github.com/norbertruff/go-graphql/graphql/generated/gqlgen"
	"github.com/norbertruff/go-graphql/internal/pkg/auth"
	_ "github.com/norbertruff/go-graphql/internal/pkg/auth"
	database "github.com/norbertruff/go-graphql/internal/pkg/db/db_driver"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	//database.Migrate()
	server := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &graphql.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
