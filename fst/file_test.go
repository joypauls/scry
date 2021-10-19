package fst

import (
	"testing"
)

// testing BytesSI type constants for their actual values
func TestBytesSIConstantsValues(t *testing.T) {
	if res, exp := KB, 1000.0; res != exp {
		t.Errorf("Result: %f, Expected: %f", res, exp)
	}

	if res, exp := GB, 1000000000.0; res != exp {
		t.Errorf("Result: %f, Expected: %f", res, exp)
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
		res := BytesSI(table.bytes).String()
		if res != table.expected {
			t.Errorf("Result: %s, Expected: %s", res, table.expected)
		}
	}
}

func TestMakeFile(t *testing.T) {
	dirRaw, err := testFs.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	f := MakeFile(dirRaw[0])
	// add other checks
	if res, exp := f.Name, ".env"; res != exp {
		t.Errorf("Result: %s, Expected: %s", res, exp)
	}
	if res, exp := f.IsDir, false; res != exp {
		t.Errorf("Result: %t, Expected: %t", res, exp)
	}
}
