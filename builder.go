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
	Name        string        `yaml:"Name"`
	Description string        `yaml:"Description"`
	Action      string        `yaml:"Action"`
	Method      string        `yaml:"Method"`
	Fields      []ConfigField `yaml:"Fields"`
}

// The YAML config struct for a field
type ConfigField struct {
	Name    string   `yaml:"Name"`
	Label   string   `yaml:"Label"`
	Type    string   `yaml:"Type"`
	Default string   `yaml:"Default"`
	Values  []string `yaml:"Values"`
	Rules   []string `yaml:"Rules"`
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
	b.Form.Method = b.Config.Method

	// Now we can add fields
	for _, field := range b.Config.Fields {
		f := &Field{
			Name:    field.Name,
			Label:   field.Label,
			Type:    field.Type,
			Default: field.Default,
			Values:  field.Values,
		}

		b.Form.addField(f)
	}
}
