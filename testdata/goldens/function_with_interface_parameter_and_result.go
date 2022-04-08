package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo21(t *testing.T) {
	type args struct {
		i interface{}
	}

	testCases := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Foo21(testCase.args.i)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo21() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
