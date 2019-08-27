package codegenerator

import (
	"strings"
)

type Resource struct {
	Methods map[string]*Method `json:"methods"`
}

type Method struct {
	Path           string                             `json:"path"`
	ID             string                             `json:"id"`
	Request        *MethodRequestType                 `json:"request"`
	Description    string                             `json:"description"`
	Response       *MethodResponseType                `json:"response"`
	ParameterOrder []string                           `json:"parameterOrder"`
	HTTPMethod     string                             `json:"httpMethod"`
	Scopes         []string                           `json:"scopes"`
	Parameters     map[string]*MethodRequestParameter `json:"parameters"`
	FlatPath       string                             `json:"flatPath"`
}

type MethodRequestType struct {
	Ref string `json:"$ref"`
}

type MethodResponseType struct {
	Ref string `json:"$ref"`
}

type MethodRequestParameter struct {
	Location    string `json:"location"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
	Pattern     string `json:"pattern"`
}

func (r *Resource) GetCallName(resourceName string, methodName string) string {
	return strings.Title(resourceName) + strings.Title(methodName) + "Call"
}

func (r *Resource) GetName(parameterName string) string {
	return strings.Title(parameterName)
}

func (p *MethodRequestParameter) GetType() (string, error) {
	return getTypeFromJsonType(p.Type)
}

func (r *Method) GetRequestName() string {
	return strings.ToLower(r.Request.Ref)
}

func (r *Method) GetRequestType() string {
	return r.Request.Ref
}

func (r *Method) GetResponseType() string {
	return r.Response.Ref
}

func (r *Method) HasRequest() bool {
	return r.Request != nil
}

func (r *Method) GetDescription() string {
	return re.ReplaceAllString(r.Description, "// $1  ")
}

func (m *Method) IsPaged(d *DiscoveryDocument) bool {
	if _, ok := m.Parameters["pageToken"]; ok {
		return true
	}

	if m.Request != nil {
		for _, schemaDefinition := range d.Schemas {
			if schemaDefinition.ID == m.Request.Ref {
				if _, ok := schemaDefinition.Properties["pageToken"]; ok {
					return true
				}
			}
		}
	}

	return false
}

func (m *Method) IsGet() bool {
	return m.HTTPMethod == "GET"
}

func (m *Method) IsPost() bool {
	return m.HTTPMethod == "POST"
}

func (m *Method) HasParameters() bool {
	return len(m.GetParameters()) > 0
}

func (m *Method) GetParameters() map[string]*MethodRequestParameter {
	parameters := make(map[string]*MethodRequestParameter)

	for name, value := range m.Parameters {
		if name == "pageSize" {
			continue
		}

		if name == "pageToken" {
			continue
		}

		parameters[name] = value
	}

	return parameters
}
