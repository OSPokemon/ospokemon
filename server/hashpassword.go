package server

import (
	"crypto/md5"
	"github.com/ospokemon/ospokemon/util"
)

func hashpassword(password string) string {
	hash := md5.Sum([]byte(password + util.Opt("passwordsalt")))
	password = string(hash[:])
	return password
}
