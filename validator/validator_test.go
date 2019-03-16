package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	result := Validate("my string")
	assert.True(t, result)
}
