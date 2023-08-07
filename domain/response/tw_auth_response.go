package response

type RawData struct {
	ContributorsEnabled bool   `json:"contributors_enabled"`
	CreatedAt           string `json:"created_at"`
	DefaultProfile      bool   `json:"default_profile"`
	DefaultProfileImage bool   `json:"default_profile_image"`
	Description         string `json:"description"`
	Email               string `json:"email"`
	Entities            struct {
		Description struct {
			URLs []interface{} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	FavouritesCount                int           `json:"favourites_count"`
	FollowRequestSent              bool          `json:"follow_request_sent"`
	FollowersCount                 int           `json:"followers_count"`
	Following                      bool          `json:"following"`
	FriendsCount                   int           `json:"friends_count"`
	GeoEnabled                     bool          `json:"geo_enabled"`
	HasExtendedProfile             bool          `json:"has_extended_profile"`
	ID                             int64         `json:"id"`
	IDStr                          string        `json:"id_str"`
	IsTranslationEnabled           bool          `json:"is_translation_enabled"`
	IsTranslator                   bool          `json:"is_translator"`
	Lang                           string        `json:"lang"`
	ListedCount                    int           `json:"listed_count"`
	Location                       string        `json:"location"`
	Name                           string        `json:"name"`
	NeedsPhoneVerification         bool          `json:"needs_phone_verification"`
	Notifications                  bool          `json:"notifications"`
	ProfileBackgroundColor         string        `json:"profile_background_color"`
	ProfileBackgroundImageURL      string        `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string        `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool          `json:"profile_background_tile"`
	ProfileImageURL                string        `json:"profile_image_url"`
	ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
	ProfileLinkColor               string        `json:"profile_link_color"`
	ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string        `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
	Protected                      bool          `json:"protected"`
	ScreenName                     string        `json:"screen_name"`
	StatusesCount                  int           `json:"statuses_count"`
	Suspended                      bool          `json:"suspended"`
	TimeZone                       string        `json:"time_zone"`
	TranslatorType                 string        `json:"translator_type"`
	URL                            string        `json:"url"`
	UTCOffset                      string        `json:"utc_offset"`
	Verified                       bool          `json:"verified"`
	WithheldInCountries            []interface{} `json:"withheld_in_countries"`
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
