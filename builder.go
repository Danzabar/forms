package forms

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// The builder struct
type Builder struct {
	// The location of the YAML file
	Filename string
	// The extracted configuration settings
	Config Config
	// The form struct that will be populated from the yaml file
	Form *Form
	// Instance of the http request
	Request http.Request
}

// The YAML config struct
type Config struct {
	Name        string            `yaml:"Name"`
	Description string            `yaml:"Description"`
	Action      string            `yaml:"Action"`
	Method      string            `yaml:"Method"`
	Props       map[string]string `yaml:"Props"`
	Fields      []ConfigField     `yaml:"Fields"`
}

// The YAML config struct for a field
type ConfigField struct {
	Name    string            `yaml:"Name"`
	Label   string            `yaml:"Label"`
	Type    string            `yaml:"Type"`
	Default string            `yaml:"Default"`
	Props   map[string]string `yaml:"Props"`
	Values  []string          `yaml:"Values"`
	Rules   []ConfigRule      `yaml:"Rules"`
}

// The validation struct
type ConfigRule struct {
	Type   string `yaml:"Type"`
	Err    string `yaml:"Err"`
	Regex  string `yaml:"Regex"`
	Length int    `yaml:"Length"`
}

// Creates a new form builder and runs the extract method
func NewBuilder(file string, r http.Request) *Builder {
	path, _ := filepath.Abs(file)

	b := &Builder{
		Filename: path,
		Request:  r,
	}

	b.extract()

	return b
}

// Extracts values from the YAML file and builds a config struct
func (b *Builder) extract() {
	data, _ := ioutil.ReadFile(b.Filename)

	var c Config

	yaml.Unmarshal([]byte(data), &c)

	b.Config = c
}

// Creates the form object from the configuration
func (b *Builder) build() {
	// Create the form struct
	b.Form = NewForm(b.Request)
	b.Form.Name = b.Config.Name
	b.Form.Description = b.Config.Description
	b.Form.Action = b.Config.Action
	b.Form.Props = b.Config.Props
	b.Form.Method = b.Config.Method

	// Now we can add fields
	for _, field := range b.Config.Fields {
		f := &Field{
			Name:    field.Name,
			Label:   field.Label,
			Type:    field.Type,
			Props:   field.Props,
			Default: field.Default,
			Values:  field.Values,
		}

		// Create validation items
		createValidationItems(field.Rules, f)

		b.Form.addField(f)
	}
}

// Loops through rules and returns a list of validation rules
func createValidationItems(rules []ConfigRule, f *Field) {
	for _, rule := range rules {

		var r Validation

		switch rule.Type {
		case "Required":
			r = &Required{
				Err: rule.Err,
			}
		case "Regex":
			r = &RegexValidation{
				Err:   rule.Err,
				Regex: rule.Regex,
			}
		case "Length":
			r = &LengthValidation{
				Err: rule.Err,
				Len: rule.Length,
			}
		}

		f.addValidation(r)
	}
}
