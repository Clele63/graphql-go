package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"workbench/graphql-app/db"
	"workbench/graphql-app/graph/connect"
	"workbench/graphql-app/graph/exec"
	"workbench/graphql-app/graph/resolver"
	"workbench/graphql-app/middlewares"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "6000"
const defaultAppEnv = "dev"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = defaultAppEnv
	}

	// connect.InitDB()
	database, err := db.SetupDB()
	if err != nil {
		log.Fatalf("failed to setup the DB: %v", err)
	}
	err = database.InitDB()
	if err != nil {
		log.Panicf("failed to initialize the DB: %v", err)
	}

	resolver := &resolver.Resolver{
		Queries: connect.GetQueries(),
	}

	router := chi.NewRouter()

	router.Use(middlewares.JwtMiddleware())

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Bearer"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	srv := handler.New(exec.NewExecutableSchema(exec.Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	if appEnv == "prod" {
		fs := http.FileServer(http.Dir(filepath.Join("client", "build")))
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join("client", "build", r.URL.Path)
			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				http.ServeFile(w, r, filepath.Join("client", "build", "index.html"))
				return
			}
			fs.ServeHTTP(w, r)
		})
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
