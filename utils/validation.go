package utils

import (
	"errors"
	"regexp"
)

var (
  uppercasePattern   = regexp.MustCompile(`[A-Z]`)
  lowercasePattern   = regexp.MustCompile(`[a-z]`)
  digitPattern       = regexp.MustCompile(`[0-9]`)
  specialCharPattern = regexp.MustCompile(`[^\w\s]`)
  spacePattern       = regexp.MustCompile(`\s`)
)

func ValidateDisplayName(value string) string {
	if value == "" {
		return "Display name is required"
	}
	if len(value) > 50 {
		return "Display name cannot exceed 50 characters"
	}
	re := regexp.MustCompile(`^[a-zA-Z0-9 _-]+$`)
	if !re.MatchString(value) {
		return "Display name can only contain letters, numbers, spaces, underscores, and hypnens"
	}
	return ""
}

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	// Basic email regex pattern
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidateUsername(value string) string {
	if value == "" {
		return "Username is required"
	}
	if len(value) > 30 {
		return "Username cannot exceed 30 characters"
	}
	// Allow only alphanumeric characters (no spaces or special characters)
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(value) {
		return "Username can only contain letters and numbers"
	}
	return ""
}

func ValidatePassword(password string) error {
  if password == "" {
    return errors.New("password is required")
  }
  if len(password) < 8 {
    return errors.New("password must be at least 8 characters")
  }
  if !uppercasePattern.MatchString(password) {
    return errors.New("password must contain at least one uppercase letter")
  }
  if !lowercasePattern.MatchString(password) {
    return errors.New("password must contain at least one lowercase letter")
  }
  if !digitPattern.MatchString(password) {
    return errors.New("password must contain at least one digit")
  }
  if !specialCharPattern.MatchString(password) {
    return errors.New("password must contain at least one special character")
    }
  if spacePattern.MatchString(password) {
    return errors.New("password cannot contain spaces")
  }
  return nil
}