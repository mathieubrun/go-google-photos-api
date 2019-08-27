package codegenerator_test

import (
	"encoding/json"
	"github.com/mathieubrun/go-google-photos-api/generator/codegenerator"
	"gotest.tools/assert"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_Process(t *testing.T) {
	fd, err := os.Open("../photoslibrary/v1/photoslibrary-api.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	doc := &codegenerator.DiscoveryDocument{}
	json.NewDecoder(fd).Decode(doc)

	var validInputs = []struct {
		tpl      string
		function func(doc *codegenerator.DiscoveryDocument, wr io.Writer, tpl string) error
	}{
		{"../templates/meta.tpl", codegenerator.Process},
		{"../templates/resources.tpl", codegenerator.Process},
		{"../templates/services.tpl", codegenerator.Process},
		{"../templates/schemas.tpl", codegenerator.Process},
	}

	for _, tt := range validInputs {
		t.Run(tt.tpl, func(t *testing.T) {
			var b strings.Builder
			err = tt.function(doc, &b, tt.tpl)

			assert.NilError(t, err)
			assert.Assert(t, b.String() != "")
		})
	}
}
