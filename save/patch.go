package save

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/mattes/migrate/driver/sqlite3"
	"github.com/mattes/migrate/migrate"
	"github.com/ospokemon/ospokemon/util"
)

const PATCH uint64 = 1

func CheckPatch() bool {
	if util.Opt("patchpath") != "" {
		Patch()
		return false
	} else if patch, _ := migrate.Version("sqlite3://"+util.Opt("dbpath"), util.Opt("patchpath")); patch != PATCH {
		logrus.WithFields(logrus.Fields{
			"Found":    patch,
			"Expected": PATCH,
		}).Fatal("Database patch mismatch")
		return false
	}

	return true
}

func Patch() {
	errors, ok := migrate.UpSync("sqlite3://"+util.Opt("dbpath"), util.Opt("patchpath"))

	if !ok {
		for _, err := range errors {
			logrus.Error(err)
		}
	}
}
