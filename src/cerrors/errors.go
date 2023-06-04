package cerrors

import "errors"

var (
	ErrUrlIsRequired     = errors.New("url is required")
	ErrInvalidURLFormat  = errors.New("url format is invalid")
	ErrCantPingUrl       = errors.New("our server cant ping the hostname of url, try another")
	ErrNameAlreadyExists = errors.New("url name already exists")
	ErrUrlAlreadyExists  = errors.New("url already exists")
	ErrNameIsRequired    = errors.New("name is required")
	ErrUrlNotFound       = errors.New("url not found in our database")
)
