package creditcard

type IsExistsCcAction interface {
	IsExistsCc(userId uint64, ccNumber string) (bool, error)
}
