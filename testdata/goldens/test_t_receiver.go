package testdata

import (
	"errors"
	"testing"
)

func TestTestReceiver_FooMethod(t *testing.T) {
	testCases := []struct {
		name    string
		tr      *TestReceiver
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		tr := &TestReceiver{}
		if err := tr.FooMethod(); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. TestReceiver.FooMethod() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
