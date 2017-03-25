package log

import (
	"github.com/Sirupsen/logrus"
	"os"
)

type Entries map[string]interface{}

var log = logrus.New()

func SetLevel(level string) {
	switch level {
	case "debug":
		log.Level = logrus.DebugLevel
		break
	case "info":
		log.Level = logrus.InfoLevel
		break
	case "warn":
		log.Level = logrus.WarnLevel
		break
	case "error":
		log.Level = logrus.ErrorLevel
		break
	case "fatal":
		log.Level = logrus.FatalLevel
		break
	case "panic":
		log.Level = logrus.PanicLevel
		break
	default:
		log.Level = logrus.InfoLevel
		log.Warn("Log level invalid: ", level)
		break
	}
}

func SetFile(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Formatter = &logrus.JSONFormatter{}
		log.Out = file
	} else {
		log.Warn("Failed to log to file, using default stderr")
	}
}

func New() Entries {
	return Entries{}
}

func Add(name string, value interface{}) Entries {
	return New().Add(name, value)
}

func (entries Entries) Add(name string, value interface{}) Entries {
	entries[name] = value
	return entries
}

func Debug(message string) {
	New().Debug(message)
}

func (entries Entries) Debug(message string) {
	log.WithFields(logrus.Fields(entries)).Debug(message)
}

func Info(message string) {
	New().Info(message)
}

func (entries Entries) Info(message string) {
	log.WithFields(logrus.Fields(entries)).Info(message)
}

func Warn(message string) {
	New().Warn(message)
}

func (entries Entries) Warn(message string) {
	log.WithFields(logrus.Fields(entries)).Warn(message)
}

func Error(message interface{}) {
	New().Error(message)
}

func (entries Entries) Error(message interface{}) {
	log.WithFields(logrus.Fields(entries)).Error(message)
}
