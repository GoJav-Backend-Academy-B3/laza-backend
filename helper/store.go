package helper

import "github.com/gorilla/sessions"

func GetStore() *sessions.CookieStore {
	key := "SAgn9qi1Zt3OV8xfDaMRiBQvK"
	maxAge := 86400 * 30 // 30 days
	// isProd := true

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	// store.Options.Secure = isProd
	return store
}
