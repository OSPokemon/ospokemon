package util

import (
	"github.com/Sirupsen/logrus"
)

func loginit() {
	switch Opt("log") {
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
		logrus.Warn("Log level invalid: ", Opt("log"))
		break
	}
}
