package testdata

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo18(t *testing.T) {
	type args struct {
		t *os.File
	}

	testCases := []struct {
		name string
		args args
		want *os.File
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Foo18(testCase.args.t)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo18() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
