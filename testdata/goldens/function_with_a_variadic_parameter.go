package testdata

import "testing"

func TestFoo20(t *testing.T) {
	type args struct {
		strs []string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := Foo20(testCase.args.strs...); got != testCase.want {
			t.Errorf("%q. Foo20() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
