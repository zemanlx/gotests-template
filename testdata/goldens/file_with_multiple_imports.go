package testdata

import (
	"errors"
	"go/ast"
	"go/types"
	"io"
	"testing"
)

func TestFoo24(t *testing.T) {
	type args struct {
		r io.Reader
		x ast.Expr
		t types.Type
	}

	testCases := []struct {
		name    string
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if err := Foo24(testCase.args.r, testCase.args.x, testCase.args.t); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo24() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
