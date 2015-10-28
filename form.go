package forms

import (
	"fmt"
)

type Form struct {
	// The form name
	Name string
	// The form description
	Description string
	// The form action
	Action string
	// The HTTP Method
	Method string
	// A list of fields
	Fields []Field
}

func NewForm() *Form {
	return &Form{
		Name:        "Form",
		Description: "",
		Action:      "",
		Method:      "POST",
	}
}

func (f Form) open() string {
	return fmt.Sprintf("<form method=\"%s\" action=\"%s\">", f.Method, f.Action)
}

func (f Form) close() string {
	return fmt.Sprintf("</form>")
}
