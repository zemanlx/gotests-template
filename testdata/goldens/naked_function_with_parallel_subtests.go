package testdata

import "testing"

func Test_main(t *testing.T) {
	testCases := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			main()
		})
	}
}

func Test_do(t *testing.T) {
	testCases := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			do()
		})
	}
}
