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

	assert.Equal(t, "<input type=\"text\" name=\"foo\" id=\"foo\" value=\"\" />", tag)
}

func TestFieldOutputTextArea(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Type = "textarea"

	tag := field.output()

	assert.Equal(t, "<textarea name=\"foo\" id=\"foo\"></textarea>", tag)
}

func TestFieldAddError(t *testing.T) {
	field := NewField()
	field.addError("A test error")

	assert.Equal(t, 1, len(field.Errors))
}
