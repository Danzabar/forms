package forms

import (
	"regexp"
)

type Validation interface {
	validate()
	setValue(val string)
	getErr() string
	isValid() bool
}

// Required Checks whether any value is set
type Required struct {
	Value string
	Err   string
	Valid bool
}

// Checks for a specific length
type LengthValidation struct {
	Value string
	Err   string
	Valid bool
	Len   int
}

// Custom Regex validation
type RegexValidation struct {
	Value string
	Err   string
	Valid bool
	Regex string
}

// Methods for the Required validation struct
func (r *Required) validate() {
	// We only follow the happy path here because go sets Valid as
	// its null value (false)
	if len(r.Value) > 0 {
		r.Valid = true
	}
}

func (r *Required) setValue(val string) {
	r.Value = val
}

func (r Required) getErr() string {
	return r.Err
}

func (r Required) isValid() bool {
	return r.Valid
}

// Methods for the Length validation struct
func (l *LengthValidation) validate() {
	if len(l.Value) == l.Len {
		l.Valid = true
	}
}

func (l *LengthValidation) setValue(val string) {
	l.Value = val
}

func (l LengthValidation) getErr() string {
	return l.Err
}

func (l LengthValidation) isValid() bool {
	return l.Valid
}

// Methods for the Regex Validation struct
func (r *RegexValidation) validate() {
	match, _ := regexp.MatchString(r.Regex, r.Value)
	r.Valid = match
}

func (r *RegexValidation) setValue(val string) {
	r.Value = val
}

func (r RegexValidation) getErr() string {
	return r.Err
}

func (r RegexValidation) isValid() bool {
	return r.Valid
}
