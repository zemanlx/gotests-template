package testdata

import (
	"fmt"
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

func wrapToString(in []int) []string {
	var result []string
	for _, x := range in {
		result = append(result, fmt.Sprintf("%v", x))
	}
	return result
}

func Test_wrapToString(t *testing.T) {
	type args struct {
		in []int
	}

	testCases := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := wrapToString(testCase.args.in)
			if diff := cmp.Diff(got, testCase.want); diff != "" {
				t.Errorf("wrapToString() diff (-got +want)\n%s", diff)
			}
		})
	}
}
