package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo23(t *testing.T) {
	type args struct {
		ch chan bool
	}

	testCases := []struct {
		name string
		args args
		want chan string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Foo23(testCase.args.ch)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo23() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
