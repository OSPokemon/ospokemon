package util

import (
	"github.com/Sirupsen/logrus"
)

var Log = logrus.New()

func loginit() {
	switch Opt("log") {
	case "debug":
		Log.Level = logrus.DebugLevel
		break
	case "info":
		Log.Level = logrus.InfoLevel
		break
	case "warn":
		Log.Level = logrus.WarnLevel
		break
	case "error":
		Log.Level = logrus.ErrorLevel
		break
	case "fatal":
		Log.Level = logrus.FatalLevel
		break
	case "panic":
		Log.Level = logrus.PanicLevel
		break
	default:
		Log.Level = logrus.InfoLevel
		Log.Warn("Log level invalid: ", Opt("log"))
		break
	}
}
