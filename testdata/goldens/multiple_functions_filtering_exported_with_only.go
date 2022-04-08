package testdata

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFooFilter(t *testing.T) {
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
		got, err := FooFilter(testCase.args.strs)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. FooFilter() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}

		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. FooFilter() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
