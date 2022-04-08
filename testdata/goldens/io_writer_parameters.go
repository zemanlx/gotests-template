package testdata

import (
	"bytes"
	"errors"
	"testing"
)

func TestBar_Write(t *testing.T) {
	testCases := []struct {
		name    string
		b       *Bar
		wantW   string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		w := &bytes.Buffer{}
		if err := testCase.b.Write(w); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Bar.Write() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if gotW := w.String(); gotW != testCase.wantW {
			t.Errorf("%q. Bar.Write() = %v, want %v", testCase.name, gotW, tt.wantW)
		}
	}
}

func TestWrite(t *testing.T) {
	type args struct {
		data string
	}

	testCases := []struct {
		name    string
		args    args
		wantW   string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		w := &bytes.Buffer{}
		if err := Write(w, testCase.args.data); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Write() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if gotW := w.String(); gotW != testCase.wantW {
			t.Errorf("%q. Write() = %v, want %v", testCase.name, gotW, tt.wantW)
		}
	}
}

func TestMultiWrite(t *testing.T) {
	type args struct {
		data string
	}

	testCases := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantW1  string
		wantW2  string
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		w1 := &bytes.Buffer{}
		w2 := &bytes.Buffer{}
		got, got1, err := MultiWrite(w1, w2, testCase.args.data)
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. MultiWrite() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. MultiWrite() got = %v, want %v", testCase.name, got, tt.want)
		}
		if got1 != testCase.want1 {
			t.Errorf("%q. MultiWrite() got1 = %v, want %v", testCase.name, got1, tt.want1)
		}
		if gotW1 := w1.String(); gotW1 != testCase.wantW1 {
			t.Errorf("%q. MultiWrite() gotW1 = %v, want %v", testCase.name, gotW1, tt.wantW1)
		}
		if gotW2 := w2.String(); gotW2 != testCase.wantW2 {
			t.Errorf("%q. MultiWrite() gotW2 = %v, want %v", testCase.name, gotW2, tt.wantW2)
		}
	}
}
