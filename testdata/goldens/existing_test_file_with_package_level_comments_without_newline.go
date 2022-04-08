// Lorem ipsum dolor sit amet, consectetur adipiscing
// elit, sed do eiusmod tempor incididunt ut labore et
// dolore magna aliqua. Ut enim ad minim veniam, quis
// nostrud exercitation ullamco laboris nisi ut aliquip
// ex ea commodo consequat. Duis aute irure dolor in
// reprehenderit in voluptate velit esse cillum dolore
// eu fugiat nulla pariatur. Excepteur sint occaecat
// cupidatat non proident, sunt in culpa qui officia
// deserunt mollit anim id est laborum.

// Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
// tempor incididunt ut labore et dolore magna aliqua.
package testdata

import "testing"

func TestBeforeCommentNewLine(t *testing.T) {
	testCases := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := BeforeCommentNewLine(); got != testCase.want {
			t.Errorf("%q. BeforeCommentNewLine() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
