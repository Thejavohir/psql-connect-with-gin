package helper

import (
	"errors"
	"regexp"
)

func ValidPinfl(pinfl string) error {
	if pinfl == "" {
		return errors.New("error application passport_pinfl requirement body to model")
	}

	pattern := regexp.MustCompile(`^([0-9]{14})$`)

	if !(pattern.MatchString(pinfl)) {
		return errors.New("passport_pinfl must contain 14 digits")
	}
	return nil
}

func ValidPassportNumber(number string) error {
	if number == "" {
		return errors.New("error application passport_number requirement body to model")
	}

	pattern := regexp.MustCompile(`^([0-9]{7})$`)

	if !(pattern.MatchString(number)) {
		return errors.New("passport_number must contain 7 digits")
	}
	return nil
}

func IsValidPhoneNumber(phoneNumber string) bool {
	r := regexp.MustCompile(`^\+998[0-9]{2}[0-9]{7}$`)
	return r.MatchString(phoneNumber)
}

func IsValidEmail(email string) bool {
	r := regexp.MustCompile(`^([\w.*-]+@([\w-]+\.)+[\w-]{2,4})?$`)
	return r.MatchString(email)
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	return r.MatchString(uuid)
}

func IsValidUUIDv1(uuid string) bool {
	r := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-1[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	return r.MatchString(uuid)
}