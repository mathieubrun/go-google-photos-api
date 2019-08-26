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
		{&SchemaProperty{Type: STRING}, "string"},
		{&SchemaProperty{Ref: "SomeStruct"}, "*SomeStruct"},
		{&SchemaProperty{Type: ARRAY, Items: &SchemaPropertyItem{Type: STRING}}, "[]string"},
		{&SchemaProperty{Type: ARRAY, Items: &SchemaPropertyItem{Ref: "SomeStruct"}}, "[]*SomeStruct"},
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
		{&SchemaProperty{Type: "invalid"}, &ErrUnknownJsonType{}},
		{&SchemaProperty{Type: ARRAY, Items: &SchemaPropertyItem{Type: "invalid"}}, &ErrUnknownJsonType{}},
	}

	for _, tt := range invalidInputs {
		t.Run("invalidInputs", func(t *testing.T) {
			s, err := tt.in.GetType()
			assert.Equal(t, s, "")
			assert.ErrorType(t, err, tt.err)
		})
	}
}

func Test_getTypeFronJsonType(t *testing.T) {
	var validInputs = []struct {
		in  string
		out string
	}{
		{"integer", "int64"},
		{"string", "string"},
		{"number", "float64"},
		{"boolean", "bool"},
		{"object", "googleapi.RawMessage"},
	}

	for _, tt := range validInputs {
		t.Run(tt.in, func(t *testing.T) {
			s, err := getTypeFromJsonType(tt.in)
			assert.Equal(t, s, tt.out)
			assert.NilError(t, err)
		})
	}

	var invalidInputs = []struct {
		in  string
		err error
	}{
		{"unknown", &ErrUnknownJsonType{}},
	}

	for _, tt := range invalidInputs {
		t.Run(tt.in, func(t *testing.T) {
			s, err := getTypeFromJsonType(tt.in)
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
