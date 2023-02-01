package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Hash256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash[:])
}
