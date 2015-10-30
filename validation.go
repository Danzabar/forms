package forms

type Validation interface {
	validate()
}

// Required Checks whether any value is set
type Required struct {
	Value string
	Err   string
	Valid bool
}

func (r *Required) validate() {
	// We only follow the happy path here because go sets Valid as
	// its null value (false)
	if len(r.Value) > 0 {
		r.Valid = true
	}
}
