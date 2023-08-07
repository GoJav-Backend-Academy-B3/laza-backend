package helper

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"strings"
)

func GenerateRandomTokenString(len int) (string, error) {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("could not random token %w", err)
	}

	result := strings.ToLower(base32.StdEncoding.EncodeToString(randomBytes)[:len])
	return result, nil
}
