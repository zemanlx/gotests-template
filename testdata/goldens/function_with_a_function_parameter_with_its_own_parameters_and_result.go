package testdata

import (
	"errors"
	"testing"
)

func TestFoo14(t *testing.T) {
	type args struct {
		f func(string, int) string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if err := Foo14(testCase.args.f); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo14() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
