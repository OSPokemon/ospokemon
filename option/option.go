package option

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type option struct {
	Value  string
	Prompt string
}

var options = map[string]*option{
	"log":          &option{"info", "One of [debug,info,warn,error,fatal,panic]"},
	"logpath":      &option{"ospokemon.log", "Log file path"},
	"port":         &option{"80", "Server port to open"},
	"dbpath":       &option{"ospokemon.db", "Server port to open"},
	"webpath":      &option{"www", "Server homepage path"},
	"interact":     &option{"false", "Launch interactive mode"},
	"refresh":      &option{"250", "Refresh rate in milliseconds"},
	"sessionlife":  &option{"180", "Session lifetime in seconds"},
	"passwordsalt": &option{"ospokemonsalt", "Salt string for password hashing"},
}

func init() {
	flags := setupflags()
	readfile("settings.txt")
	bindflags(flags)

	loginit()
}

func String(name string) string {
	return options[name].Value
}

func Int(name string) int {
	i, _ := strconv.Atoi(String(name))
	return i
}

func Duration(name string) time.Duration {
	d, _ := time.ParseDuration(String(name))
	return d
}

func Bool(name string) bool {
	b, _ := strconv.ParseBool(String(name))
	return b
}

func setupflags() map[string]*string {
	read := make(map[string]*string)
	for key, option := range options {
		read[key] = flag.String(key, "", option.Prompt)
	}
	flag.Parse()
	return read
}

func readfile(path string) error {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}

	for _, line := range strings.Split(string(file), "\n") {
		if line == "" || line[0] == "#"[0] {
			continue
		}

		setting := strings.Split(line, "=")
		if options[setting[0]] == nil {
			logrus.WithFields(logrus.Fields{
				"Setting": setting,
			}).Warn("Setting name not recognized")
			continue
		}
		if len(setting) == 2 {
			options[setting[0]].Value = setting[1]
		}
	}

	return nil
}

func bindflags(read map[string]*string) {
	for key, value := range read {
		if *value == "" {
			continue
		}

		options[key].Value = *value
	}
}

func loginit() {
	switch String("log") {
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
		logrus.Warn("Log level invalid: ", String("log"))
		break
	}
}