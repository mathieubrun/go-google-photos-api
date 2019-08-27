{{ range $schemaName, $schemaDefinition := .Schemas }}
{{ $schemaDefinition.GetDescription }}
type {{ $schemaDefinition.ID }} struct {
	{{range $propertyName, $propertyDefinition := $schemaDefinition.Properties}}{{ $propertyDefinition.GetDescription }}
	{{ $schemaDefinition.GetName $propertyName }} {{ $propertyDefinition.GetType }} `json:"{{ $propertyName }},omitempty,{{ $propertyDefinition.GetJsonTag }}"`
	{{end}}{{ if $.IsResponse $schemaDefinition }}
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`{{end}}
	ForceSendFields []string `json:"-"`
	NullFields []string `json:"-"`
}

func (s *{{ $schemaDefinition.ID }}) MarshalJSON() ([]byte, error) {
	type NoMethod {{ $schemaDefinition.ID }}
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}
{{ end }}
