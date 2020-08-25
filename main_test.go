package countrycode

import "testing"

func TestConvertNumeric(t *testing.T) {
	var c []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"GB", "GBR", "826",
		"XY", "XYZ", "000",
	}
	for _, i := range c {
		str, err := Convert(i, Numeric)
		if err != nil {
			if err != ErrUnsupportedCountryCode && err != ErrUnsupportedFormat {
				t.Fail()
			}
			t.Log(err.Error())
			continue
		}
		t.Log(str)
	}
}
