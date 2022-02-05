package lib

import (
	"crypto/rand"
	"encoding/base64"
)

func GenRandomString(salt []byte) string {
	b := make([]byte, 100)
	rand.Read(b)
	b = b[:len(salt)-1]
	b = append(b, salt...)
	encodedString := base64.URLEncoding.EncodeToString(b)
	return encodedString
}
