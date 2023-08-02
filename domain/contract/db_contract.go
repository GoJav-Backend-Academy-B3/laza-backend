package contract

type Dbs interface {
	OpenConnection()
	StartTrx()
	DoneTrx(err error)
}
