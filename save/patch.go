package save

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/mattes/migrate/driver/sqlite3"
	"github.com/mattes/migrate/migrate"
	"github.com/ospokemon/ospokemon/option"
)

func CheckPatch() uint64 {
	patch, _ := migrate.Version("sqlite3://"+option.String("dbpath"), "migrations")

	return patch
}

func Patch() {
	errors, ok := migrate.UpSync("sqlite3://"+option.String("dbpath"), "migrations")

	if !ok {
		for _, err := range errors {
			logrus.WithFields(logrus.Fields{
				"Path": option.String("dbpath"),
			}).Error(err.Error())
		}
	}
}
