package security

import (
	"crypto/md5"
	"ospokemon.com/option"
)

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password + option.String("passwordsalt")))
	password = string(hash[:])
	return password
}
