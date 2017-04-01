package server

import (
	"crypto/md5"
	"ospokemon.com/option"
)

func hashpassword(password string) string {
	hash := md5.Sum([]byte(password + option.String("passwordsalt")))
	password = string(hash[:])
	return password
}
