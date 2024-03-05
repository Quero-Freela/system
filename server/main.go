//go:build !lambda && amd64

// nolint
// Description: Main entry point for local server
package main

import (
	"github.com/Quero-Freela/system/server/locals"
	m "github.com/Quero-Freela/system/server/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8000"
	}

	if port[0] != ':' {
		port = ":" + port
	}

	log.Fatal(http.ListenAndServe(port, locals.AddRoutes(m.Middlewares(router))))
}
