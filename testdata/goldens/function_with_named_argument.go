package testdata

import "testing"

func TestFoo3(t *testing.T) {
	type args struct {
		s string
	}

	testCases := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		Foo3(testCase.args.s)
	}
}
