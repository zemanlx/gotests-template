package testdata

import "testing"

func TestCelsius_String(t *testing.T) {
	testCases := []struct {
		name string
		c    Celsius
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.c.String(); got != testCase.want {
			t.Errorf("%q. Celsius.String() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func TestFahrenheit_String(t *testing.T) {
	testCases := []struct {
		name string
		f    Fahrenheit
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.f.String(); got != testCase.want {
			t.Errorf("%q. Fahrenheit.String() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
