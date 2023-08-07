package response

type RawData struct {
	Description     string `json:"description"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
	Username        string `json:"username"`
}

type Data struct {
	RawData           RawData `json:"RawData"`
	Provider          string  `json:"Provider"`
	Email             string  `json:"Email"`
	Name              string  `json:"Name"`
	FirstName         string  `json:"FirstName"`
	LastName          string  `json:"LastName"`
	NickName          string  `json:"NickName"`
	Description       string  `json:"Description"`
	UserID            string  `json:"UserID"`
	AvatarURL         string  `json:"AvatarURL"`
	Location          string  `json:"Location"`
	AccessToken       string  `json:"AccessToken"`
	AccessTokenSecret string  `json:"AccessTokenSecret"`
	RefreshToken      string  `json:"RefreshToken"`
	ExpiresAt         string  `json:"ExpiresAt"`
	IDToken           string  `json:"IDToken"`
}

type TwitterResponse struct {
	Status  string `json:"status"`
	IsError bool   `json:"isError"`
	Data    Data   `json:"data"`
}
