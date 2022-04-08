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
		t.Run(testCase.name, func(t *testing.T) {
			got, err := SameName()
			if !errors.Is(err, testCase.wantErr) {
				t.Errorf("SameName() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			if got != testCase.want {
				t.Errorf("SameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(testCase.name, func(t *testing.T) {
			got, err := sameName()
			if !errors.Is(err, testCase.wantErr) {
				t.Errorf("sameName() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			if got != testCase.want {
				t.Errorf("sameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(testCase.name, func(t *testing.T) {
			got, err := testCase.tr.SameName()
			if !errors.Is(err, testCase.wantErr) {
				t.Errorf("SameTypeName.SameName() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			if got != testCase.want {
				t.Errorf("SameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(testCase.name, func(t *testing.T) {
			got, err := testCase.tr.sameName()
			if !errors.Is(err, testCase.wantErr) {
				t.Errorf("SameTypeName.sameName() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			if got != testCase.want {
				t.Errorf("SameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(testCase.name, func(t *testing.T) {
			got, err := testCase.tr.SameName()
			if !errors.Is(err, testCase.wantErr) {
				t.Errorf("sameTypeName.SameName() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			if got != testCase.want {
				t.Errorf("sameTypeName.SameName() = %v, want %v", got, tt.want)
			}
		})
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
		t.Run(testCase.name, func(t *testing.T) {
			got, err := testCase.tr.sameName()
			if !errors.Is(err, testCase.wantErr) {
				t.Errorf("sameTypeName.sameName() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			if got != testCase.want {
				t.Errorf("sameTypeName.sameName() = %v, want %v", got, tt.want)
			}
		})
	}
}
