package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormTags(t *testing.T) {
	form := NewForm()
	form.Action = "/test/uri"
	form.Method = "GET"

	oTag := form.open()
	cTag := form.close()

	assert.Equal(t, oTag, "<form method=\"GET\" action=\"/test/uri\">")
	assert.Equal(t, cTag, "</form>")
}
