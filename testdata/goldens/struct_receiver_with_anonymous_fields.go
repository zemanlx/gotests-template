package testdata

import "testing"

func TestDoctor_SayHello(t *testing.T) {
	type args struct {
		r *Person
	}

	testCases := []struct {
		name string
		d    *Doctor
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.d.SayHello(testCase.args.r); got != testCase.want {
			t.Errorf("%q. Doctor.SayHello() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
