package countrycode

import (
	"testing"
)

func TestConvertAlpha2(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"us", "usa", "8400",
	}
	for _, input := range cases {
		output, err := Convert(input, Alpha2)
		if err != nil {
			if err == ErrInvalidCountryCode || err == ErrInvalidFormat {
				t.Logf("INPUT: %+v --> OUTPUT: %+v", input, err)
				continue
			}
			t.Fatal(err)
		}
		t.Logf("INPUT: %+v --> OUTPUT: %+v", input, output)
	}
}

func TestConvertAlpha3(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"us", "usa", "8400",
	}
	for _, input := range cases {
		output, err := Convert(input, Alpha3)
		if err != nil {
			if err == ErrInvalidCountryCode || err == ErrInvalidFormat {
				t.Logf("INPUT: %+v --> OUTPUT: %+v", input, err)
				continue
			}
			t.Fatal(err)
		}
		t.Logf("INPUT: %+v --> OUTPUT: %+v", input, output)
	}
}

func TestConvertNumeric(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"us", "usa", "8400",
	}
	for _, input := range cases {
		output, err := Convert(input, Numeric)
		if err != nil {
			if err == ErrInvalidCountryCode || err == ErrInvalidFormat {
				t.Logf("INPUT: %+v --> OUTPUT: %+v", input, err)
				continue
			}
			t.Fatal(err)
		}
		t.Logf("INPUT: %+v --> OUTPUT: %+v", input, output)
	}
}

func TestConvertError(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"us", "usa", "8400",
	}
	for _, input := range cases {
		output, err := Convert(input, "hello_world")
		if err != nil {
			if err == ErrInvalidCountryCode || err == ErrInvalidFormat {
				t.Logf("INPUT: %+v --> OUTPUT: %+v", input, err)
				continue
			}
			t.Fatal(err)
		}
		t.Logf("INPUT: %+v --> OUTPUT: %+v", input, output)
	}
}

func TestGetCountry(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"us", "usa", "8400",
	}
	for _, input := range cases {
		output, err := getCountry(input)
		if err != nil {
			if err == ErrInvalidCountryCode {
				t.Logf("INPUT:%+v --> OUTPUT:%+v", input, err)
				continue
			}
			t.Fatal(err)
		}
		t.Logf("INPUT:%+v --> OUTPUT:%+v", input, output)
	}
}

func TestGetFormat(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"us", "usa", "8400",
	}
	for _, input := range cases {
		country, err := getCountry(input)
		if err != nil {
			if err == ErrInvalidCountryCode {
				t.Logf("INPUT:%+v --> OUTPUT:%+v", input, err)
				continue
			}
			t.Fatal(err)
		}
		output, err := getFormat(country, "numeric")
		if err != nil {
			if err == ErrInvalidFormat {
				t.Logf("INPUT:%+v --> OUTPUT:%+v", input, err)
				continue
			}
			t.Fatal(err)
		}
		t.Logf("INPUT:%+v --> OUTPUT:%+v", input, output)
	}
}
