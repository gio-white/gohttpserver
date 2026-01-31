package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("no authorization header found")
	}
	
	splitHeader := strings.Split(authHeader, " ")
	if len(splitHeader) != 2 || splitHeader[0] != "Bearer" {
		return "", fmt.Errorf("malformed authorization header")
	}
	return splitHeader[1], nil
}
