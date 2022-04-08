package testdata

import "testing"

func TestFoo4(t *testing.T) {
	testCases := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := Foo4(); got != testCase.want {
			t.Errorf("%q. Foo4() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
