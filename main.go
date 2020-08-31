package countrycode

import (
	"strings"
)

// https://www.iban.com/country-codes

// Constants
const (
	Alpha2  string = "ALPHA2"
	Alpha3  string = "ALPHA3"
	Numeric string = "NUMERIC"
)

// Country defines the Country struct containing the Alpha2, Alpha3 and Numeric ISO codes
type Country struct {
	Alpha2  string
	Alpha3  string
	Numeric string
}

// Convert returns the converted ISO code string or an error
func Convert(code string, format string) (string, error) {
	country, err := getCountry(code)
	if err != nil {
		return "", err
	}
	convertedCode, err := getFormat(country, format)
	if err != nil {
		return "", err
	}
	return convertedCode, nil
}

func getCountry(code string) (Country, error) {
	code = strings.ToUpper(code)
	for _, country := range Countries {
		switch code {
		case country.Alpha2:
			return country, nil
		case country.Alpha3:
			return country, nil
		case country.Numeric:
			return country, nil
		default:
			continue
		}
	}
	return Country{}, ErrInvalidCountryCode
}

func getFormat(country Country, format string) (string, error) {
	format = strings.ToUpper(format)
	switch format {
	case Alpha2:
		return country.Alpha2, nil
	case Alpha3:
		return country.Alpha3, nil
	case Numeric:
		return country.Numeric, nil
	default:
		return "", ErrInvalidFormat
	}
}
