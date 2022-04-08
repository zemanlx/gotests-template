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

func TestBar_BarFilter(t *testing.T) {
	type args struct {
		i interface{}
	}

	testCases := []struct {
		name    string
		b       *Bar
		args    args
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if err := testCase.b.BarFilter(testCase.args.i); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Bar.BarFilter() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}

func Test_bazFilter(t *testing.T) {
	type args struct {
		f *float64
	}

	testCases := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := bazFilter(testCase.args.f); got != testCase.want {
			t.Errorf("%q. bazFilter() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
