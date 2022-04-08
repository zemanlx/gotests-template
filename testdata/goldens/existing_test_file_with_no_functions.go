package testdata

import (
	"fmt"
	"testing"
)

var example102 = fmt.Sprintf("test%v", 1)

func TestFoo102(t *testing.T) {
	type args struct {
		s string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := Foo102(testCase.args.s); got != testCase.want {
			t.Errorf("%q. Foo102() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
