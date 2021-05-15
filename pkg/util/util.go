package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Password(pass *string) {
	hash := md5.Sum([]byte(*pass))
	*pass = hex.EncodeToString(hash[:])
}
