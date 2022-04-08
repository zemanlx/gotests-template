package syntaxtest

import (
	"os"
	"testing"
)

// Plural all the types.
func Foo(s strings) errors {
	// Incorrect return type.
	return ""
}

func TestNot(t *testing.T) {
	type args struct {
		this *os.File
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := Not(testCase.args.this); got != testCase.want {
			t.Errorf("%q. Not() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
