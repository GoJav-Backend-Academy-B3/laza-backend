package model

type Brand struct {
	Id      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	LogoUrl string `json:"logo_url,omitempty"`
}
