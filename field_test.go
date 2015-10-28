package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldOutputText(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Type = "text"

	tag := field.output()

	assert.Equal(t, tag, "<input type=\"text\" name=\"foo\" id=\"foo\" value=\"\" />")
}

func TestFieldOutputTextArea(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Type = "textarea"

	tag := field.output()

	assert.Equal(t, tag, "<textarea name=\"foo\" id=\"foo\"></textarea>")
}
