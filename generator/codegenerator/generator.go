package codegenerator

import (
	"encoding/json"
	"html/template"
	"io"
)

func Process(doc *DiscoveryDocument, wr io.Writer, templatePath string) error {
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	err = tpl.Execute(wr, doc)
	if err != nil {
		return err
	}

	return nil
}

func ProcessAll(document io.Reader, output io.Writer) (err error) {
	doc := &DiscoveryDocument{}
	json.NewDecoder(document).Decode(doc)

	err = Process(doc, output, "templates/meta.tpl")
	if err != nil {
		return err
	}

	err = Process(doc, output, "templates/services.tpl")
	if err != nil {
		return err
	}

	err = Process(doc, output, "templates/schemas.tpl")
	if err != nil {
		return err
	}

	err = Process(doc, output, "templates/resources.tpl")
	if err != nil {
		return err
	}

	return nil
}
