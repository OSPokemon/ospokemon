package save

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/mattes/migrate/driver/sqlite3"
	"github.com/mattes/migrate/migrate"
	"github.com/ospokemon/ospokemon/util"
)

func CheckPatch() uint64 {
	patch, err := migrate.Version("sqlite3://"+util.Opt("dbpath"), util.Opt("patchpath"))

	if err != nil {
		logrus.Error(err)
	}

	return patch
}

func Patch() {
	oldpatch := CheckPatch()

	errors, ok := migrate.UpSync("sqlite3://"+util.Opt("dbpath"), util.Opt("patchpath"))

	if !ok {
		for _, err := range errors {
			logrus.Error(err)
		}
	}

	newpatch := CheckPatch()

	logrus.WithFields(logrus.Fields{
		"OldPatch": oldpatch,
		"NewPatch": newpatch,
	}).Warn("save.Patch")
}
