package app

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	tables := []struct {
		a        int
		b        int
		expected int
	}{
		{300, 301, 300},
		{301, 300, 300},
		{300, 300, 300},
	}

	for _, table := range tables {
		result := minInt(table.a, table.b)
		if result != table.expected {
			t.Errorf("Result: %d, Expected: %d", result, table.expected)
		}
	}
}

func TestMaxInt(t *testing.T) {
	tables := []struct {
		a        int
		b        int
		expected int
	}{
		{300, 301, 301},
		{301, 300, 301},
		{300, 300, 300},
	}

	for _, table := range tables {
		result := maxInt(table.a, table.b)
		if result != table.expected {
			t.Errorf("Result: %d, Expected: %d", result, table.expected)
		}
	}
}
