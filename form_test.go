package forms

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

// The test suite struct
type FormTestSuite struct {
	suite.Suite
	Request http.Request
}

// Set up test env
func (suite *FormTestSuite) SetupTest() {
	suite.Request = http.Request{
		Method: "POST",
	}
}

func (suite *FormTestSuite) TestFormTags() {
	form := NewForm(suite.Request)
	form.Action = "/test/uri"
	form.Method = "GET"

	oTag := form.open()
	cTag := form.close()

	assert.Equal(suite.T(), "<form method=\"GET\" action=\"/test/uri\">", oTag)
	assert.Equal(suite.T(), "</form>", cTag)
}

func TestRunnerTestSuite(t *testing.T) {
	suite.Run(t, new(FormTestSuite))
}
