package save

import (
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
		util.Log.WithFields(map[string]interface{}{
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
			util.Log.Error(err)
		}
	}
}
