package db

type Dbs interface {
	OpenConnection()
	StartTrx()
	DoneTrx(err error)
}
