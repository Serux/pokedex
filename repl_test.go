package main

import "testing"

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
			input:    "  Charmander  Bulbasur PIKACHU  ",
			expected: []string{"charmander", "bulbasur", "pikachu"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("GOT: %v , EXPECTED:  %v", len(actual), len(c.expected))

		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			t.Log(word, "---", expectedWord)

			if word != expectedWord {
				t.Errorf("GOT: %v , EXPECTED:  %v", word, expectedWord)

			}
		}
	}
}
