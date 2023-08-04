package facebook_auth

type FacebookAuthUsecase interface {
	Execute() (err error)
}
