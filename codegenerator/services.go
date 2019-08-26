package codegenerator

import (
	"html/template"
	"io"
	"strings"
)

func ProcessServices(doc *DiscoveryDocument, wr io.Writer) error {

	tpl := template.Must(template.New("servicesTemplate").Parse(servicesTemplate))

	err := tpl.Execute(wr, doc)
	if err != nil {
		return err
	}

	return nil
}

func (r *Resource) GetResourceName(resourceName string) string {
	return strings.Title(resourceName)
}

func (r *Resource) GetServiceName(resourceName string) string {
	return r.GetResourceName(resourceName) + "Service"
}

const servicesTemplate = `
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	{{ range $resourceName, $resourceDefinition := .Resources }}
	s.{{ $resourceDefinition.GetResourceName $resourceName }} = New{{ $resourceDefinition.GetServiceName $resourceName }}(s)
	{{ end }}
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Albums *AlbumsService

	MediaItems *MediaItemsService

	SharedAlbums *SharedAlbumsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

{{ range $resourceName, $resourceDefinition := .Resources }}
func New{{ $resourceDefinition.GetServiceName $resourceName }}(s *Service) *{{ $resourceDefinition.GetServiceName $resourceName }} {
	rs := &{{ $resourceDefinition.GetServiceName $resourceName }}{s: s}
	return rs
}

type {{ $resourceDefinition.GetServiceName $resourceName }} struct {
	s *Service
}
{{ end }}
`
