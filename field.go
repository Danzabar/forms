package forms

import (
	"fmt"
	"strings"
)

type Field struct {
	// The field name
	Name string
	// The field label
	Label string
	// The field type ie text, textarea
	Type string
	// The default value
	Default string
	// The actual value
	Value string
	// A list of values to use for multi-value inputs
	Values []string
	// A list of validation items to add to this field
	Rules []Validation
	// A list of errors on this field
	Errors []string
}

// Creates a new field struct
func NewField() *Field {
	return &Field{
		Name: "Field",
		Type: "text",
	}
}

// Returns the string representation of a field, depending on its type
func (f Field) output() string {
	switch f.Type {
	case "textarea":
		return f.outputTextArea()
	case "radio":
	case "checkbox":
		return f.outputBox()
	case "select":
		return f.outputSelect()
	}

	return f.outputText()
}

// Adds an error to the field
func (f *Field) addError(err string) {
	f.Errors = append(f.Errors, err)
}

// Runs all the validation associated with this field
func (f *Field) validate() {
	for _, rule := range f.Rules {
		rule.setValue(f.Value)
		rule.validate()

		if !rule.isValid() {
			f.addError(rule.getErr())
		}
	}
}

// Returns the label for a field
func (f Field) label() string {
	return fmt.Sprintf("<label for=\"%s\">%s</label>", f.Name, f.Label)
}

// Returns the string representation of a standard text input
func (f Field) outputText() string {
	return fmt.Sprintf("<input type=\"%[1]s\" name=\"%[2]s\" id=\"%[2]s\" value=\"%[3]s\" />", f.Type, f.Name, f.Value)
}

// Returns the string representation of a textarea input
func (f Field) outputTextArea() string {
	return fmt.Sprintf("<textarea name=\"%[1]s\" id=\"%[1]s\">%[2]s</textarea>", f.Name, f.Value)
}

// Returns the string representation of a checkbox/radio
// Currently, this is exactly the same as outputText, however there
// Will be some difference once request form values are added
func (f Field) outputBox() string {
	return fmt.Sprintf("<input type=\"%[1]s\" name=\"%[2]s\" id=\"%[2]s\" value=\"%[3]s\" />", f.Type, f.Name, f.Value)
}

// Returns the string representation of a select box
func (f Field) outputSelect() string {
	var options []string

	for _, val := range f.Values {
		options = append(options, fmt.Sprintf("<option value=\"%[1]s\">%[1]s</option>", val))
	}

	return fmt.Sprintf("<select name=\"%[1]s\" id=\"%[1]s\">%[2]s</select>", f.Name, strings.Join(options, ""))
}
