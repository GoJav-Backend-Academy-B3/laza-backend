package helper

import (
	"encoding/json"
	"net/http"

	"github.com/phincon-backend/laza/config"
	"golang.org/x/oauth2"
)

type GoogleUserResult struct {
	Id             string
	Email          string
	Verified_email bool
	Name           string
	Given_name     string
	Family_name    string
	Picture        string
	Locale         string
}

func GetGoogleUser(token *oauth2.Token) (*GoogleUserResult, error) {
	// Fetch User Data from google server
	response, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		return nil, err
	}

	// Parse user data JSON Object
	defer response.Body.Close()
	var userBody GoogleUserResult
	err = json.NewDecoder(response.Body).Decode(&userBody)
	if err != nil {
		return nil, err
	}

	return &userBody, nil
}
