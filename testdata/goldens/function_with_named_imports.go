package testdata

import (
	ht "html/template"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo22(t *testing.T) {
	type args struct {
		t *ht.Template
	}

	testCases := []struct {
		name string
		args args
		want *ht.Template
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Foo22(testCase.args.t)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo22() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
