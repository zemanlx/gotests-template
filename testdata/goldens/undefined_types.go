package undefinedtypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUndefined_Do(t *testing.T) {
	type args struct {
		es Something
	}

	testCases := []struct {
		name    string
		u       *Undefined
		args    args
		want    *Unknown
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := testCase.u.Do(testCase.args.es)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Undefined.Do() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}

		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Undefined.Do() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
