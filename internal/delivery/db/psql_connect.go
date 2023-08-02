package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/domain/contract"

	_ "github.com/lib/pq"
)

var psqldb PsqlDB

type PsqlDB struct {
	Dbs *sql.DB
	Trx *sql.Tx
}

func GetPsqlConnection() contract.Dbs {
	var nilDB PsqlDB
	if psqldb != nilDB {
		return &psqldb
	} else {
		psqldb.OpenConnection()
		return &psqldb
	}
}

func (d *PsqlDB) OpenConnection() {
	appConfig := config.AppConfig()

	user := appConfig.PSQL_USER
	pass := appConfig.PSQL_PASS
	host := appConfig.PSQL_HOST
	port := appConfig.PSQL_PORT
	dbname := appConfig.PSQL_DBNAME
	timezone := appConfig.PSQL_TIMEZONE

	connString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s"
	db_, err := sql.Open("postgres", fmt.Sprintf(connString, host, user, pass, dbname, port, timezone))
	if err != nil {
		panic(err)
	}

	db_.SetConnMaxIdleTime(10 * time.Minute)
	db_.SetConnMaxLifetime(12 * time.Hour)
	db_.SetMaxIdleConns(10)
	db_.SetMaxOpenConns(100)

	d.Dbs = db_
}

func (d *PsqlDB) StartTrx() {
	trx, err := psqldb.Dbs.Begin()
	if err != nil {
		return
	}
	psqldb.Trx = trx
}

func (d *PsqlDB) DoneTrx(err error) {
	if err != nil {
		psqldb.Trx.Rollback()
		psqldb.Trx = &sql.Tx{}
	} else {
		psqldb.Trx.Commit()
		psqldb.Trx = &sql.Tx{}
	}
}
