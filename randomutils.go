package easygo

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// RandomLetters used for Random Generating functions
var RandomLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandomIDwithPrefix generates random Letters with Prefix (letters can be modified by overwriting easygo.RandomLetters)
func RandomIDwithPrefix(prefix string, length int) string {

	b := make([]rune, length)
	for i := range b {
		b[i] = RandomLetters[rand.Intn(len(RandomLetters))]
	}
	return prefix + string(b)
}
