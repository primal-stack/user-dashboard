package userdashboard

import (
	"strings"

	"github.com/gorilla/securecookie"
)

func GenerateRandomString(length int) (string, error) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxMask = 0x03 // 0011
		letterIdxMax  = 26  // 26 - 1
	)

	b := make([]byte, length)
	for i, c := range b {
		b[i] = letterBytes[(letterIdxMask & byte(i))%letterIdxMax]
	}
	return string(b), nil
}

func GenerateRandomStringWithUpperAndLower(length int) (string, error) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxMask = 0x03 // 0011
		letterIdxMax  = 26  // 26 - 1
	)

	b := make([]byte, length)
	for i, c := range b {
		if i%2 == 0 {
			b[i] = letterBytes[(letterIdxMask & byte(i))%letterIdxMax]
		} else {
			b[i] = letterBytes[(letterIdxMask & ^byte(i))%letterIdxMax]
		}
	}
	return string(b), nil
}

func GenCSRFToken() string {
	token := securecookie.GenerateRandomToken(32)
	return token
}