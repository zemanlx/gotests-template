package testdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBar_Foo9(t *testing.T) {
	testCases := []struct {
		name string
		b    Bar
		want Bar
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := testCase.b.Foo9()
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Bar.Foo9() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
