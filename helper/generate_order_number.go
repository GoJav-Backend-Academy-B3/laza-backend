package helper

import (
	"fmt"
	"time"
)

func GenerateOrderNumber() string {

	year, month, day := time.Now().Date()
	result := fmt.Sprintf("ORD-%d%d%d-", year, month, day)

	result += GenerateRandomNumericString(10)

	return result
}
