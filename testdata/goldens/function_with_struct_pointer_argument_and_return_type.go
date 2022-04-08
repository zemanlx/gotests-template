package testdata

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo8(t *testing.T) {
	type args struct {
		b *Bar
	}

	testCases := []struct {
		name    string
		args    args
		want    *Bar
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := Foo8(testCase.args.b)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo8() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}

		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo8() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
