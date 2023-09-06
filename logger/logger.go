package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Init() {
	logger = logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	if envLevel, present := os.LookupEnv("NOMESS_LOG_LEVEL"); present {
		parsedLevel, err := logrus.ParseLevel(envLevel)
		if err != nil {
			panic(err)
		}

		logger.SetLevel(parsedLevel)
	} else {
		logger.SetLevel(logrus.ErrorLevel)
	}
}

func Panic(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Info(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Debug(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Trace(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}
