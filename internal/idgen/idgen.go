package idgen

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID generates a unique 128-bit identifier as a string encoded in hexadecimal format.
func GenerateID() string {
	bytes := make([]byte, 4) // 128-bit ID
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
