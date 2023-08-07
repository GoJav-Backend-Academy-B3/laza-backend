package user

type ExistsUsername interface {
	ExistsUsername(username string) bool
}
