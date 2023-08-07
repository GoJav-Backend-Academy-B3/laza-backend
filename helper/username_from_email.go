package helper

import "strings"

func ExtractUsernameFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[0]
	}
	return ""
}
