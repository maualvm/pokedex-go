package main

import "testing"

type cleanInputTestCase struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	testCases := []cleanInputTestCase{
		{
			input:    "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "Foo BAR baZ",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, testCase := range testCases {
		got := cleanInput(testCase.input)

		if len(got) != len(testCase.expected) {
			t.Errorf("cleanInput returned %d strings, expected %d", len(got), len(testCase.expected))
		}

		for i := range got {
			word := got[i]
			expected := testCase.expected[i]

			if word != expected {
				t.Errorf("got %v; expected %v", got, testCase.expected)
			}
		}
	}
}
