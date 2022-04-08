package undefinedtypes

import "testing"

func TestSomeStruct_Do(t *testing.T) {
	type fields struct {
		Doer some.Doer
	}

	testCases := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		c := &SomeStruct{
			Doer: testCase.fields.Doer,
		}
		c.Do()
	}
}
