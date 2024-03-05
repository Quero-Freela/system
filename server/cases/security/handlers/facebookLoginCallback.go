package handlers

import (
	"errors"
	"fmt"
	"github.com/Quero-Freela/system/server/cases/security/repositories"
	"github.com/Quero-Freela/system/server/cases/security/services"
	"github.com/Quero-Freela/system/server/cases/security/services/providers"
	"github.com/Quero-Freela/system/server/exceptions"
	"github.com/Quero-Freela/system/server/utils"
	"net/http"
)

func HandleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")

	if state == "" {
		ReHandler(w, r, HandleFacebookLogin)
		return
	}

	attempt, err := repositories.FindAttemptByState(string(utils.FacebookAPI), state)

	if err != nil || attempt.StateRequest != state {
		services.GoToSignErrorPage(w, r, exceptions.NewUnauthorizedError("invalid state"))
		return
	}

	code := r.FormValue("code")
	fbUser, errFb := providers.GetUserInfoFacebook(code, utils.GetBaseUrl(r))
	user, errUsr := repositories.FindUserByEmail(fbUser.Email)

	if errFb != nil || errUsr != nil {
		services.GoToUserErrorPage(w, r, errors.Join(errFb, errUsr), fbUser.Email, code, attempt)
		return
	}

	token, errTk := services.CreateToken(user)

	if errTk != nil || token == nil {
		services.GoToSignErrorPage(w, r, exceptions.NewUnauthorizedError("cannot retrieve token"))
		return
	}

	err = repositories.UpdateAttemptUserToken(attempt, user, code, *token)

	if err != nil {
		utils.LogError(err)
	}

	appURL := fmt.Sprintf("%s?token=%s", utils.GetAppUrl(r), *token)
	http.Redirect(w, r, appURL, http.StatusTemporaryRedirect)
}
