package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetIntegerStr(length int) string {
	return GetRandomStr("0123456789", length)
}

func GetAlphabeticStr(length int) string {
	return GetRandomStr("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
}

func GetMixStr(length int) string {
	return GetRandomStr("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", length)
}

func GetRandomStr(str string, length int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune(str)
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func EncryptPassword(password, salt string) string {
	hash1 := md5.New()
	hash1.Write([]byte(password))
	pwdMd := hex.EncodeToString(hash1.Sum(nil))
	hash2 := md5.New()
	hash2.Write([]byte(pwdMd + salt))
	return hex.EncodeToString(hash2.Sum(nil))
}
