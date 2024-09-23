package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqsrtuvwxyz"

var domains = []string{"gmail.com", "yahoo.com", "hotmail.com", "example.com"} 

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder 

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomNameString() string{
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s", RandomString(6), domains[rand.Intn(len(domains))])
}


func RandomMessageString() string{
	return RandomString(110)
}

