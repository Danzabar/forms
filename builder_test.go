package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseYamlFile(t *testing.T) {
	b := NewBuilder("example.yml")

	assert.Equal(t, "Test", b.Config.Name)
}
