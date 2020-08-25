package countrycode

import "strings"

// Constants
const (
	Alpha2  string = "alpha2"
	Alpha3  string = "alpha3"
	Numeric string = "numeric"
)

// Country defines the country struct
type Country struct {
	Alpha2  string
	Alpha3  string
	Numeric string
}

// Convert converts the country code into the desired output type
func Convert(input string, output string) (string, error) {
	input = strings.ToUpper(input)
	switch input {
	case US.Alpha2, US.Alpha3, US.Numeric:
		switch output {
		case Alpha2:
			return Countries["US"].Alpha2, nil
		case Alpha3:
			return Countries["US"].Alpha3, nil
		case Numeric:
			return Countries["US"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case CA.Alpha2, CA.Alpha3, CA.Numeric:
		switch output {
		case Alpha2:
			return Countries["CA"].Alpha2, nil
		case Alpha3:
			return Countries["CA"].Alpha3, nil
		case Numeric:
			return Countries["CA"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case MX.Alpha2, MX.Alpha3, MX.Numeric:
		switch output {
		case Alpha2:
			return Countries["MX"].Alpha2, nil
		case Alpha3:
			return Countries["MX"].Alpha3, nil
		case Numeric:
			return Countries["MX"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case GB.Alpha2, GB.Alpha3, GB.Numeric:
		switch output {
		case Alpha2:
			return Countries["GB"].Alpha2, nil
		case Alpha3:
			return Countries["GB"].Alpha3, nil
		case Numeric:
			return Countries["GB"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case DK.Alpha2, DK.Alpha3, DK.Numeric:
		switch output {
		case Alpha2:
			return Countries["DK"].Alpha2, nil
		case Alpha3:
			return Countries["DK"].Alpha3, nil
		case Numeric:
			return Countries["DK"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case FI.Alpha2, FI.Alpha3, FI.Numeric:
		switch output {
		case Alpha2:
			return Countries["FI"].Alpha2, nil
		case Alpha3:
			return Countries["FI"].Alpha3, nil
		case Numeric:
			return Countries["FI"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case FR.Alpha2, FR.Alpha3, FR.Numeric:
		switch output {
		case Alpha2:
			return Countries["FR"].Alpha2, nil
		case Alpha3:
			return Countries["FR"].Alpha3, nil
		case Numeric:
			return Countries["FR"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case DE.Alpha2, DE.Alpha3, DE.Numeric:
		switch output {
		case Alpha2:
			return Countries["DE"].Alpha2, nil
		case Alpha3:
			return Countries["DE"].Alpha3, nil
		case Numeric:
			return Countries["DE"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case HU.Alpha2, HU.Alpha3, HU.Numeric:
		switch output {
		case Alpha2:
			return Countries["HU"].Alpha2, nil
		case Alpha3:
			return Countries["HU"].Alpha3, nil
		case Numeric:
			return Countries["HU"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case IS.Alpha2, IS.Alpha3, IS.Numeric:
		switch output {
		case Alpha2:
			return Countries["IS"].Alpha2, nil
		case Alpha3:
			return Countries["IS"].Alpha3, nil
		case Numeric:
			return Countries["IS"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case IT.Alpha2, IT.Alpha3, IT.Numeric:
		switch output {
		case Alpha2:
			return Countries["IT"].Alpha2, nil
		case Alpha3:
			return Countries["IT"].Alpha3, nil
		case Numeric:
			return Countries["IT"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case NL.Alpha2, NL.Alpha3, NL.Numeric:
		switch output {
		case Alpha2:
			return Countries["NL"].Alpha2, nil
		case Alpha3:
			return Countries["NL"].Alpha3, nil
		case Numeric:
			return Countries["NL"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case NO.Alpha2, NO.Alpha3, NO.Numeric:
		switch output {
		case Alpha2:
			return Countries["NO"].Alpha2, nil
		case Alpha3:
			return Countries["NO"].Alpha3, nil
		case Numeric:
			return Countries["NO"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case PL.Alpha2, PL.Alpha3, PL.Numeric:
		switch output {
		case Alpha2:
			return Countries["PL"].Alpha2, nil
		case Alpha3:
			return Countries["PL"].Alpha3, nil
		case Numeric:
			return Countries["PL"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case ES.Alpha2, ES.Alpha3, ES.Numeric:
		switch output {
		case Alpha2:
			return Countries["ES"].Alpha2, nil
		case Alpha3:
			return Countries["ES"].Alpha3, nil
		case Numeric:
			return Countries["ES"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case SE.Alpha2, SE.Alpha3, SE.Numeric:
		switch output {
		case Alpha2:
			return Countries["SE"].Alpha2, nil
		case Alpha3:
			return Countries["SE"].Alpha3, nil
		case Numeric:
			return Countries["SE"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	case CH.Alpha2, CH.Alpha3, CH.Numeric:
		switch output {
		case Alpha2:
			return Countries["CH"].Alpha2, nil
		case Alpha3:
			return Countries["CH"].Alpha3, nil
		case Numeric:
			return Countries["CH"].Numeric, nil
		default:
			return "", ErrUnsupportedFormat
		}
	}
	return "", ErrUnsupportedCountryCode
}
