package countrycode

import (
	"testing"
)

// func TestMain(m *testing.M) {
// log.SetOutput(ioutil.Discard)
// os.Exit(m.Run())
// }

func TestConvert(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"GB", "GBR", "826",
	}
	for _, input := range cases {
		t.Log(input)
		c, err := Convert(input, Numeric)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(c)
	}
}

func TestGetCountry(t *testing.T) {
	var cases []string = []string{
		"US", "USA", "840",
		"CA", "CAN", "124",
		"MX", "MEX", "484",
		"GB", "GBR", "826",
	}
	for _, input := range cases {
		t.Log(input)
		c, err := getCountry(input)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(c)
	}
}

func TestGetFormat(t *testing.T) {

}
