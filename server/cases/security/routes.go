package security

import (
	"github.com/Quero-Freela/system/server/cases/security/handlers"
	"github.com/gorilla/mux"
)

func UseRoutes(router *mux.Router) {
	// facebook
	router.HandleFunc("/login/facebook", handlers.HandleFacebookLogin).Methods("GET")
	router.HandleFunc("/login/facebook", handlers.HandleFacebookCallback).Methods("POST")

	// github
	router.HandleFunc("/login/github", handlers.HandleGHLogin).Methods("GET")
	router.HandleFunc("/login/github", handlers.HandleGHCallback).Methods("POST")

	// google
	router.HandleFunc("/login/google", handlers.HandleGoogleLogin).Methods("GET")
	router.HandleFunc("/login/google", handlers.HandleGoogleCallback).Methods("POST")
}
