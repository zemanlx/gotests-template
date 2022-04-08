package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo16(t *testing.T) {
	type args struct {
		in Bazzar
	}

	testCases := []struct {
		name string
		args args
		want Bazzar
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Foo16(testCase.args.in)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo16() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
