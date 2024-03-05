package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Quero-Freela/system/server/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"time"
)

type GitHubUser struct {
	Login  string `json:"login"`
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

func GetGHConfig(baseURI string) (*oauth2.Config, error) {
	secret, err := utils.GetSecret[Secret](utils.GithubAPI)

	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		RedirectURL:  baseURI + "/login/github",
		ClientID:     secret.Id,
		ClientSecret: secret.Secret,
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint,
	}, nil
}

func GetGHUserInfo(code string) (*GitHubUser, error) {
	token, err := GetGHAccessToken(code)
	if err != nil {
		return &GitHubUser{}, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := GetGHOauthUser(token)
	if err != nil {
		return &GitHubUser{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	return response, nil
}

func GetGHAccessToken(code string) (string, error) {
	client := ghHTTPClient()
	secret, err := utils.GetSecret[Secret](utils.GithubAPI)

	if err != nil {
		return "", err
	}

	data, err := json.Marshal(map[string]interface{}{
		"client_id":     secret.Id,
		"client_secret": secret.Secret,
		"code":          code,
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req) //nolint:bodyclose

	if err != nil {
		return "", err
	}
	defer utils.Closer(resp.Body.Close)

	var body map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", err
	}
	if body["error"] != "" {
		return "", fmt.Errorf("error in response body: %s", err.Error())
	}
	return body["access_token"], nil
}

func GetGHOauthUser(accessToken string) (*GitHubUser, error) {
	client := ghHTTPClient()
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req) //nolint:bodyclose
	if err != nil {
		return nil, err
	}
	defer utils.Closer(resp.Body.Close)

	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	email, err := GetGHEmail(accessToken)
	if err != nil {
		return nil, err
	}
	user.Email = email
	return &user, nil
}

func GetGHEmail(accessToken string) (string, error) {
	client := ghHTTPClient()
	req, err := http.NewRequest("GET", "https://api.github.com/user/public_emails", nil)
	if err != nil {
		return "", err
	}
	req.Close = true
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req) //nolint:bodyclose
	if err != nil {
		return "", err
	}
	defer utils.Closer(resp.Body.Close)

	var emails []struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return "", err
	}
	if len(emails) == 0 {
		return "", nil
	}
	return emails[0].Email, nil
}

func ghHTTPClient() *http.Client {
	return &http.Client{Timeout: 5 * time.Second}
}
