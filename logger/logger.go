package logger

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

type Logger struct {
	log   *log.Logger
	level int
}

const (
	DEBUG     = 8
	LOG       = 7
	INFO      = 6
	WARN      = 5
	ERR       = 4
	CRIT      = 3
	ALERT     = 2
	EMERGENCY = 1
)

var logger *Logger

func Init() *Logger {
	if logger != nil {
		return logger
	}

	var buf bytes.Buffer

	if envLevel, present := os.LookupEnv("NOMESS_LOG_LEVEL"); present {
		intLevel, err := strconv.Atoi(envLevel)
		if err != nil {
			return &Logger{log.New(&buf, "[-]", log.Lshortfile|log.Ltime), intLevel}
		}
	}

	logger = &Logger{log.New(&buf, "[ERR]", log.Lshortfile|log.Ltime), ERR}

	return logger
}

func Debug(msg string) {
	if logger.level == DEBUG {
		logger.log.SetPrefix("[DEBUG]")
		logger.log.Output(2, msg)
	}
}

func Log(msg string) {
	if logger.level >= LOG {
		logger.log.SetPrefix("[LOG]")
		logger.log.Output(2, msg)
	}
}

func Info(msg string) {
	if logger.level >= INFO {
		logger.log.SetPrefix("[INFO]")
		logger.log.Output(2, msg)
	}
}

func Warn(msg string) {
	if logger.level >= WARN {
		logger.log.SetPrefix("[WARN]")
		logger.log.Output(2, msg)
	}
}

func Err(msg string) {
	if logger.level >= ERR {
		logger.log.SetPrefix("[ERR]")
		logger.log.Output(2, msg)
	}
}

func Crit(msg string) {
	if logger.level >= CRIT {
		logger.log.SetPrefix("[CRIT]")
		logger.log.Output(2, msg)
	}
}

func Alert(msg string) {
	if logger.level >= ALERT {
		logger.log.SetPrefix("[ALERT]")
		logger.log.Output(2, msg)
		os.Exit(1)
	}
}

func Emergency(msg string) {
	if logger.level >= EMERGENCY {
		logger.log.SetPrefix("[EMERGENCY]")
		logger.log.Output(2, msg)
		panic(msg)
	}
}
