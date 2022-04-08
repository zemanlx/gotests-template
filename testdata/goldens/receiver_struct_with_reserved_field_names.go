package testdata

import "testing"

func TestReserved_DontFail(t *testing.T) {
	testCases := []struct {
		name string
		r    *Reserved
		want string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		if got := testCase.r.DontFail(); got != testCase.want {
			t.Errorf("%q. Reserved.DontFail() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
