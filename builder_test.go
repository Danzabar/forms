package forms

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

var request http.Request

// Test suite
type BuilderTestSuite struct {
	suite.Suite
	Req http.Request
}

// Test setup method
func (suite *BuilderTestSuite) SetupTest() {
	suite.Req = http.Request{
		Method: "POST",
	}
}

// Testing basic parsing of YAML form
func (suite *BuilderTestSuite) TestParseYamlFile() {
	b := NewBuilder("example.yml", suite.Req)

	assert.Equal(suite.T(), "Test", b.Config.Name)
}

// Test building a usable form from the yaml file
func (suite *BuilderTestSuite) TestBuildingForm() {
	b := NewBuilder("example.yml", suite.Req)

	b.build()
	assert.Equal(suite.T(), 2, len(b.Form.Fields))
}

// Test runner
func BuilderTestRunnerTestSuite(t *testing.T) {
	suite.Run(t, new(BuilderTestSuite))
}

// Build benchmark
func BenchmarkBuildFormFromConfig(t *testing.B) {
	b := NewBuilder("example.yml", request)

	b.build()
}
