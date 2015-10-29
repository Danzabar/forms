package forms

import (
	"fmt"
)

type Field struct {
	// The field name
	Name string
	// The field type ie text, textarea
	Type string
	// The default value
	Default string
	// The actual value
	Value string
	// A list of errors on this field
	Errors []string
}

func NewField() *Field {
	return &Field{
		Name: "Field",
		Type: "text",
	}
}

func (f Field) output() string {
	switch f.Type {
	case "textarea":
		return f.outputTextArea()
	}

	return f.outputText()
}

func (f *Field) addError(err string) {
	f.Errors = append(f.Errors, err)
}

func (f Field) outputText() string {
	return fmt.Sprintf("<input type=\"%[1]s\" name=\"%[2]s\" id=\"%[2]s\" value=\"%[3]s\" />", f.Type, f.Name, f.Value)
}

func (f Field) outputTextArea() string {
	return fmt.Sprintf("<textarea name=\"%[1]s\" id=\"%[1]s\">%[2]s</textarea>", f.Name, f.Value)
}
