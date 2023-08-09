package helper

import (
	"github.com/gorilla/sessions"
)

var store = new(sessions.CookieStore)

func init() {
	key := "_gothic_session"
	maxAge := 86400 * 30 // 30 days
	// isProd := true

	store = sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
}

func GetStore() *sessions.CookieStore {
	return store
}
