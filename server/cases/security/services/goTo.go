package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Quero-Freela/system/server/cases/security/repositories"
	"github.com/Quero-Freela/system/server/exceptions"
	"github.com/Quero-Freela/system/server/models"
	"github.com/Quero-Freela/system/server/utils"
	"net/http"
)

func GoToErrorPage(w http.ResponseWriter, r *http.Request, err error) {
	appURL := utils.GetAppUrl(r)
	jsonError, _ := json.Marshal(err)
	base64Error := base64.StdEncoding.EncodeToString(jsonError)
	url := fmt.Sprintf("%s/error?message=%s", appURL, base64Error)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoToSignErrorPage(w http.ResponseWriter, r *http.Request, err error) {
	appURL := utils.GetAppUrl(r)
	jsonError, _ := json.Marshal(err)
	base64Error := base64.StdEncoding.EncodeToString(jsonError)
	url := fmt.Sprintf("%s/sign_error?message=%s", appURL, base64Error)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoToUserErrorPage(w http.ResponseWriter, r *http.Request, err error, email, code string, attempt *models.AttemptSign) {
	if errors.Is(err, &exceptions.NotFoundError{}) {
		e := repositories.UpdateAttemptUserToken(attempt, nil, code, "")

		if e != nil {
			utils.LogError(e)
		} else {
			GoToSignUpPage(w, r, email, attempt)
			return
		}
	}

	GoToSignErrorPage(w, r, err)
}

func GoToSignUpPage(w http.ResponseWriter, r *http.Request, email string, attempt *models.AttemptSign) {
	appURL := utils.GetAppUrl(r)
	url := fmt.Sprintf("%s/signup?email=%s&state=%s&code=%s", appURL, email, attempt.StateRequest, attempt.Code)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
