package util

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
	"logpath":      &option{"/var/log/ospokemon/ospokemon.log", "Log file path"},
	"port":         &option{"80", "Server port to open"},
	"webpath":      &option{"/srv/http/ospokemon.com/", "Server homepage path"},
	"dbpath":       &option{"/var/lib/ospokemon/ospokemon.db", "Database sqlite3 path"},
	"patchpath":    &option{"", "Database migrations path"},
	"interact":     &option{"false", "Launch interactive mode"},
	"refresh":      &option{"250", "Refresh rate in milliseconds"},
	"sessionlife":  &option{"180", "Session lifetime in seconds"},
	"passwordsalt": &option{"ospokemonsalt", "Salt string for password hashing"},
}

func init() {
	flags := setupflags()

	if args := flag.Args(); len(args) == 1 {
		readfile(args[0])
	} else {
		bindflags(flags)
	}

	loginit()
	logoptions()
}

func setupflags() map[string]*string {
	read := make(map[string]*string)
	for key, option := range options {
		read[key] = flag.String(key, option.Value, option.Prompt)
	}
	flag.Parse()
	return read
}

func Opt(name string) string {
	return options[name].Value
}

func OptInt(name string) int {
	i, _ := strconv.Atoi(Opt(name))
	return i
}

func OptDuration(name string) time.Duration {
	d, _ := time.ParseDuration(Opt(name))
	return d
}

func OptBool(name string) bool {
	b, _ := strconv.ParseBool(Opt(name))
	return b
}

func readfile(path string) {
	if file, e := ioutil.ReadFile(path); e != nil {
		logrus.Warn(e.Error())
	} else {
		for _, line := range strings.Split(string(file), "\n") {
			if line[0] == "#"[0] {
				continue
			}

			setting := strings.Split(line, "=")
			if len(setting) == 2 && setting[1] != "" {
				options[setting[0]].Value = setting[1]
			}
		}
	}
}

func bindflags(read map[string]*string) {
	for key, value := range read {
		options[key].Value = *value
	}
}

func logoptions() {
	for key, option := range options {
		logrus.Debug(key + ": " + option.Value)
	}
}
