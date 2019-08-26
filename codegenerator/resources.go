package codegenerator

import (
	"io"
	"strings"
	"text/template"
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

func ProcessResources(doc *DiscoveryDocument, wr io.Writer) error {
	tpl := template.Must(template.New("resources").Parse(resourcesTemplates))
	err := tpl.Execute(wr, doc)
	if err != nil {
		return err
	}

	return nil
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

const resourcesTemplates = `
{{range $resourceName, $resourceDefinition := .Resources }}
{{range $methodName, $methodDefinition := $resourceDefinition.Methods }}
// method id "photoslibrary.{{ $resourceName }}.{{ $methodName }}":

type {{ $resourceDefinition.GetCallName $resourceName $methodName }} struct {
	s                           *Service
	{{range $parameterName, $parameterDefinition := $methodDefinition.GetParameters }}{{ $parameterName }} {{ $parameterDefinition.GetType }}
	{{end}}{{ if $methodDefinition.HasRequest }}{{ $methodDefinition.GetRequestName }}	*{{ $methodDefinition.GetRequestType }}
	{{ end }}{{ if $methodDefinition.IsGet }}ifNoneMatch_ string
	{{ end }}urlParams_                  gensupport.URLParams
	ctx_                        context.Context
	header_                     http.Header
}

// {{ $methodName }}: {{ $methodDefinition.GetDescription }}
func (r *{{ $resourceDefinition.GetServiceName $resourceName }}) {{ $resourceDefinition.GetResourceName $methodName }}({{ if $methodDefinition.HasParameters  }}{{range $parameterName, $parameterDefinition := $methodDefinition.GetParameters }}{{ $parameterName }} {{ $parameterDefinition.GetType }}{{end}}, {{end}}{{ if $methodDefinition.HasRequest }}{{ $methodDefinition.GetRequestName }}	*{{ $methodDefinition.GetRequestType }}{{ end }}) *{{ $resourceDefinition.GetCallName $resourceName $methodName }} {
	c := &{{ $resourceDefinition.GetCallName $resourceName $methodName }}{s: r.s, urlParams_: make(gensupport.URLParams)}
	{{range $parameterName, $parameterDefinition := $methodDefinition.GetParameters }}c.{{ $parameterName }} = {{ $parameterName }}
	{{end}}{{ if $methodDefinition.HasRequest }}c.{{ $methodDefinition.GetRequestName }} = {{ $methodDefinition.GetRequestName }}{{ end }}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) Fields(s ...googleapi.Field) *{{ $resourceDefinition.GetCallName $resourceName $methodName }} {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) Context(ctx context.Context) *{{ $resourceDefinition.GetCallName $resourceName $methodName }} {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

{{ if $methodDefinition.IsPaged $ }}
// PageSize sets the optional parameter "pageSize": Maximum number of
// albums to return in the response. The default number of
// albums to return at a time is 20. The maximum page size is 50.
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) PageSize(pageSize int64) *{{ $resourceDefinition.GetCallName $resourceName $methodName }} {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to get the next page of the results. Adding this to
// the request will return the rows after the pageToken. The pageToken
// should
// be the value returned in the nextPageToken parameter in the response
// to the
// listSharedAlbums request.
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) PageToken(pageToken string) *{{ $resourceDefinition.GetCallName $resourceName $methodName }} {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) Pages(ctx context.Context, f func(*{{ $methodDefinition.GetResponseType }}) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
{{end}}

{{ if $methodDefinition.IsGet }}
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{{ $methodDefinition.Path }}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	{{ if $methodDefinition.HasParameters  }}googleapi.Expand(req.URL, map[string]string{
		{{ range $parameterName, $parameterDefinition := $methodDefinition.GetParameters }}"{{ $parameterName }}": fmt.Sprintf("%v", c.{{ $parameterName }}),
		{{ end}}
	}){{end}}
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}
{{end}}

{{ if $methodDefinition.IsPost }}
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.{{ $methodDefinition.GetRequestName }})
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{{ $methodDefinition.Path }}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	{{ if $methodDefinition.HasParameters  }}googleapi.Expand(req.URL, map[string]string{
		{{ range $parameterName, $parameterDefinition := $methodDefinition.GetParameters }}"{{ $parameterName }}": fmt.Sprintf("%v", c.{{ $parameterName }}),
		{{ end}}
	}){{ end}}
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}
{{end}}

// Do executes the "photoslibrary.{{ $resourceName }}.{{ $methodName }}" call.
// Exactly one of *{{ $methodDefinition.GetResponseType }} or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *{{ $methodDefinition.GetResponseType }}.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *{{ $resourceDefinition.GetCallName $resourceName $methodName }}) Do(opts ...googleapi.CallOption) (*{{ $methodDefinition.GetResponseType }}, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &{{ $methodDefinition.GetResponseType }}{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
}

{{end}}
{{end}}
`
