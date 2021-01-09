package caliber

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringN(n int) string {
	N := len(letterRunes)
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(N)]
	}
	return string(b)
}
