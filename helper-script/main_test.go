package main

import (
	"reflect"
	"testing"
)

func TestScrubText(t *testing.T) {
	// Table of test cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Simple clean", "  Hello  ", "Hello"},
		{"Newlines and tabs", "Line1\n\t\rLine2", "Line1 Line2"},
		{"Multiple spaces", "Word1    Word2", "Word1 Word2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := scrubText(tt.input)
			if actual != tt.expected {
				t.Errorf("scrubText() = %q, want %q", actual, tt.expected)
			}
		})
	}
}

func TestParseContent(t *testing.T) {
	input := []byte(`<div data-twomi18n="key1 key2[attr] key3">  Some Text  </div>`)
	expected := map[string]string{
		"key1": "Some Text",
		"key2": "",
		"key3": "Some Text",
	}

	actual := parseContent(input)

	// In Go, we use reflect.DeepEqual to compare maps
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("parseContent() = %v, want %v", actual, expected)
	}
}
