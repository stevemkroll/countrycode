package countrycode

import "errors"

// Errors
var (
	ErrInvalidFormat      = errors.New("error: invalid format")
	ErrInvalidCountryCode = errors.New("error: invalid country code")
)
