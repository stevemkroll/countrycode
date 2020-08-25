package countrycode

import "errors"

// Errors
var (
	ErrUnsupportedFormat      = errors.New("Error: unsupported format")
	ErrUnsupportedCountryCode = errors.New("Error: unsupported country code")
)
