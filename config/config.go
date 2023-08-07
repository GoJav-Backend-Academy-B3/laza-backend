package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type configStruct struct {
	PSQL_USER         string
	PSQL_PASS         string
	PSQL_HOST         string
	PSQL_PORT         string
	PSQL_DBNAME       string
	PSQL_TIMEZONE     string
	PSQL_SEARCH_PATH  string
	PSQL_TimeoutQuick time.Duration
	PSQL_TimeoutMid   time.Duration
	PSQL_TimeoutSlow  time.Duration
}

var appConfig = new(configStruct)

func init() {
	godotenv.Load("d:/phincon/laza-backend/.env")
	appConfig.PSQL_USER = os.Getenv("PSQL_USER")
	appConfig.PSQL_PASS = os.Getenv("PSQL_PASS")
	appConfig.PSQL_HOST = os.Getenv("PSQL_HOST")
	appConfig.PSQL_PORT = os.Getenv("PSQL_PORT")
	appConfig.PSQL_DBNAME = os.Getenv("PSQL_DBNAME")
	appConfig.PSQL_TIMEZONE = os.Getenv("PSQL_TIMEZONE")
	appConfig.PSQL_SEARCH_PATH = os.Getenv("PSQL_SEARCH_PATH")

	durationQuick, err := strconv.Atoi(os.Getenv("PSQL_TIMEOUT_1"))
	if err != nil {
		panic(err)
	}
	appConfig.PSQL_TimeoutQuick = time.Duration(durationQuick) * time.Second

	durationMid, err := strconv.Atoi(os.Getenv("PSQL_TIMEOUT_2"))
	if err != nil {
		panic(err)
	}
	appConfig.PSQL_TimeoutMid = time.Duration(durationMid) * time.Second

	durationSlow, err := strconv.Atoi(os.Getenv("PSQL_TIMEOUT_3"))
	if err != nil {
		panic(err)
	}
	appConfig.PSQL_TimeoutSlow = time.Duration(durationSlow) * time.Second
}

func AppConfig() configStruct {
	return *appConfig
}
