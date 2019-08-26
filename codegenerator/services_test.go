package codegenerator

import (
	"gotest.tools/assert"
	"testing"
)

func Test_GetResourceName(t *testing.T) {
	var validInputs = []struct {
		in  string
		out string
	}{
		{"someValue", "SomeValue"},
		{"SomeValue", "SomeValue"},
	}

	sut := &Resource{}

	for _, tt := range validInputs {
		t.Run(tt.in, func(t *testing.T) {
			s := sut.GetResourceName(tt.in)
			assert.Equal(t, s, tt.out)
		})
	}
}

func Test_GetServiceName(t *testing.T) {
	var validInputs = []struct {
		in  string
		out string
	}{
		{"someValue", "SomeValueService"},
		{"SomeValue", "SomeValueService"},
	}

	sut := &Resource{}

	for _, tt := range validInputs {
		t.Run(tt.in, func(t *testing.T) {
			s := sut.GetServiceName(tt.in)
			assert.Equal(t, s, tt.out)
		})
	}
}
