package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		input    string
		expected Counter
	}{
		{"Hello World\n", Counter{bytes: 12, chars: 12, words: 2, lines: 1}},
		{"One two three\nFour five six\n", Counter{words: 6, lines: 2, bytes: 28, chars: 28}},
		{"SingleWord", Counter{words: 1, lines: 0, bytes: 10, chars: 10}},
		{"Multiple\nNew\nLines\n", Counter{words: 3, lines: 3, bytes: 19, chars: 19}},
		{"", Counter{words: 0, lines: 0, bytes: 0, chars: 0}},
		// Test for unicode characters - 4 characters (3bytes per character)
		{"你好 世界\n", Counter{words: 2, lines: 1, bytes: 14, chars: 6}},
	}

	for _, test := range tests {
		reader := bufio.NewReader(strings.NewReader(test.input))
		counter := Counter{}
		count(reader, &counter)

		if counter != test.expected {
			t.Errorf("Expected %v, got %v - text: %v", test.expected, counter, test.input)
		}
	}
}
