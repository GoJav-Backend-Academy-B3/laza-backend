package helper

import (
	"fmt"
	"time"
)

func GetExpiryDate(minutesToAdd time.Duration, timezone string) (time.Time, error) {
	// Load the specified timezone location
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, fmt.Errorf("error loading timezone: %v", err)
	}

	// Get the current time in the specified timezone
	currentTimeInTimezone := time.Now().In(location)

	// Add minutes to the current time
	expiryDate := currentTimeInTimezone.Add(minutesToAdd)

	return expiryDate, nil
}
