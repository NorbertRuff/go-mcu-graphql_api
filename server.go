package main

import (
	"github.com/joho/godotenv"
	"github.com/norbertruff/go-graphql/graphql/resolvers"
	gqlgen "github.com/norbertruff/go-graphql/graphql/server"
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
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	database.InitDB()
	//database.Migrate()
	server := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{Resolvers: &resolvers.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
