package main

import (
	"testing"
)

func TestCleanText(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		got := textClean(c.input)
		if len(got) != len(c.expected) {
			t.Errorf("Expected %v, got %v", c.expected, got)
			continue
		}
		for i := range got {
			if got[i] != c.expected[i] {
				t.Errorf("Expected %v, got %v", c.expected, got)
			}
		}
	}
}
