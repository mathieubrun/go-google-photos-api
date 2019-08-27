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
