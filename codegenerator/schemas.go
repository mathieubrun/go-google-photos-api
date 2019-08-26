package codegenerator

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"text/template"
)

type ErrUnknownJsonType struct {
	Type string
}

type Schema struct {
	ID          string                     `json:"id"`
	Description string                     `json:"description"`
	Type        string                     `json:"type"`
	Properties  map[string]*SchemaProperty `json:"properties"`
}

type SchemaProperty struct {
	Ref              string              `json:"$ref"`
	Type             string              `json:"type"`
	Format           string              `json:"format"`
	Description      string              `json:"description"`
	EnumDescriptions []string            `json:"enumDescriptions,omitempty"`
	Items            *SchemaPropertyItem `json:"items,omitempty"`
}

type SchemaPropertyItem struct {
	Ref    string   `json:"$ref"`
	Type   string   `json:"type"`
	Format string   `json:"format"`
	Enum   []string `json:"enum"`
}

func (e *ErrUnknownJsonType) Error() string {
	return fmt.Sprintf("Unknown type: %v", e.Type)
}

func ProcessSchemas(doc *DiscoveryDocument, wr io.Writer) error {
	tpl := template.Must(template.New("schema").Parse(structsTemplate))
	err := tpl.Execute(wr, doc)
	if err != nil {
		return err
	}

	return nil
}

func (p *Schema) GetName(propertyName string) string {
	return strings.Title(propertyName)
}

func (p *SchemaProperty) GetType() (string, error) {
	if p.Type == "array" {
		if p.Items.Ref != "" {
			return "[]*" + p.Items.Ref, nil
		}

		typeStr, err := getTypeFromJsonType(p.Items.Type)
		if err != nil {
			return "", err
		}

		return "[]" + typeStr, nil
	}

	if p.Ref != "" {
		return "*" + p.Ref, nil
	}

	typeStr, err := getTypeFromJsonType(p.Type)
	if err != nil {
		return "", err
	}

	return typeStr, nil
}

var re = regexp.MustCompile(`(.*\n?)`)

func (p *SchemaProperty) GetDescription() string {
	return re.ReplaceAllString(p.Description, "// $1  ")
}

func (p *Schema) GetDescription() string {
	return re.ReplaceAllString(p.Description, "// $1  ")
}

func (p *DiscoveryDocument) IsResponse(s *Schema) bool {
	for _, resourceDefinition := range p.Resources {
		for _, methodDefinition := range resourceDefinition.Methods {
			if methodDefinition.Response != nil && methodDefinition.Response.Ref == s.ID {
				return true
			}
		}
	}

	return false
}

func getTypeFromJsonType(json string) (string, error) {
	switch json {
	case "boolean":
		return "bool", nil
	case "integer":
		return "int64", nil
	case "number":
		return "float64", nil
	case "string":
		return "string", nil
	case "object":
		return "googleapi.RawMessage", nil
	default:
		return "", &ErrUnknownJsonType{Type: json}
	}
}

const structsTemplate = `
{{ range $schemaName, $schemaDefinition := .Schemas }}
{{ $schemaDefinition.GetDescription }}
type {{ $schemaDefinition.ID }} struct {
	{{range $propertyName, $propertyDefinition := $schemaDefinition.Properties}}{{ $propertyDefinition.GetDescription }}
	{{ $schemaDefinition.GetName $propertyName }} {{ $propertyDefinition.GetType }} ` + "`" + `json:"{{$propertyName}},omitempty" ` + "`" + `
	{{end}}{{ if $.IsResponse $schemaDefinition }}
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse ` + "`" + `json:"-"` + "`" + `{{end}}
	ForceSendFields []string ` + "`" + `json:"-"` + "`" + `
	NullFields []string ` + "`" + `json:"-"` + "`" + `
}

func (s *{{ $schemaDefinition.ID }}) MarshalJSON() ([]byte, error) {
	type NoMethod {{ $schemaDefinition.ID }}
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}
{{ end }}
`
