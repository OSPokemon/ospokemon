package security

import (
	"crypto/md5"

	"ztaylor.me/env"
)

func HashPassword(env env.Provider, password string) string {
	salt := env.Get("passwordsalt")
	hash := md5.Sum([]byte(password + salt))
	return string(hash[:])
}
