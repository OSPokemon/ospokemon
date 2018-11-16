package security

import (
	"crypto/md5"
)

func HashPassword(salt string, password string) string {
	hash := md5.Sum([]byte(password + salt))
	return string(hash[:])
}
