package codegenerator

import (
	"encoding/json"
	"io"
)

func ProcessAll(document io.Reader, output io.Writer) (err error) {
	doc := &DiscoveryDocument{}
	json.NewDecoder(document).Decode(doc)

	err = ProcessMeta(doc, output)
	if err != nil {
		return err
	}

	err = ProcessServices(doc, output)
	if err != nil {
		return err
	}

	err = ProcessSchemas(doc, output)
	if err != nil {
		return err
	}

	err = ProcessResources(doc, output)
	if err != nil {
		return err
	}

	return nil
}
