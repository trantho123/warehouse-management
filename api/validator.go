package api

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	hasUpperCase    = regexp.MustCompile(`[A-Z]`).MatchString
	hasSpecialChar  = regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString
	hasNumber       = regexp.MustCompile(`[0-9]`).MatchString
)

func passwordValidator(fl validator.FieldLevel) bool {
	pass := fl.Field().String()
	if len(pass) < 6 || len(pass) > 100 {
		return false
	}
	if !hasUpperCase(pass) {
		return false
	}
	if !hasSpecialChar(pass) {
		return false
	}
	if !hasNumber(pass) {
		return false
	}

	return true
}

func usernameValidator(fl validator.FieldLevel) bool {
	userName := fl.Field().String()
	if len(userName) < 5 || len(userName) > 100 {
		return false
	}
	return isValidUsername(userName)
}

func validationErrorMessage(e validator.FieldError) error {
	switch e.Field() {
	case "Username":
		if e.Tag() == "required" {
			return fmt.Errorf("username is required")
		}
		if e.Tag() == "username" {
			return fmt.Errorf("username must contain only lowercase letters, numbers and underscores")
		}
	case "Password":
		if e.Tag() == "required" {
			return fmt.Errorf("password is required")
		}
		if e.Tag() == "passwd" {
			return fmt.Errorf("password must be at least 6 characters and contain both letters and numbers")
		}
	case "Email":
		if e.Tag() == "required" {
			return fmt.Errorf("email is required")
		}
		if e.Tag() == "email" {
			return fmt.Errorf("invalid email format")
		}
	}
	return fmt.Errorf("invalid field: %s", e.Field())
}
