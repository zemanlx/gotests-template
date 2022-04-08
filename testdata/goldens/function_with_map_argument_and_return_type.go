package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo10(t *testing.T) {
	type args struct {
		m map[string]int32
	}

	testCases := []struct {
		name string
		args args
		want map[string]*Bar
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Foo10(testCase.args.m)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo10() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
