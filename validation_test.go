package forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequiredValidation(t *testing.T) {
	req := &Required{
		Err: "This field is required",
	}

	req.validate()

	assert.Equal(t, false, req.Valid)
}
