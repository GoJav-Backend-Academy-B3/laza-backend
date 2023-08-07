package provider

import (
	d "github.com/phincon-backend/laza/domain/handlers"
	h "github.com/phincon-backend/laza/internal/handler/twitter_auth"
)

func NewtwitterAuthHandler() d.HandlerInterface {

	return h.NewtwitterAuthHandler("/auth", "/auth/twitter/callback")
}
