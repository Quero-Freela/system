package handlers

import (
	"fmt"
	"github.com/Quero-Freela/system/server/cases/security/repositories"
	"github.com/Quero-Freela/system/server/cases/security/services"
	"github.com/Quero-Freela/system/server/cases/security/services/providers"
	"github.com/Quero-Freela/system/server/exceptions"
	"github.com/Quero-Freela/system/server/utils"
	"golang.org/x/oauth2"
	"net/http"
)

func HandleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	var baseURI = fmt.Sprintf("%s://%s", r.URL.Scheme, r.URL.Host)

	var cfg, err = providers.GetFbConfig(baseURI)

	if err != nil {
		services.GoToSignErrorPage(w, r, err)
		return
	}

	callbackURL := r.URL.Query().Get("return_url")

	if callbackURL == "" {
		services.GoToSignErrorPage(w, r, exceptions.NewUnauthorizedError("return_url is required"))
		return
	}

	attempt, err := repositories.CreateAttemptSign(string(utils.FacebookAPI), callbackURL)

	if err != nil {
		services.GoToSignErrorPage(w, r, err)
		return
	}

	queryStringOption := oauth2.SetAuthURLParam("return_url", callbackURL)
	url := cfg.AuthCodeURL(attempt.StateRequest, queryStringOption)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
