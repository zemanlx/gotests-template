package testdata

import (
	"errors"
	"testing"
)

func TestFoo13(t *testing.T) {
	type args struct {
		f func()
	}

	testCases := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if err := Foo13(testCase.args.f); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo13() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
