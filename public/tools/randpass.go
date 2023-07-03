package tools

import (
	"crypto/rand"
	"math/big"
)

const (
	passwordLength = 8
	letters        = "abcdefghijklmnopqrstu@vwxyzABCDEFGHIJKL#MNOP*QRSTUVWXYZ0123456789"
	lettersLength  = len(letters)
)

func GenerateRandomPassword() string {
	password := make([]byte, passwordLength)

	for i := range password {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(lettersLength)))
		password[i] = letters[index.Int64()]
	}

	return string(password)
}
