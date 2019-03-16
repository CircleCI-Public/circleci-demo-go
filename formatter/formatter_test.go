package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	result := Format("my string")
	assert.Equal(t, result, "Formatted: my string")
}
