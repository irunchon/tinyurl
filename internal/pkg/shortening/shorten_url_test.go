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
			expected: "RFOAIuuhZH",
		},
		{
			name:     "Hash for string '1'",
			url:      "1",
			expected: "BAllHGT1pM",
		},
		{
			name:     "Hash for string 'qwerty'",
			url:      "qwerty",
			expected: "SabvwfL9DM",
		},
		{
			name:     "Hash for url 'https://habr.com/'",
			url:      "https://habr.com/",
			expected: "QKUs71zwCU",
		},
		{
			name:     "Hash for url to habr article",
			url:      "https://habr.com/ru/articles/310460/",
			expected: "Id_h4FJtgx",
		},
		{
			name:     "Hash for long url on wikipedia",
			url:      "https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%BA%D0%BE%D0%B4%D0%BE%D0%B2_%D1%81%D0%BE%D1%81%D1%82%D0%BE%D1%8F%D0%BD%D0%B8%D1%8F_HTTP#404",
			expected: "K7Yy7Wh3Ih",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, GenerateHashForURL(tc.url))
		})
	}
}

func TestDecimalNumToBase63String(t *testing.T) {
	for _, tc := range []struct {
		name     string
		number   uint64
		expected string
	}{
		{
			name:     "Number 0",
			number:   0,
			expected: "A",
		},
		{
			name:     "Number 1",
			number:   1,
			expected: "B",
		},
		{
			name:     "Number 63",
			number:   63,
			expected: "BA",
		},
		{
			name:     "Number 64",
			number:   64,
			expected: "BB",
		},
		{
			name:     "Number 11157",
			number:   11157,
			expected: "CzG",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, decimalNumToBase63String(tc.number))
		})
	}
}

func TestDecimalNumberConversionToBaseNNumbers(t *testing.T) {
	for _, tc := range []struct {
		name     string
		number   uint64
		base     uint64
		expected []uint64
	}{
		{
			name:     "Number 123 to base 10",
			number:   123,
			base:     10,
			expected: []uint64{3, 2, 1},
		},
		{
			name:     "Number 123456 to base 8",
			number:   123456,
			base:     8,
			expected: []uint64{0, 0, 1, 1, 6, 3},
		},
		{
			name:     "Number 123 to base 2",
			number:   123,
			base:     2,
			expected: []uint64{1, 1, 0, 1, 1, 1, 1},
		},
		{
			name:     "Number 0 to base 8",
			number:   0,
			base:     8,
			expected: []uint64{0},
		},
		{
			name:     "Number 0 to base 0",
			number:   0,
			base:     0,
			expected: []uint64{0},
		},
		{
			name:     "Number 7 to base 8",
			number:   7,
			base:     8,
			expected: []uint64{7},
		},
		{
			name:     "Number 9 to base 8",
			number:   9,
			base:     8,
			expected: []uint64{1, 1},
		},
		{
			name:     "Number 10 to base 2",
			number:   10,
			base:     2,
			expected: []uint64{0, 1, 0, 1},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, decimalNumberConversionToBaseNNumbers(tc.number, tc.base))
		})
	}
}
