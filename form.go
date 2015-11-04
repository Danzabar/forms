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
	Request http.Request
	// Boolean flag whether form is valid or not
	Valid bool
}

// Creates the form struct
func NewForm(r http.Request) *Form {
	return &Form{
		Name:        "Form",
		Description: "",
		Action:      "",
		Method:      "POST",
		Request:     r,
		Valid:       false,
	}
}

// Adds a field and checks for a value in the request
func (f *Form) addField(field *Field) {
	// Check for a value in the request obj
	val := f.Request.FormValue(field.Name)
	field.Value = val

	f.Fields = append(f.Fields, field)
}

// Returns the form open tags
func (f Form) open() string {
	return fmt.Sprintf("<form method=\"%s\" action=\"%s\">", f.Method, f.Action)
}

// Returns the form close tags
func (f Form) close() string {
	return fmt.Sprintf("</form>")
}

// Method to validate the whole form
func (f *Form) validate() bool {
	f.Valid = true
	for _, field := range f.Fields {
		field.validate()

		if !field.Valid {
			f.Valid = false
		}
	}

	return f.Valid
}
