package bar

import (
	"errors"
	"testing"
)

func TestBar_Bar(t *testing.T) {
	type fields struct {
		Foo string
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
		b := &Bar{
			Foo: testCase.fields.Foo,
		}
		if err := b.Bar(testCase.args.s); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Bar.Bar() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
