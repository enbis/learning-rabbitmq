package utils

import (
	"math/rand"
	"strings"
	"time"
)

func RandomStr() string {
	tn := time.Now().UnixNano()
	rand.Seed(tn)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}

func RandomInt(min, max int) int {
	tn := time.Now().Unix()
	rand.Seed(tn)
	return rand.Intn(max-min) + min
}

func SwitchBulb(i int) string {
	if i%2 == 0 {
		return "On"
	}
	return "Off"
}
