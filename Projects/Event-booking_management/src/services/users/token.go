package users

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetAccessAndRefreshToken(l int) (string, string) {
	accessToken := RandStringBytes(l)
	refreshToken := RandStringBytes(l)
	return accessToken, refreshToken
}

func GetAuthorizationCode(l int) string {
	return RandStringBytes(l)
}