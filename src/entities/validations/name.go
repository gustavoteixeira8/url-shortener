package validations

import (
	"errors"
	"regexp"
)

func ValidateName(name string) error {
	nameLength := len(name)
	checkLength := nameLength > 2 && nameLength < 255
	testName, err := regexp.MatchString(`[a-zá-ýA-ZÁ-Ý]gm`, name)

	if err != nil {
		return errors.New("name is invalid")
	}

	if !testName && !checkLength {
		return errors.New("name is invalid")
	}

	return nil
}
