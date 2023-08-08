package helper

import "github.com/gorilla/sessions"

func GetStore() *sessions.CookieStore {
	key := "_gothic_session"
	maxAge := 60 // 30 days
	// isProd := true

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	// store.Options.Secure = isProd
	return store
}
