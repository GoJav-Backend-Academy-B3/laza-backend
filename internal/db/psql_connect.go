package db

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/phincon-backend/laza/config"
	"github.com/phincon-backend/laza/domain/db"

	_ "github.com/lib/pq"
)

var psqldb PsqlDB

type PsqlDB struct {
	Dbs *gorm.DB
	Trx *gorm.DB
}

func GetPostgreSQLConnection() db.Dbs {
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
	searchPath := appConfig.PSQL_SEARCH_PATH

	connString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=%s search_path=%s"
	db_, err := sql.Open("postgres", fmt.Sprintf(connString, host, user, pass, dbname, port, timezone, searchPath))
	if err != nil {
		panic(err)
	}
	db_.SetConnMaxIdleTime(10 * time.Minute)
	db_.SetConnMaxLifetime(12 * time.Hour)
	db_.SetMaxIdleConns(10)
	db_.SetMaxOpenConns(100)

	gormdb, err := gorm.Open(postgres.New(postgres.Config{Conn: db_}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	d.Dbs = gormdb
}

func (d *PsqlDB) StartTrx() {
	db := psqldb.Dbs.Begin()
	psqldb.Trx = db
}

func (d *PsqlDB) DoneTrx(err error) {
	if err != nil {
		psqldb.Trx.Rollback()
		psqldb.Trx = &gorm.DB{}
	} else {
		psqldb.Trx.Commit()
		psqldb.Trx = &gorm.DB{}
	}
}
