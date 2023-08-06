package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomNumericString(length int) string {
	const charset = "0123456789"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}
	return string(b)
}
