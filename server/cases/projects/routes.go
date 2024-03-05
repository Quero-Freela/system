package projects

import (
	"github.com/Quero-Freela/system/server/cases/projects/handlers"
	"github.com/gorilla/mux"
)

func UseRoutes(router *mux.Router) {
	router.HandleFunc("/projects", handlers.List).Methods("GET")
}
