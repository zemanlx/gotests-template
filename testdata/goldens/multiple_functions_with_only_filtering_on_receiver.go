package testdata

import (
	"errors"
	"testing"
)

func TestBar_BarFilter(t *testing.T) {
	type args struct {
		i interface{}
	}

	testCases := []struct {
		name    string
		b       *Bar
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if err := testCase.b.BarFilter(testCase.args.i); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
