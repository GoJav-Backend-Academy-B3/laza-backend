package helpers

import (
	"fmt"
	"github.com/phincon-backend/laza/helper"
	"time"
)

func GenerateOrderNumber() string {

	year, month, day := time.Now().Date()
	result := fmt.Sprintf("ORD-%d%d%d-", year, month, day)

	result += helper.GenerateRandomNumericString(10)

	return result
}
