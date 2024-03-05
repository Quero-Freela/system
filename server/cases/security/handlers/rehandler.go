package handlers

import (
	"github.com/Quero-Freela/system/server/cases/security/services"
	"net/http"
)

func ReHandler(w http.ResponseWriter, r *http.Request, handle func(w http.ResponseWriter, r *http.Request)) {
	newRequest, err := http.NewRequest("GET", r.URL.String(), nil)

	if err != nil {
		services.GoToSignErrorPage(w, r, err)
		return
	}

	newRequest.URL.RawQuery = r.URL.RawQuery
	handle(w, newRequest)
}
