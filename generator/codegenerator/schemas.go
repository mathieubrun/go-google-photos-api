package codegenerator

import (
	"fmt"
	"regexp"
	"strings"
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
	PropertyDefinition
	Description      string              `json:"description"`
	EnumDescriptions []string            `json:"enumDescriptions,omitempty"`
	Items            *SchemaPropertyItem `json:"items,omitempty"`
}

type SchemaPropertyItem struct {
	PropertyDefinition
	Enum []string `json:"enum"`
}

type PropertyDefinition struct {
	Type   string `json:"type"`
	Format string `json:"format"`
	Ref    string `json:"$ref"`
}

func (e *ErrUnknownJsonType) Error() string {
	return fmt.Sprintf("Unknown type: %v", e.Type)
}

func (p *Schema) GetName(propertyName string) string {
	return strings.Title(propertyName)
}

func (p *SchemaProperty) GetJsonTag() string {
	if t, _ := p.GetType(); p.Format != "" && t != "string" {
		return p.Type
	}

	return ""
}

func (p *SchemaProperty) GetType() (string, error) {
	if p.Type == "array" {
		if p.Items.Ref != "" {
			return "[]*" + p.Items.Ref, nil
		}

		typeStr, err := getTypeFromJson(p.Items.PropertyDefinition)
		if err != nil {
			return "", err
		}

		return "[]" + typeStr, nil
	}

	if p.Ref != "" {
		return "*" + p.Ref, nil
	}

	typeStr, err := getTypeFromJson(p.PropertyDefinition)
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

func getTypeFromJson(t PropertyDefinition) (string, error) {
	if t.Format != "" {
		return getTypeFromFormat(t.Format)
	}

	return getTypeFromJsonType(t.Type)
}

func getTypeFromFormat(format string) (string, error) {
	switch format {
	case "int32":
		return "int", nil
	case "int64":
		return "int64", nil
	case "float":
		return "float32", nil
	case "double":
		return "float32", nil
	case "google-datetime":
		return "string", nil
	case "google-duration":
		return "string", nil
	default:
		return "", &ErrUnknownJsonType{Type: format}
	}
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
