package shortening

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHashForURL(t *testing.T) {
	for _, tc := range []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "Hash for empty string",
			url:      "",
			expected: "zTXKKJACG4",
		},
		{
			name:     "Hash for string '1'",
			url:      "1",
			expected: "4Ga3TjdhVd",
		},
		{
			name:     "Hash for string 'qwerty'",
			url:      "qwerty",
			expected: "24ockDRqPA",
		},
		{
			name:     "Hash for url 'https://habr.com/'",
			url:      "https://habr.com/",
			expected: "wVQAYJ3G7Y",
		},
		{
			name:     "Hash for url to habr article",
			url:      "https://habr.com/ru/articles/310460/",
			expected: "VXMwTeb6Mb",
		},
		{
			name:     "Hash for long url on wikipedia",
			url:      "https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%BA%D0%BE%D0%B4%D0%BE%D0%B2_%D1%81%D0%BE%D1%81%D1%82%D0%BE%D1%8F%D0%BD%D0%B8%D1%8F_HTTP#404",
			expected: "dXwf4pGtrq",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, GenerateHashForURL(tc.url))
		})
	}
}

func TestDecimalNumToBaseNString(t *testing.T) {
	for _, tc := range []struct {
		name     string
		number   uint64
		expected string
	}{
		{
			name:     "Number 0",
			number:   0,
			expected: "",
		},
		{
			name:     "Number 56",
			number:   56,
			expected: "21",
		},
		{
			name:     "Number 100",
			number:   100,
			expected: "2n",
		},
		{
			name:     "Number 11157",
			number:   11157,
			expected: "4ZE",
		},
		{
			name:     "Number 123456789",
			number:   123456789,
			expected: "DYzbX",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, decimalNumToBaseNString(tc.number))
		})
	}
}
