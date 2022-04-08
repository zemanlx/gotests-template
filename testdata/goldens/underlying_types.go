package testdata

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCelsius_ToFahrenheit(t *testing.T) {
	testCases := []struct {
		name string
		c    Celsius
		want Fahrenheit
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := testCase.c.ToFahrenheit()
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Celsius.ToFahrenheit() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}

func TestHourToSecond(t *testing.T) {
	type args struct {
		h time.Duration
	}

	testCases := []struct {
		name string
		args args
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := HourToSecond(testCase.args.h); got != testCase.want {
			t.Errorf("%q. HourToSecond() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
