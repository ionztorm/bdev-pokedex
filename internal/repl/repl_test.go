package repl

import (
	"pokedex/internal/utils"
	"testing"
)

type testCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	cases := []testCase{
		{"TEST LOWER CASE 1", []string{"test", "lower", "case", "1"}},
		{"TeSt LoWeR CaSe 2", []string{"test", "lower", "case", "2"}},
		{"  Test preceeding white space", []string{"test", "preceeding", "white", "space"}},
		{"Test post whitespace    ", []string{"test", "post", "whitespace"}},
		{"", []string{}},
		{"    ", []string{}},
	}

	for _, c := range cases {
		output := utils.CleanInput(c.input)
		if len(output) != len(c.expected) {
			t.Errorf("For input: %s\nExpected length of slice to be %d but got %d.", c.input, len(c.expected), len(output))
		}

		for i := range c.expected {
			word := output[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("For input: '%s'\nAt index %d: Expected '%s' but got '%s'", c.input, i, c.expected[i], output[i])
			}
		}
	}
}
