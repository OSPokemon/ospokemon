package log

import (
	"github.com/Sirupsen/logrus"
)

type Log struct {
	message string
	fields  map[string]interface{}
}

func SetLevel(level string) {
	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		break
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
		break
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
		break
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
		break
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
		break
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
		break
	default:
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warn("Log level invalid: ", level)
		break
	}
}

func New() *Log {
	return &Log{
		fields: make(map[string]interface{}),
	}
}

func Add(name string, value interface{}) *Log {
	return New().Add(name, value)
}

func (log *Log) Add(name string, value interface{}) *Log {
	log.fields[name] = value
	return log
}

func Debug(message string) {
	New().Debug(message)
}

func (log *Log) Debug(message string) {
	logrus.WithFields(log.fields).Debug(message)
}

func Info(message string) {
	New().Info(message)
}

func (log *Log) Info(message string) {
	logrus.WithFields(log.fields).Info(message)
}

func Warn(message string) {
	New().Warn(message)
}

func (log *Log) Warn(message string) {
	logrus.WithFields(log.fields).Warn(message)
}

func Error(message string) {
	New().Error(message)
}

func (log *Log) Error(message string) {
	logrus.WithFields(log.fields).Error(message)
}
