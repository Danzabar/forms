package forms

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
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

// Test building a form with pre-defined values
func TestBuildPreValues(t *testing.T) {
	r := http.Request{
		Method: "POST",
		Body:   ioutil.NopCloser(strings.NewReader("test=foo&test2=bar&test3=test1")),
		Header: http.Header{
			"Content-Type": {"application/x-www-form-urlencoded"},
		},
	}

	b := NewBuilder("example.yml", r)
	b.build()

	for _, field := range b.Form.Fields {
		switch field.Name {
		case "test":
			assert.Equal(t, "foo", field.Value)
		case "test2":
			assert.Equal(t, "bar", field.Value)
		case "test3":
			assert.Equal(t, "test1", field.Value)
		}
	}
}

// Build benchmark
func BenchmarkBuildFormFromConfig(t *testing.B) {
	b := NewBuilder("example.yml", request)

	b.build()
}
