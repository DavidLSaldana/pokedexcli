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
		//add more cases
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match actual:%d and expected: %d", len(actual), len(c.expected))
			continue

		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

		}
	}

}
