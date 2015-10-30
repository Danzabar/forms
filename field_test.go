package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Testing the output of a text field
func TestFieldOutputText(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Type = "text"

	tag := field.output()

	assert.Equal(t, "<input type=\"text\" name=\"foo\" id=\"foo\" value=\"\" />", tag)
}

// Testing the output of a checkbox/radio
func TestFieldOutputBox(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Type = "checkbox"

	tag := field.output()

	assert.Equal(t, "<input type=\"checkbox\" name=\"foo\" id=\"foo\" value=\"\" />", tag)
}

// Testing the output of a textarea field
func TestFieldOutputTextArea(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Type = "textarea"

	tag := field.output()

	assert.Equal(t, "<textarea name=\"foo\" id=\"foo\"></textarea>", tag)
}

// Testing the output of a select field
func TestFieldOutputSelect(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Type = "select"
	field.Values = []string{"test"}

	tag := field.output()

	assert.Equal(t, "<select name=\"foo\" id=\"foo\"><option value=\"test\">test</option></select>", tag)
}

// Testing the output of a label
func TestFieldOutputLabel(t *testing.T) {
	field := NewField()
	field.Name = "foo"
	field.Label = "Foo"

	assert.Equal(t, "<label for=\"foo\">Foo</label>", field.label())
}

// Testing the add error method
func TestFieldAddError(t *testing.T) {
	field := NewField()
	field.addError("A test error")

	assert.Equal(t, 1, len(field.Errors))
}

// Testing validation on fields
func TestFieldValidate(t *testing.T) {
	req := &Required{
		Err: "This field is required",
	}

	field := NewField()
	field.Name = "test"
	field.Value = ""
	field.Rules = append(field.Rules, req)
	field.Type = "text"

	field.validate()

	assert.Equal(t, 1, len(field.Errors))
}
