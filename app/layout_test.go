package app

import "testing"

// Just test some logical stuff and sanity checks.
func TestMakeLayout(t *testing.T) {
	result := MakeLayout(20, 10)
	expectedW := 20
	expectedH := 10
	if result.width != expectedW {
		t.Errorf("Result: %d, Expected: %d", result.width, expectedW)
	} else if result.height != expectedH {
		t.Errorf("Result: %d, Expected: %d", result.height, expectedH)
	}
}
