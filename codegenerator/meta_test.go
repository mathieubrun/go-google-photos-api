package codegenerator

import (
	"gotest.tools/assert"
	"testing"
)

func Test_GetScopeName(t *testing.T) {
	var validInputs = []struct {
		in  string
		out string
	}{
		{"https://www.googleapis.com/auth/photoslibrary", "PhotoslibraryScope"},
		{"https://www.googleapis.com/auth/photoslibrary.readonly", "PhotoslibraryReadonlyScope"},
	}

	sut := &OAuthScope{}

	for _, tt := range validInputs {
		t.Run(tt.in, func(t *testing.T) {
			s := sut.GetName(tt.in)
			assert.Equal(t, s, tt.out)
		})
	}
}
