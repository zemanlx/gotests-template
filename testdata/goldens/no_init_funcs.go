package testdata

import "testing"

func Test_initFuncStruct_init(t *testing.T) {
	type fields struct {
		field int
	}

	testCases := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		i := initFuncStruct{
			field: testCase.fields.field,
		}
		if got := i.init(); got != testCase.want {
			t.Errorf("%q. initFuncStruct.init() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func Test_initFieldStruct_getInit(t *testing.T) {
	type fields struct {
		init int
	}

	testCases := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		i := initFieldStruct{
			init: testCase.fields.init,
		}
		if got := i.getInit(); got != testCase.want {
			t.Errorf("%q. initFieldStruct.getInit() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
