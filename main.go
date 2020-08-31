package countrycode

import (
	"log"
	"strings"
)

// https://www.iban.com/country-codes

// Constants
const (
	Alpha2  string = "alpha2"
	Alpha3  string = "alpha3"
	Numeric string = "numeric"
)

// Country defines the Country struct containing the Alpha2, Alpha3 and Numeric ISO codes
type Country struct {
	Alpha2  string
	Alpha3  string
	Numeric string
}

// Convert returns the converted ISO code string or an error
func Convert(code string, format string) (string, error) {
	code = strings.ToUpper(code)
	c := Country{}
	for i, v := range Countries {
		switch true {
		case v.Alpha2 == code, v.Alpha3 == code, v.Numeric == code:
			log.Println(Countries[i])
			c = Countries[i]
		default:
			return "", ErrInvalidCountryCode
		}
	}
	switch format {
	case Alpha2:
		return c.Alpha2, nil
	case Alpha3:
		return c.Alpha3, nil
	case Numeric:
		return c.Numeric, nil
	default:
		return "", ErrInvalidFormat
	}
}
