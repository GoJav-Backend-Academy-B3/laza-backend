package requests

type POSTSizeRequest struct {
	Sizes []string `json:"sizes"`
}

type PUTSizeRequest struct {
	Size string `json:"size"`
}
