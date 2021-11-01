package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopWord(t *testing.T) {
	testTable := []struct {
		text     string
		n        int
		expected []string
	}{
		{
			text:     "Hello, world!",
			n:        2,
			expected: []string{"Hello", "world"},
		},
		{
			text:     "I run, I run, I ran.",
			n:        1,
			expected: []string{"I"},
		},
		{
			text:     "I run, I run, I ran.",
			n:        2,
			expected: []string{"I", "run"},
		},
	}

	for _, testCase := range testTable {
		result := topWords(testCase.text, testCase.n)

		assert.Equal(t, testCase.expected, result,
			fmt.Errorf("Incorrect result. Expected: %v, Got: %v", testCase.expected, result))
	}
}
