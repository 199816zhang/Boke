package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(data []byte) string {
	md5New := md5.New()
	md5New.Write(data)
	cipherStr := md5New.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
