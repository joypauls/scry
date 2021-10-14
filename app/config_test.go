package app

import "testing"

// Just test some logical stuff and sanity checks.
func TestMakeConfig(t *testing.T) {
	c := MakeConfig()
	// test default behavior
	if res, exp := c.ShowHidden, false; res != exp {
		t.Errorf("Result: %t, Expected: %t", res, exp)
	}
	if res, exp := c.UseEmoji, false; res != exp {
		t.Errorf("Result: %t, Expected: %t", res, exp)
	}
	if res := c.Home; res == nil {
		t.Error("Result: nil, Expected: non-nil")
	}
}
