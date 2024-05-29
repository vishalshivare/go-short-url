package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestURL(t *testing.T) {
	app := RequestURL{}
	err := app.Validate()
	assert.EqualError(t, err, "field 'URL' can not be empty")
}
