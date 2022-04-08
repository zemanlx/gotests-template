package testdata

import (
	"go/types"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestImporter_Foo35(t *testing.T) {
	type args struct {
		t types.Type
	}

	testCases := []struct {
		name string
		i    *Importer
		args args
		want *types.Var
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := testCase.i.Foo35(testCase.args.t)
		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Importer.Foo35() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}
