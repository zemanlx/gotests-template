package testdata

import (
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo17(t *testing.T) {
	type args struct {
		r io.Reader
	}

	testCases := []struct {
		name string
		args args
		want io.Reader
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Foo17(testCase.args.r)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo17() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
