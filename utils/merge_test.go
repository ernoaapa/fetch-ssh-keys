package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKeys(t *testing.T) {
	result := MergeKeys(map[string][]string{
		"user-1": []string{"key1"},
		"user-2": []string{"key1"},
	}, map[string][]string{
		"user-2": []string{"key2"},
		"user-3": []string{"key1"},
	})

	assert.Equal(t, map[string][]string{
		"user-1": []string{"key1"},
		"user-2": []string{"key1", "key2"},
		"user-3": []string{"key1"},
	}, result)
}
