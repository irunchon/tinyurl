package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func (s Service) GetLongURL(ctx context.Context, request *pb.Hash) (*pb.LongURL, error)
//t.Run("", func(t *testing.T){})

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
