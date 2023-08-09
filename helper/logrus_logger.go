package helper

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

func GetLogrusLogger() *logrus.Logger {
	return logger
}
