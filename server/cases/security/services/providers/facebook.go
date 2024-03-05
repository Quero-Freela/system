package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Quero-Freela/system/server/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"io"
	"net/http"
)

type FacebookUser struct {
	Name  string `json:"first_name"`
	Email string `json:"email"`
}

func GetFbConfig(baseURI string) (*oauth2.Config, error) {
	secret, err := utils.GetSecret[Secret](utils.FacebookAPI)

	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		RedirectURL:  baseURI + "/login/facebook",
		ClientID:     secret.Id,
		ClientSecret: secret.Secret,
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}, nil
}

func GetUserInfoFacebook(code, baseURI string) (*FacebookUser, error) {
	var cfg, err = GetFbConfig(baseURI)

	if err != nil {
		return nil, fmt.Errorf("failed getting facebook config: %s", err.Error())
	}

	token, err := cfg.Exchange(context.TODO(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	url := fmt.Sprintf("https://graph.facebook.com/me?fields=first_name,email&access_token=%s", token.AccessToken)

	response, err := http.Get(url) //nolint:bodyclose,gosec
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer utils.Closer(response.Body.Close)
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var fbUser *FacebookUser
	_ = json.Unmarshal(contents, &fbUser)

	return fbUser, nil
}
