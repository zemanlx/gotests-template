package testdata

import (
	"errors"
	"testing"
)

func TestBar_Foo7(t *testing.T) {
	type args struct {
		i int
	}

	testCases := []struct {
		name    string
		b       *Bar
		args    args
		want    string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := testCase.b.Foo7(testCase.args.i)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Bar.Foo7() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. Bar.Foo7() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
