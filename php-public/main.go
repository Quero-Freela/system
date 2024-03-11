package main

import (
	"github.com/Quero-Freela/system/server/locals"
	m "github.com/Quero-Freela/system/server/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/schivei/php-go/php" //nolint:typecheck
)

func init() {
	_ = godotenv.Load()

	router := mux.NewRouter()

	locals.AddRoutes(m.Middlewares(router))

	php.Start("querofreela", router)
}

func main() {}
