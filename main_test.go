package countrycode

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

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
			if err != ErrInvalidCountryCode && err != ErrInvalidFormat {
				t.Fail()
			}
			t.Log(err.Error())
			continue
		}
		t.Log(str)
	}
}
