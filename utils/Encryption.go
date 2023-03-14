package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Encryption(str string) string {
	has := md5.New()
	has.Write([]byte(str))
	b := has.Sum(nil)
	return hex.EncodeToString(b)
}
