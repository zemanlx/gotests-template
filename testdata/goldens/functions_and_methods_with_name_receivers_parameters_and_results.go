package testdata

import "testing"

func Test_name_Name(t *testing.T) {
	type args struct {
		n string
	}

	testCases := []struct {
		name string
		n    name
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.n.Name(testCase.args.n); got != testCase.want {
			t.Errorf("%q. name.Name() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func TestName_Name1(t *testing.T) {
	type args struct {
		n string
	}

	testCases := []struct {
		name string
		n    *Name
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.n.Name1(testCase.args.n); got != testCase.want {
			t.Errorf("%q. Name.Name1() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func TestName_Name2(t *testing.T) {
	type args struct {
		name string
	}

	testCases := []struct {
		name string
		n    *Name
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.n.Name2(testCase.args.name); got != testCase.want {
			t.Errorf("%q. Name.Name2() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func TestName_Name3(t *testing.T) {
	type args struct {
		nn string
	}

	testCases := []struct {
		name     string
		n        *Name
		args     args
		wantName string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if gotName := testCase.n.Name3(testCase.args.nn); gotName != testCase.wantName {
			t.Errorf("%q. Name.Name3() = %v, want %v", testCase.name, gotName, tt.wantName)
		}
	}
}
