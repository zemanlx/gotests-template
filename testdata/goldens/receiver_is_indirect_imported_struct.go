package testdata

import "testing"

func Test_someIndirectImportedStruct_Foo037(t *testing.T) {
	testCases := []struct {
		name string
		smtg *someIndirectImportedStruct
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		smtg := &someIndirectImportedStruct{}
		smtg.Foo037()
	}
}
