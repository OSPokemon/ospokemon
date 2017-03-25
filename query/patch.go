package query

import (
	_ "github.com/mattes/migrate/driver/sqlite3"
	"github.com/mattes/migrate/migrate"
	"github.com/ospokemon/ospokemon/log"
	"github.com/ospokemon/ospokemon/option"
)

func Patch() {
	errors, ok := migrate.UpSync("sqlite3://"+option.String("dbpath"), option.String("patchpath"))

	if !ok {
		for _, err := range errors {
			log.Add("Path", option.String("dbpath")).Error(err.Error())
		}
	}
}

func CheckPatch() uint64 {
	patch, _ := migrate.Version("sqlite3://"+option.String("dbpath"), option.String("patchpath"))

	return patch
}
