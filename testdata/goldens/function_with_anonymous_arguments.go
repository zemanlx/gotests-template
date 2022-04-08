package testdata

import "testing"

func TestFoo2(t *testing.T) {
	type args struct {
		in0 string
		in1 int
	}

	testCases := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		Foo2(testCase.args.in0, testCase.args.in1)
	}
}
