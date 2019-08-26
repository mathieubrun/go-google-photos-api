package codegenerator_test

import (
	"encoding/json"
	"github.com/mathieubrun/go-google-photos-api/codegenerator"
	"gotest.tools/assert"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_Process(t *testing.T) {
	fd, err := os.Open("../discoveryDocumentV1.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	doc := &codegenerator.DiscoveryDocument{}
	json.NewDecoder(fd).Decode(doc)

	var validInputs = []struct {
		name     string
		function func(doc *codegenerator.DiscoveryDocument, wr io.Writer) error
	}{
		{"codegenerator.ProcessMeta", codegenerator.ProcessMeta},
		{"codegenerator.ProcessServices", codegenerator.ProcessServices},
		{"codegenerator.ProcessSchemas", codegenerator.ProcessSchemas},
		{"codegenerator.ProcessResources", codegenerator.ProcessResources},
	}

	for _, tt := range validInputs {
		t.Run(tt.name, func(t *testing.T) {
			var b strings.Builder
			err = tt.function(doc, &b)

			assert.NilError(t, err)
			assert.Assert(t, b.String() != "")
		})
	}
}
