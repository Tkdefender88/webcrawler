package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputUrl string
		expected string
	}{
		{
			name:     "remove scheme",
			inputUrl: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "longer path",
			inputUrl: "https://blog.boot.dev/path/to/something",
			expected: "blog.boot.dev/path/to/something",
		},
		{
			name:     "query params ",
			inputUrl: "https://blog.boot.dev/path/to/something?query=params",
			expected: "blog.boot.dev/path/to/something",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalUrl, err := NormalizeURL(tt.inputUrl)
			require.NoError(t, err)
			require.Equal(t, tt.expected, normalUrl)
		})
	}
}
