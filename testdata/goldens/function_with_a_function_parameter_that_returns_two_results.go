package testdata

import (
	"errors"
	"testing"
)

func TestFoo15(t *testing.T) {
	type args struct {
		f func(string) (string, error)
	}

	testCases := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if err := Foo15(testCase.args.f); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo15() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
