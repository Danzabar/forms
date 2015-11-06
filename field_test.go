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
func TestFieldValidateRequired(t *testing.T) {
	req := &Required{
		Err: "This field is required",
	}

	field := NewField()
	field.Name = "test"
	field.Value = ""
	field.Type = "text"
	field.addValidation(req)

	field.validate()

	assert.Equal(t, 1, len(field.Errors))
}

func TestFieldValidateLength(t *testing.T) {
	req := &LengthValidation{
		Err: "This field should be 5 characters",
		Len: 5,
	}

	field := NewField()
	field.Name = "test"
	field.Value = ""
	field.Type = "text"
	field.addValidation(req)

	field.validate()

	assert.Equal(t, 1, len(field.Errors))
}

func TestFieldValidationRequired(t *testing.T) {
	req := &RegexValidation{
		Err:   "An error",
		Regex: "^([A-Z])$",
	}

	field := NewField()
	field.Name = "test"
	field.Value = "test"
	field.addValidation(req)

	field.validate()

	assert.Equal(t, 1, len(field.Errors))
}

// Testing class output
func TestAddingHtmlProperties(t *testing.T) {
	m := make(map[string]string)
	m["class"] = "testclass"
	m["data-id"] = "3"

	field := NewField()
	field.Props = m

	prop := field.getHtmlProperties()

	assert.Equal(t, "class=\"testclass\" data-id=\"3\" ", prop)
}
