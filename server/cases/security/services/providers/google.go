package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Quero-Freela/system/server/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
)

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func GetGoogleOauthConfig(baseURI string) (*oauth2.Config, error) {
	secret, err := utils.GetSecret[Secret](utils.GoogleAPI)

	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		RedirectURL:  baseURI + "/login/google",
		ClientID:     secret.Id,
		ClientSecret: secret.Secret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: google.Endpoint,
	}, nil
}

func GetUserInfoGoogle(code, baseURI string) (*GoogleUser, error) {
	googleOauthConfig, err := GetGoogleOauthConfig(baseURI)

	if err != nil {
		return &GoogleUser{}, err
	}

	token, err := googleOauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		return &GoogleUser{}, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	reqURL := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", token.AccessToken)

	response, err := http.Get(reqURL) //nolint:bodyclose,gosec
	if err != nil {
		return &GoogleUser{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer utils.Closer(response.Body.Close)
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return &GoogleUser{}, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var user *GoogleUser
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return &GoogleUser{}, fmt.Errorf("failed parsing response body: %s", err.Error())
	}

	return user, nil
}
