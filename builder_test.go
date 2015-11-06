package forms

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var request http.Request

// Testing basic parsing of YAML form
func TestParseYamlFile(t *testing.T) {
	b := NewBuilder("example.yml", request)

	assert.Equal(t, "Test", b.Config.Name)
}

// Test building a usable form from the yaml file
func TestBuildingForm(t *testing.T) {
	b := NewBuilder("example.yml", request)

	b.build()
	assert.Equal(t, 4, len(b.Form.Fields))
}

// Build benchmark
func BenchmarkBuildFormFromConfig(t *testing.B) {
	b := NewBuilder("example.yml", request)

	b.build()
}
