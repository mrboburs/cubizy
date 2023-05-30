package srand

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const numberCharset = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//StringWithCharset gives randome string of given length from given charset
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//String gives randome string of given length from full charset from A-z and 0-9
func String(length int) string {
	return StringWithCharset(length, charset)
}

//Number gives randome string of given length from full charset from 0-9
func Number(length int) string {
	return StringWithCharset(length, numberCharset)
}
