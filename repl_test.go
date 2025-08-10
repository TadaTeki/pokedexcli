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
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		// 長さ比較
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch: expected %d, got %d", len(c.expected), len(actual))
			continue // 長さが違えば以降の比較はできない
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("word mismatch at index %d: expected %q, got %q", i, c.expected[i], actual[i])
			}
		}
	}

}
