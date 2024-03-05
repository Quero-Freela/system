package locals

import (
	"github.com/Quero-Freela/system/server/cases/projects"
	"github.com/Quero-Freela/system/server/cases/security"
	"github.com/gorilla/mux"
)

func AddRoutes(router *mux.Router) *mux.Router {
	security.UseRoutes(router)

	sub := router.PathPrefix("/api").Subrouter()

	projects.UseRoutes(sub)

	return router
}
