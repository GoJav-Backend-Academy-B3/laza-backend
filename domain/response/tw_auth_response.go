package response

import (
	"time"

	"github.com/markbates/goth"
)

type RawData struct {
	Description     string `json:"description"`
	Id              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageUrl string `json:"profile_image_url"`
	UserName        string `json:"username"`
}

type TwitterUser struct {
	RawData           RawData
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}

func (t *TwitterUser) FillEntity(usr goth.User) {
	t.RawData.Description = usr.RawData["description"].(string)
	t.RawData.Id = usr.RawData["id"].(string)
	t.RawData.Name = usr.RawData["name"].(string)
	t.RawData.ProfileImageUrl = usr.RawData["profile_image_url"].(string)
	t.RawData.UserName = usr.RawData["username"].(string)
	t.Provider = usr.Provider
	t.Email = usr.Email
	t.Name = usr.Name
	t.FirstName = usr.FirstName
	t.LastName = usr.LastName
	t.NickName = usr.NickName
	t.Description = usr.Description
	t.UserID = usr.UserID
	t.AvatarURL = usr.AvatarURL
	t.Location = usr.Location
	t.AccessToken = usr.AccessToken
	t.AccessTokenSecret = usr.AccessTokenSecret
	t.RefreshToken = usr.RefreshToken
	t.ExpiresAt = usr.ExpiresAt
	t.IDToken = usr.IDToken
}
