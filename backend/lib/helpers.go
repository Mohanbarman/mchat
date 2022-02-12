package lib

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"
)

func GenRandomString(salt []byte) string {
	b := make([]byte, 100)
	rand.Read(b)
	b = b[:len(salt)-1]
	b = append(b, salt...)
	encodedString := base64.URLEncoding.EncodeToString(b)
	return encodedString
}

func MustGetInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(i)
}
