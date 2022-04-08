package testdata

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBarBar100(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       interface{}
		wantErr bool
	}{
		{
			name:    "Basic test",
			b:       &Bar{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if err := tt.b.Bar100(tt.i); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar100() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestBaz100(t *testing.T) {
	tests := []struct {
		name string
		f    *float64
		want float64
	}{
		{
			name: "Basic test",
			f:    func() *float64 { var x float64 = 64; return &x }(),
			want: 64,
		},
	}
	// TestBaz100 contains a comment.
	for _, tt := range tests {
		if got := baz100(tt.f); got != tt.want {
			t.Errorf("%q. baz100() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFoo100(t *testing.T) {
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
		got, err := Foo100(testCase.args.strs)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Foo100() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}

		if diff := cmp.Diff(got, testCase.want); diff != "" {
			t.Errorf("%q. Foo100() diff (-got +want)\n%s", testCase.name, diff)
		}
	}
}

func TestBar_Bar100(t *testing.T) {
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
		if err := testCase.b.Bar100(testCase.args.i); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Bar.Bar100() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}

func Test_baz100(t *testing.T) {
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
		if got := baz100(testCase.args.f); got != testCase.want {
			t.Errorf("%q. baz100() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
