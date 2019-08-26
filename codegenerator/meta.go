package codegenerator

import (
	"html/template"
	"io"
	"strings"
)

const INTEGER = "integer"
const STRING = "string"
const NUMBER = "number"
const ARRAY = "array"

type DiscoveryDocument struct {
	Version                      string                `json:"version"`
	BaseURL                      string                `json:"baseUrl"`
	ServicePath                  string                `json:"servicePath"`
	Kind                         string                `json:"kind"`
	Description                  string                `json:"description"`
	BasePath                     string                `json:"basePath"`
	ID                           string                `json:"id"`
	DocumentationLink            string                `json:"documentationLink"`
	Revision                     string                `json:"revision"`
	DiscoveryVersion             string                `json:"discoveryVersion"`
	VersionModule                bool                  `json:"version_module"`
	Protocol                     string                `json:"protocol"`
	RootURL                      string                `json:"rootUrl"`
	OwnerDomain                  string                `json:"ownerDomain"`
	Name                         string                `json:"name"`
	BatchPath                    string                `json:"batchPath"`
	FullyEncodeReservedExpansion bool                  `json:"fullyEncodeReservedExpansion"`
	Title                        string                `json:"title"`
	OwnerName                    string                `json:"ownerName"`
	Auth                         map[string]*Auth      `json:"auth"`
	Schemas                      map[string]*Schema    `json:"schemas"`
	Resources                    map[string]*Resource  `json:"resources"`
	Parameters                   map[string]*Parameter `json:"parameters"`
}

type Icons struct {
	X16 string `json:"x16"`
	X32 string `json:"x32"`
}

type Parameter struct {
	Description      string   `json:"description"`
	Default          string   `json:"default"`
	Enum             []string `json:"enum"`
	Type             string   `json:"type"`
	EnumDescriptions []string `json:"enumDescriptions"`
	Location         string   `json:"location"`
}

type Auth struct {
	Scopes map[string]*OAuthScope `json:"scopes"`
}

type OAuthScope struct {
	Description string `json:"description"`
}

func ProcessMeta(doc *DiscoveryDocument, wr io.Writer) error {

	tpl := template.Must(template.New("metaTemplate").Parse(metaTemplate))

	err := tpl.Execute(wr, doc)
	if err != nil {
		return err
	}

	return nil
}

func (s *OAuthScope) GetName(scopeUrl string) string {

	return strings.Replace(strings.Title(strings.Replace(scopeUrl, "https://www.googleapis.com/auth/", "", 1)), ".", "", -1) + "Scope"
}

const metaTemplate = `
package photoslibrary // import "google.golang.org/api/photoslibrary/v1"

import (
	"fmt"
	"errors"
	"golang.org/x/net/context"
	"io"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"net/http"
)

const apiId = "{{ .ID }}"
const apiName = "{{ .Name }}"
const apiVersion = "{{ .Version }}"
const basePath = "{{ .BaseURL }}"
const apiRevision = "{{ .Revision }}"

// OAuth2 scopes used by this API.
const (
	{{ $oauth := index .Auth "oauth2" }}
	{{ range $scopeName, $scopeDefinition := $oauth.Scopes }}
	// {{ $scopeDefinition.Description }}
	{{ $scopeDefinition.GetName $scopeName }} = "{{ $scopeName }}"
	{{ end }}
)
`
