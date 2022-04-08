package foo

import (
	"errors"
	"testing"
)

func TestFoo_Foo(t *testing.T) {
	type fields struct {
		Bar string
	}

	type args struct {
		s string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		f := &Foo{
			Bar: testCase.fields.Bar,
		}
		if err := f.Foo(testCase.args.s); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo.Foo() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
