package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "      test     ",
			expected: []string{"test"},
		},
		{
			input:    "      test     ing   ",
			expected: []string{"test", "ing"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "nowhitespace",
			expected: []string{"nowhitespace"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "Uppercase text",
			expected: []string{"uppercase", "text"},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)
		if len(result) != len(c.expected) {
			t.Error("Length of result and expected result do not match.")
		}
		for i := range result {
			if result[i] != c.expected[i] {
				t.Error("Result and expected result do not match")
			}
		}
	}
}
