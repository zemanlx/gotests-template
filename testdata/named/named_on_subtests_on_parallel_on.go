package testdata

import "testing"

func TestFoo038(t *testing.T) {
	testCases := map[string]struct {
		want bool
	}{
		// TODO: Add test cases.
	}
	for name, testCase := range testCases {
		testCase := testCase
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Foo038(); got != testCase.want {
				t.Errorf("Foo038() = %v, want %v", got, tt.want)
			}
		})
	}
}
