package utils

import (
	"crypto/rand"
	"encoding/binary"
)

func GetRandomNumber() int {

	// Generate 8 bytes of random data
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// Interpret the bytes as an integer (example)
	randomNumber := binary.BigEndian.Uint64(b)

	return int(randomNumber)
}
