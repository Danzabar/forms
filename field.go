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
