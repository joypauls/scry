package fst

import (
	"testing"
)

// testing BytesSI type constants for their actual values
func TestBytesSIConstantsValues(t *testing.T) {
	result := KB
	expected := 1000.0
	if result != expected {
		t.Errorf("Result: %f, Wanted: %f", result, expected)
	}

	result = GB
	expected = 1000000000.0
	if result != expected {
		t.Errorf("Result: %f, Expected: %f", result, expected)
	}
}

func TestBytesSIStringFormat(t *testing.T) {
	// table of test cases
	tables := []struct {
		bytes    int64  // this would be straight from FileInfo type
		expected string // after formatting
	}{
		{73, "73 B"},
		{1234, "1.2 KB"},
		{456700000, "456.7 MB"},
		{98765432100, "98.8 GB"},
		{9876543210011, "9.9 TB"},
		{987654321001111111, "987.7 PB"},
	}

	// iterate over test tables
	for _, table := range tables {
		result := BytesSI(table.bytes).String()
		if result != table.expected {
			t.Errorf("Result: %s, Expected: %s", result, table.expected)
		}
	}
}

// func TestFileStructCreation(t *testing.T) {
// 	result := KB
// 	expected := 1000.0
// 	if result != expected {
// 		t.Errorf("Result: %f, Wanted: %f", result, expected)
// 	}

// 	result = GB
// 	expected = 1000000000.0
// 	if result != expected {
// 		t.Errorf("Result: %f, Expected: %f", result, expected)
// 	}
// }
