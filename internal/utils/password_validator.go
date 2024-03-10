package utils

import "regexp"

func ValidPassword(password string) bool {
	// Check length
	if len(password) < 8 || len(password) > 64 {
		return false
	}

	// Check for at least one letter
	if matched, _ := regexp.MatchString("[a-zA-Z]", password); !matched {
		return false
	}

	// Check for at least one digit
	if matched, _ := regexp.MatchString("[0-9]", password); !matched {
		return false
	}

	// Check for at least one special character
	if matched, _ := regexp.MatchString(`[\W_]`, password); !matched {
		return false
	}

	return true
}
