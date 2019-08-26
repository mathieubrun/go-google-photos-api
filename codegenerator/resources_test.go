package codegenerator

import (
	"gotest.tools/assert"
	"testing"
)

func Test_IsPaged(t *testing.T) {
	doc := &DiscoveryDocument{
		Schemas: map[string]*Schema{
			"paged": &Schema{
				ID: "paged",
				Properties: map[string]*SchemaProperty{
					"pageToken": &SchemaProperty{},
				},
			},
		},
	}
	var validInputs = []struct {
		in  *Method
		out bool
	}{
		{&Method{ID: "1", Parameters: map[string]*MethodRequestParameter{"pageToken": &MethodRequestParameter{}}}, true},
		{&Method{ID: "2", Request: &MethodRequestType{Ref: "paged"}}, true},
		{&Method{ID: "3", Parameters: map[string]*MethodRequestParameter{"test": &MethodRequestParameter{}}}, false},
	}

	for _, tt := range validInputs {
		t.Run(tt.in.ID, func(t *testing.T) {
			s := tt.in.IsPaged(doc)
			assert.Equal(t, s, tt.out)
		})
	}
}
