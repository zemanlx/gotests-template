package testdata

import (
	"errors"
	"testing"
)

func TestBook_Open(t *testing.T) {
	testCases := []struct {
		name    string
		b       *Book
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		b := &Book{}
		if err := b.Open(); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. Book.Open() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}

func Test_door_Open(t *testing.T) {
	testCases := []struct {
		name    string
		d       *door
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		d := &door{}
		if err := d.Open(); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. door.Open() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}

func Test_xml_Open(t *testing.T) {
	testCases := []struct {
		name    string
		x       *xml
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		x := &xml{}
		if err := x.Open(); !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. xml.Open() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)
		}
	}
}
