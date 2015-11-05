package forms

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
}

// The YAML config struct
type Config struct {
	Name        string `yaml:"Name"`
	Description string `yaml:"Description"`
	Action      string `yaml:"Action"`
	Method      string `yaml:"Method"`
}

func NewBuilder(file string) *Builder {
	path, _ := filepath.Abs(file)

	b := &Builder{
		Filename: path,
	}
	// Extract settings
	b.extract()

	return b
}

func (b *Builder) extract() {
	data, _ := ioutil.ReadFile(b.Filename)

	var c Config

	yaml.Unmarshal([]byte(data), &c)

	b.Config = c
}
