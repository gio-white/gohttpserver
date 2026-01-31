package auth

import (
	"crypto/rand"
	"encoding/hex"

)

func MakeRefreshToken() (string, error) {
	token := make([]byte, 32)
	rand.Read(token)
	encodedToken := hex.EncodeToString(token)
	return encodedToken, nil
}