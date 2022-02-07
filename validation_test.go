package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const validInput = "[()({[{}]})]()"

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid Setup",
			input:    validInput,
			expected: true,
		},
		{
			name:     "Invalid Setup",
			input:    "[{()[([)({[]})]{[]}",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, validate(tt.input, len(tt.input)))
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		validate(validInput, len(validInput))
	}
}
