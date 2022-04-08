package testdata

import (
	"errors"
	"testing"
)

func TestFoo12(t *testing.T) {
	type args struct {
		str string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if err := Foo12(testCase.args.str); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo12() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
