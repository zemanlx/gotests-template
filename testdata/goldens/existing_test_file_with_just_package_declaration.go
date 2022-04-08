package testdata

import "testing"

func TestFoo101(t *testing.T) {
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
		if got := Foo101(testCase.args.s); got != testCase.want {
			t.Errorf("%q. Foo101() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
