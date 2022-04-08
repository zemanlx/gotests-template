package testdata

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo25(t *testing.T) {
	type args struct {
		in0 interface{}
	}

	testCases := []struct {
		name    string
		args    args
		want    string
		want1   []byte
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, got1, err := Foo25(testCase.args.in0)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo25() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. Foo25() got = %v, want %v", testCase.name, got, tt.want)
		}

		if diff := cmp.Diff(got1, testCase.want1); diff != "" {
			t.Errorf("%q. Foo25() diff (-got1 +want1)\n%s", testCase.name, diff)
		}
	}
}
