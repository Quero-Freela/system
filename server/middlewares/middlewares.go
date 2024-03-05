package middlewares

import (
	"github.com/gorilla/mux"
)

func Middlewares(router *mux.Router) *mux.Router {
	//router.Use(AuthenticationMiddleware)
	return router
}
