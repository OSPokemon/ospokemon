package security

import (
	"crypto/md5"

	"ztaylor.me/env"
)

func HashPassword(password string) string {
	env := env.Global()
	hash := md5.Sum([]byte(password + env.Get("passwordsalt")))
	password = string(hash[:])
	return password
}
