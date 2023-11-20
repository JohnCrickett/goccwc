package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCalculateStats(t *testing.T) {
	cases := []struct {
		Description string
		input       string
		Want        stats
	}{
		{"Empty", "", stats{bytes: 0, words: 0, lines: 0, chars: 0}},
		{"Single char", "s", stats{bytes: 1, words: 1, lines: 0, chars: 1}},
		{"Multibyte chars", "s⌘ f", stats{bytes: 6, words: 2, lines: 0, chars: 4}},
		{"Trailing newline", "this is a sentence\n\nacross multiple\nlines\n", stats{bytes: 42, words: 7, lines: 4, chars: 42}},
		{"No trailing newline", "this is a sentence\n\nacross multiple\nlines", stats{bytes: 41, words: 7, lines: 3, chars: 41}},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			bufferedString := bufio.NewReader(strings.NewReader(test.input))
			got := CalculateStats(bufferedString)

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}
