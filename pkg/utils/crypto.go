package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Crypto(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func Md5CryptoWithSalt(password string, salt string) string {
	data := []byte(password + salt)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
