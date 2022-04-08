package testdata

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo26(t *testing.T) {
	type args struct {
		v interface{}
	}

	testCases := []struct {
		name    string
		args    args
		want    string
		wantI   int
		want2   []byte
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, gotI, got2, err := Foo26(testCase.args.v)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo26() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. Foo26() got = %v, want %v", testCase.name, got, tt.want)
		}
		if gotI != testCase.wantI {
			t.Errorf("%q. Foo26() gotI = %v, want %v", testCase.name, gotI, tt.wantI)
		}

		if diff := cmp.Diff(got2, testCase.want2); diff != "" {
			t.Errorf("%q. Foo26() diff (-got2 +want2)\n%s", testCase.name, diff)
		}
	}
}
