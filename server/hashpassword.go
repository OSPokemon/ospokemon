package server

import (
	"crypto/md5"
	"github.com/ospokemon/ospokemon/option"
)

func hashpassword(password string) string {
	hash := md5.Sum([]byte(password + option.String("passwordsalt")))
	password = string(hash[:])
	return password
}
