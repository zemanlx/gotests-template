package testdata

import (
	"errors"
	"testing"
)

func TestFoo6(t *testing.T) {
	type args struct {
		i int
		b bool
	}

	testCases := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := Foo6(testCase.args.i, testCase.args.b)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo6(%v, %v) error = %v, wantErr %v", testCase.name, testCase.args.i, testCase.args.b, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. Foo6(%v, %v) = %v, want %v", testCase.name, testCase.args.i, testCase.args.b, got, tt.want)
		}
	}
}
