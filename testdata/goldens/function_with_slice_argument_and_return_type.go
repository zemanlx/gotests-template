package testdata

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFoo11(t *testing.T) {
	type args struct {
		strs []string
	}

	testCases := []struct {
		name    string
		args    args
		want    []*Bar
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := Foo11(testCase.args.strs)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo11() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}

		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo11() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
