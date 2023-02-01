package utils

import (
	"os"
	"regexp"
)

func PhoneNumberValidator(phone string) bool {
	val := regexp.MustCompile(`[^0-9]*1[34578][0-9]{9}[^0-9]*`)
	return val.MatchString(phone)
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
