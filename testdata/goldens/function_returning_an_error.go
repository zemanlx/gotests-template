package testdata

import (
	"errors"
	"testing"
)

func TestFoo5(t *testing.T) {
	testCases := []struct {
		name    string
		want    string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := Foo5()
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo5() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. Foo5() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
