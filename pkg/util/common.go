package util

import (
	"math/rand"
	"strings"
	"time"
)

var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Replace(table map[string]string, s string) string {
	for key, value := range table {
		s = strings.Replace(s, key, value, -1)
	}
	return s
}
