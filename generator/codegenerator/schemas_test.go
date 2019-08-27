package codegenerator

import (
	"gotest.tools/assert"
	"testing"
)

func Test_getPropertyType(t *testing.T) {
	var validInputs = []struct {
		in  *SchemaProperty
		out string
	}{
		{&SchemaProperty{PropertyDefinition: PropertyDefinition{Type: "string"}}, "string"},
		{&SchemaProperty{PropertyDefinition: PropertyDefinition{Ref: "SomeStruct"}}, "*SomeStruct"},
		{&SchemaProperty{PropertyDefinition: PropertyDefinition{Type: "array"}, Items: &SchemaPropertyItem{PropertyDefinition: PropertyDefinition{Type: "string"}}}, "[]string"},
		{&SchemaProperty{PropertyDefinition: PropertyDefinition{Type: "array"}, Items: &SchemaPropertyItem{PropertyDefinition: PropertyDefinition{Ref: "SomeStruct"}}}, "[]*SomeStruct"},
	}

	for _, tt := range validInputs {
		t.Run("validInputs", func(t *testing.T) {
			s, err := tt.in.GetType()
			assert.Equal(t, s, tt.out)
			assert.NilError(t, err)
		})
	}

	var invalidInputs = []struct {
		in  *SchemaProperty
		err error
	}{
		{&SchemaProperty{PropertyDefinition: PropertyDefinition{Type: "invalid"}}, &ErrUnknownJsonType{}},
		{&SchemaProperty{PropertyDefinition: PropertyDefinition{Type: "array"}, Items: &SchemaPropertyItem{PropertyDefinition: PropertyDefinition{Type: "invalid"}}}, &ErrUnknownJsonType{}},
	}

	for _, tt := range invalidInputs {
		t.Run("invalidInputs", func(t *testing.T) {
			s, err := tt.in.GetType()
			assert.Equal(t, s, "")
			assert.ErrorType(t, err, tt.err)
		})
	}
}

func Test_getPropertyDescription(t *testing.T) {
	var validInputs = []struct {
		in  *SchemaProperty
		out string
	}{
		{&SchemaProperty{Description: "oneline"}, "// oneline  "},
		{&SchemaProperty{Description: "two\nlines"}, "// two\n  // lines  "},
	}

	for _, tt := range validInputs {
		t.Run("validInputs", func(t *testing.T) {
			s := tt.in.GetDescription()
			assert.Equal(t, s, tt.out)
		})
	}
}
