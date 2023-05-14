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

	t.Run("Error - wrong URL", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		url := ""
		request := &pb.GetShortURLRequest{LongUrl: url}

		actual, err := testService.GetShortURL(ctx, request)

		assert.Nil(t, actual)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requested URL is not valid")
	})
	t.Run("OK - hash already exists in repo", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		url := "https://go.dev/play/"
		expectedHash := "E3puYxWn1Q"
		testService.repo.SetShortAndLongURLs(expectedHash, url)
		request := &pb.GetShortURLRequest{LongUrl: url}

		actual, err := testService.GetShortURL(ctx, request)

		require.Nil(t, err)
		assert.Equal(t, expectedHash, actual.ShortUrl)
	})
	t.Run("OK - new hash generated", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		url := "https://github.com/"
		expectedHash := "R2oCwKKhF6"
		request := &pb.GetShortURLRequest{LongUrl: url}

		actual, err := testService.GetShortURL(ctx, request)

		require.Nil(t, err)
		assert.Equal(t, expectedHash, actual.ShortUrl)
	})
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

func TestGenerateUniqueHashForURL(t *testing.T) {
	t.Run("Hash not found in repo", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())

		actual, err := testService.generateUniqueHashForURL("test")

		assert.Nil(t, err)
		assert.Equal(t, "O1CYK5Hql0", actual)
	})

	t.Run("Hash already exists in repo", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		testService.repo.SetShortAndLongURLs("O1CYK5Hql0", "test")

		actual, err := testService.generateUniqueHashForURL("test")

		assert.Equal(t, "O1CYK5Hql0", actual)
		assert.Equal(t, errorAlreadyExists, err)
	})
	t.Run("Hash shifting doesn't help", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		testService.repo.SetShortAndLongURLs("O1CYK5Hql0", "test0")
		testService.repo.SetShortAndLongURLs("1CYK5Hql0O", "test1")
		testService.repo.SetShortAndLongURLs("CYK5Hql0O1", "test2")
		testService.repo.SetShortAndLongURLs("YK5Hql0O1C", "test3")
		testService.repo.SetShortAndLongURLs("K5Hql0O1CY", "test4")
		testService.repo.SetShortAndLongURLs("5Hql0O1CYK", "test5")
		testService.repo.SetShortAndLongURLs("Hql0O1CYK5", "test6")
		testService.repo.SetShortAndLongURLs("ql0O1CYK5H", "test7")
		testService.repo.SetShortAndLongURLs("l0O1CYK5Hq", "test8")
		testService.repo.SetShortAndLongURLs("0O1CYK5Hql", "test9")

		actual, err := testService.generateUniqueHashForURL("test")

		assert.Equal(t, "", actual)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "fail to generate hash")
	})
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
