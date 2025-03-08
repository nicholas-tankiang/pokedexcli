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
			input:    "  hello    world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " one    two -three-    /@/ ",
			expected: []string{"one", "two", "-three-", "/@/"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("ERROR raw: %q// processed len: %d// expected len: %d", c.input, len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("ERROR index %d // returned word: %q // expected word: %q", i, word, expectedWord)
			}
		}
	}
}
