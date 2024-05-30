package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {
	hash1 := GenerateShortURL("www.example.com")
	hash2 := GenerateShortURL("www.example.com")
	assert.Equal(t, hash1, hash2, "hashes should be same")

	hash1 = GenerateShortURL("www.sample.com")
	hash2 = GenerateShortURL("www.example.com")
	assert.NotEqual(t, hash1, hash2, "hashes should not be same")
}
