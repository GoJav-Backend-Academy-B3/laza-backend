package requests

type ReviewRequest struct {
	Comment string  `validate:"required" json:"comment,omitempty"`
	Rating  float32 `validate:"required" json:"rating,omitempty"`
}
