package appcontext

import (
	"os"

	"github.com/jmoiron/sqlx"
	// Dependency of sqlx
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	testhook "github.com/sirupsen/logrus/hooks/test"
)

type appContext struct {
	db     *sqlx.DB
	logger *logrus.Logger
}

var context *appContext

func Initialize() {
	if context != nil {
		return
	}
	db, err := getDBConnection()
	if err != nil {
		panic(err)
	}

	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		panic(err)
	}

	context = &appContext{db: db, logger: createLogger(os.Getenv("ENVIRONMENT"), level)}
}

func createLogger(environment string, logLevel logrus.Level) *logrus.Logger {
	if environment == "test" {
		nullLogger, _ := testhook.NewNullLogger()
		return nullLogger
	}

	return &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logLevel,
	}
}

func GetDB() *sqlx.DB {
	return context.db
}

func LogInfo(message, info string) {
	context.logger.WithFields(logrus.Fields{"info": info}).Info(message)
}

func LogDebug(message, info string) {
	context.logger.WithFields(logrus.Fields{"info": info}).Debug(message)
}

func LogError(message, info string) {
	context.logger.WithFields(logrus.Fields{"info": info}).Error(message)
}
