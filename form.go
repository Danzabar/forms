package forms

import (
	"fmt"
	"net/http"
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
	Fields []*Field
	// The current request
	Request *http.Request
}

func NewForm(r *http.Request) *Form {
	return &Form{
		Name:        "Form",
		Description: "",
		Action:      "",
		Method:      "POST",
		Request:     r,
	}
}

func (f Form) open() string {
	return fmt.Sprintf("<form method=\"%s\" action=\"%s\">", f.Method, f.Action)
}

func (f Form) close() string {
	return fmt.Sprintf("</form>")
}
