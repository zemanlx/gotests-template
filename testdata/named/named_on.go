package testdata

import "testing"

func TestFoo038(t *testing.T) {
	testCases := map[string]struct {
		want bool
	}{
		// TODO: Add test cases.
	}
	for name, testCase := range testCases {
		if got := Foo038(); got != testCase.want {
			t.Errorf("%q. Foo038() = %v, want %v", name, got, tt.want)
		}
	}
}
