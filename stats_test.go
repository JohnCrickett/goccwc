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

func TestFormatStats(t *testing.T) {
	cases := []struct {
		Description   string
		inputStats    stats
		inputOptions  Options
		inputFilename string
		Want          string
	}{
		{"Empty", stats{bytes: 0, words: 0, lines: 0, chars: 0}, Options{true, true, true, false}, "", "0\t0\t0\t"},
		{"None Selected", stats{bytes: 11, words: 2, lines: 1, chars: 0}, Options{false, false, false, false}, "", ""},
		{"Default", stats{bytes: 11, words: 2, lines: 1, chars: 0}, Options{true, true, true, false}, "filename", "1\t2\t11\tfilename"},
		{"Chars", stats{bytes: 0, words: 0, lines: 0, chars: 100}, Options{false, false, false, true}, "filename", "100\tfilename"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := formatStats(test.inputOptions, test.inputStats, test.inputFilename)

			if got != test.Want {
				t.Errorf("got %v, want %v", got, test.Want)
			}
		})
	}
}

func TestCalculateStatsWithTotals(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("Hello, World\nLine 2\n"))
	reader2 := bufio.NewReader(strings.NewReader("Hello, World\nLine 2\nLine 3\n"))
	options := Options{true, true, true, false}
	expectedTotals := stats{bytes: 47, words: 10, lines: 5}

	var totals stats
	CalculateStatsWithTotals(reader, "Test File", options, &totals)
	CalculateStatsWithTotals(reader2, "Test File2", options, &totals)

	if totals != expectedTotals {
		t.Errorf("got %v, want %v", totals, expectedTotals)
	}
}
