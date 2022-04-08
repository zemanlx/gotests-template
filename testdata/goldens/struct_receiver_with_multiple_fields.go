package testdata

import "testing"

func TestPerson_SayHello(t *testing.T) {
	type args struct {
		r *Person
	}

	testCases := []struct {
		name string
		p    *Person
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.p.SayHello(testCase.args.r); got != testCase.want {
			t.Errorf("%q. Person.SayHello() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
