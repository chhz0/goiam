package idutil

import "crypto/rand"

// Defiens alphabet.
const (
	Alphabet62 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Alphabet36 = "abcdefghijklmnopqrstuvwxyz1234567890"
)

func randString(letters string, n int) string {
	buf := make([]byte, n)

	randomness := make([]byte, n)

	_, err := rand.Read(randomness)
	if err != nil {
		panic(err)
	}

	l := len(letters)
	for pos := range buf {
		randon := randomness[pos]

		randomPos := randon % uint8(l)

		buf[pos] = letters[randomPos]
	}

	return string(buf)
}

// func NewRandonStr62(n int) string {
// 	return randString(Alphabet62, n)
// }

// func NewRandonStr36(n int) string {
// 	return randString(Alphabet36, n)
// }

func NewSecretID() string {
	return randString(Alphabet62, 36)
}

func NewSecretKey() string {
	return randString(Alphabet62, 32)
}
