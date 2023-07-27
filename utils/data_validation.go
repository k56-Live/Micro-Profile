package utils

import (
	"errors"
	"regexp"
)

// ValidateEmail checks if the given email is in a valid format
func ValidateEmail(email string) error {
	// Implement email validation logic using regular expressions or other methods
	// For example, you can use a regular expression to check if the email matches a valid format

	// Example implementation (dummy logic)
	validEmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !validEmailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

// ValidatePassword checks if the given password meets the required criteria
func ValidatePassword(password string) error {
	// Implement password validation logic
	// For example, you can check the password length and complexity requirements

	// Example implementation (dummy logic)
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
}
