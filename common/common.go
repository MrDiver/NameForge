package common

import (
	"math/rand/v2" // https://go.dev/blog/randv2
	"strings"
)

func Lines(content string) []string {
	return strings.Split(content, "\n")
}

func SelectRandom[T any](l []T) T {
	return l[rand.N(len(l))]
}
