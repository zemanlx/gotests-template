package testdata

import "testing"

func TestFoo19(t *testing.T) {
	type args struct {
		in1 string
		in2 string
		in3 string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := Foo19(testCase.args.in1, testCase.args.in2, testCase.args.in3); got != testCase.want {
			t.Errorf("%q. Foo19() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
