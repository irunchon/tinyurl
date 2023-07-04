package app

import (
	"context"
	"testing"

	"github.com/irunchon/tinyurl/internal/pkg/storage/inmemory"
	pb "github.com/irunchon/tinyurl/pkg/tinyurl/api"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestGetLongURL(t *testing.T) {
	var ctx = context.Background()

	t.Run("Error - wrong hash", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		hash := "1"
		request := &pb.GetLongURLRequest{Hash: hash}

		actual, err := testService.GetLongURL(ctx, request)

		assert.Nil(t, actual)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "requested URL is not valid")
	})
	t.Run("Error - URL not found in repo", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		hash := "1234567891"
		request := &pb.GetLongURLRequest{Hash: hash}

		actual, err := testService.GetLongURL(ctx, request)

		assert.Nil(t, actual)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "long URL is not found in repository")
	})
	t.Run("OK", func(t *testing.T) {
		testService := New(inmemory.NewInMemoryStorage())
		expectedURL := "https://go.dev/play/"
		hash := "Gr4vnVN37M"
		_ = testService.repo.SetShortAndLongURLs(hash, expectedURL)
		request := &pb.GetLongURLRequest{Hash: hash}

		actual, err := testService.GetLongURL(ctx, request)

		require.Nil(t, err)
		assert.Equal(t, expectedURL, actual.LongUrl)
	})
}

func TestIsHashValid(t *testing.T) {
	for _, tc := range []struct {
		name     string
		hash     string
		expected bool
	}{
		{
			name:     "Valid hash - only letters",
			hash:     "qwertyuiop",
			expected: true,
		},
		{
			name:     "Valid hash with digits",
			hash:     "Qwer5Aui91",
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
		{
			name:     "Invalid symbols in hash",
			hash:     "123456789$",
			expected: false,
		},
		{
			name:     "Invalid letters in hash",
			hash:     "йцукенгшщз",
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, isHashValid(tc.hash))
		})
	}
}
