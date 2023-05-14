package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUrl(t *testing.T) {
	for _, tc := range []struct {
		name     string
		str      string
		expected bool
	}{
		{
			name:     "Valid URL - http",
			str:      "http://google.com/",
			expected: true,
		},
		{
			name:     "Valid URL - http (no slash)",
			str:      "http://google.com",
			expected: true,
		},
		{
			name:     "Valid URL - https (long)",
			str:      "https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go",
			expected: true,
		},
		{
			name:     "Valid URL - https (localhost)",
			str:      "https://localhost/",
			expected: true,
		},
		{
			name:     "Valid URL - ftp",
			str:      "ftp://mystorage.com/",
			expected: true,
		},
		{
			name:     "Invalid URL - empty string",
			str:      "",
			expected: false,
		},
		{
			name:     "Invalid URL - missing symbol",
			str:      "http//google.com",
			expected: false,
		},
		{
			name:     "Invalid URL - no protocol",
			str:      "google.com",
			expected: false,
		},
		{
			name:     "Invalid URL - words with slashes",
			str:      "/foo/bar",
			expected: false,
		},
		{
			name:     "Invalid URL - word localhost",
			str:      "localhost",
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, IsUrl(tc.str))
		})
	}
}

func TestIsHashValid(t *testing.T) {
	for _, tc := range []struct {
		name     string
		hash     string
		expected bool
	}{
		{
			name:     "Valid hash",
			hash:     "qwertyuiop",
			expected: true,
		},
		{
			name:     "Invalid hash - empty",
			hash:     "",
			expected: false,
		},
		{
			name:     "Invalid hash - short",
			hash:     "1",
			expected: false,
		},
		{
			name:     "Invalid hash - long",
			hash:     "asdfghjklzxcvbnm",
			expected: false,
		},
		// TODO: more tests when IsHashValid func finished (symbols checks)
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, IsHashValid(tc.hash))
		})
	}
}

func TestHashRingShift(t *testing.T) {
	for _, tc := range []struct {
		name     string
		hash     string
		expected string
	}{
		{
			name:     "qwerty",
			hash:     "qwerty",
			expected: "wertyq",
		},
		{
			name:     "one symbol",
			hash:     "1",
			expected: "1",
		},
		{
			name:     "same symbols",
			hash:     "qq",
			expected: "qq",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, hashRingShift(tc.hash))
		})
	}
}
