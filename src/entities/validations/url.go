package validations

import (
	"errors"
	"net/url"
	"strings"
)

func ValidateURL(u string) (string, error) {
	uParsed, err := url.ParseRequestURI(u)

	if strings.HasPrefix(uParsed.String(), "http://") {
		return "", errors.New("url should have SSL")
	}

	return uParsed.String(), err
}
