package forms

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"strings"
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

// Tests the output of the form tags
func (suite *FormTestSuite) TestFormTags() {
	form := NewForm(suite.Request)
	form.Action = "/test/uri"
	form.Method = "GET"

	oTag := form.open()
	cTag := form.close()

	assert.Equal(suite.T(), "<form method=\"GET\" action=\"/test/uri\">", oTag)
	assert.Equal(suite.T(), "</form>", cTag)
}

// Tests adding a field with no value
func (suite *FormTestSuite) TestAddFieldNoValue() {
	form := NewForm(suite.Request)
	form.Action = "/test/uri"
	form.Method = "POST"

	field := &Field{
		Name:  "test",
		Label: "Test",
		Type:  "text",
	}

	form.addField(field)
	assert.Equal(suite.T(), 1, len(form.Fields))
}

// Tests adding a field with a value set in the request
func (suite *FormTestSuite) TestAddFieldWithValue() {

	req := http.Request{
		Method: "POST",
		Body:   ioutil.NopCloser(strings.NewReader("foo=value")),
		Header: http.Header{"Content-Type": {"application/x-www-form-urlencoded;"}},
	}

	form := NewForm(req)
	form.Action = "/test/uri"
	form.Method = "POST"

	field := &Field{
		Name:  "foo",
		Label: "Foo field",
		Type:  "text",
	}

	form.addField(field)
	assert.Equal(suite.T(), "value", field.Value)
}

func TestRunnerTestSuite(t *testing.T) {
	suite.Run(t, new(FormTestSuite))
}
