Forms
=====
[![Build Status](https://travis-ci.org/Danzabar/forms.svg)](https://travis-ci.org/Danzabar/forms) [![Coverage Status](https://coveralls.io/repos/Danzabar/forms/badge.svg?branch=master&service=github)](https://coveralls.io/github/Danzabar/forms?branch=master)

A little library that allows you to control web forms through Go code. Currently in development

## Usage
The following is a basic example of how to set up and use a go form

    // Create the form
    form := NewForm(http.Request)
    form.Action = "/test/form/uri"
    form.Method = "POST"

    // Add fields
    uname := &Field{
        Name:   "username",
        Label:  "Username",
        Type:   "text",
    }

    form.addField(uname)

### YAML Form builder
You can create a yaml file to represent a form, passing this to builder will create a form with all its fields and validation rules specified, firstly you'll need to make a yaml form file, take a look at example.yml in the source code, after this pass it to the builder:

    // Create a new builder
    b := NewBuilder("example.yml", http.Request)
    // Now the form can be built
    b.build()

    // Once the form is built you can access it through the builder
    b.Form
    b.Form.validate()

### Output
The form struct will output the form tags, and each field can output the Html representation of itself, for example

    form.open()

        for _, field := range form.Fields {
            field.output()
        }

    form.close()

### Validate
The form struct can also validate all fields it has, or you can validate individual fields...

    // Upon adding a field the method will check the request
    // object for the corresponding field value
    form.validate()

    // You can validate a single field as well
    field.validate()

Both the form and field struct have the `Valid` boolean flag to signify whether they have passed validation or not

    form.validate()

    // true
    return form.Valid

### Validation rules
Remember to add the validation you need to the field, otherwise the result of `validate` will always be true!

    rule := &Required{
        Err: "This field is required!"
    }

    field.addValidation(rule)

In the above example, if the field has no value and so fails validation it will add an entry into the fields `Error` list.
