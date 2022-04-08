package testdata

import (
	"errors"
	"testing"
)

func TestSameName(t *testing.T) {
	testCases := []struct {
		name    string
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := SameName()
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. SameName() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. SameName() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func Test_sameName(t *testing.T) {
	testCases := []struct {
		name    string
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := sameName()
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. sameName() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. sameName() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func TestSameTypeName_SameName(t *testing.T) {
	testCases := []struct {
		name    string
		tr      *SameTypeName
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := testCase.tr.SameName()
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. SameTypeName.SameName() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. SameTypeName.SameName() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func TestSameTypeName_sameName(t *testing.T) {
	testCases := []struct {
		name    string
		tr      *SameTypeName
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := testCase.tr.sameName()
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. SameTypeName.sameName() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. SameTypeName.sameName() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func Test_sameTypeName_SameName(t *testing.T) {
	testCases := []struct {
		name    string
		tr      *sameTypeName
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := testCase.tr.SameName()
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. sameTypeName.SameName() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. sameTypeName.SameName() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}

func Test_sameTypeName_sameName(t *testing.T) {
	testCases := []struct {
		name    string
		tr      *sameTypeName
		want    int
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got, err := testCase.tr.sameName()
		if !errors.Is(err, testCase.wantErr) {
			t.Errorf("%q. sameTypeName.sameName() error = %v, wantErr %v", testCase.name, err, testCase.wantErr)

			continue
		}
		if got != testCase.want {
			t.Errorf("%q. sameTypeName.sameName() = %v, want %v", testCase.name, got, tt.want)
		}
	}
}
