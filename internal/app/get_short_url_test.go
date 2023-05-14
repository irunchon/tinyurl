package app

import (
	"context"
	"testing"

	"github.com/irunchon/tinyurl/internal/pkg/storage/inmemory"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetShortURL(t *testing.T) {
	var ctx = context.Background()

	t.Run("Hash for wrong URL", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		url := ""
		request := &pb.LongURL{LongUrl: url}

		actual, err := testService.GetShortURL(ctx, request)

		assert.Nil(t, actual)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requested URL is not valid")
	})

	//t.Run("", func(t *testing.T) {})
	//t.Run("", func(t *testing.T) {})
}

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