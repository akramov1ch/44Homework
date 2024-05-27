package wordcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{"Empty string", "", map[string]int{}},
		{"Single word", "hello", map[string]int{"hello": 1}},
		{"Two words", "hello world", map[string]int{"hello": 1, "world": 1}},
		{"Repeated words", "a b a b c 2 f", map[string]int{"a": 2, "b": 2, "c": 1, "2": 1, "f": 1}},
		{"Complex string", "hello hello world world world", map[string]int{"hello": 2, "world": 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordCount(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
