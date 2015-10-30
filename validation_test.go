package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Tests the required validation
func TestRequiredValidation(t *testing.T) {
	req := &Required{
		Err: "This field is required",
	}

	req.validate()

	assert.Equal(t, false, req.Valid)

	req.Value = "test"
	req.validate()

	assert.Equal(t, true, req.Valid)
}

// Tests the length validation
func TestLengthValidation(t *testing.T) {
	val := &LengthValidation{
		Err: "This field should be 5 characters in length",
		Len: 5,
	}

	val.Value = "fail"
	val.validate()

	assert.Equal(t, false, val.Valid)

	val.Value = "tests"
	val.validate()

	assert.Equal(t, true, val.Valid)
}

/**
 *	Benchmarks
 */
// Required
func BenchmarkRequiredValidation(t *testing.B) {
	val := &Required{
		Err: "This field is required",
	}

	val.validate()
}

// Length
func BenchmarkLengthValidation(t *testing.B) {
	val := &LengthValidation{
		Err: "This field should be 5 characters",
		Len: 5,
	}

	val.validate()
}
