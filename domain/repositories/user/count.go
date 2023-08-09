package user

type Count interface {
	Count() (int64, error)
}
