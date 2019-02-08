package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewProjectsGetRequest() ProjectsGetRequest {
	return ProjectsGetRequest{
		client:      c,
		queryParams: c.NewProjectsGetQueryParams(),
		pathParams:  c.NewProjectsGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewProjectsGetRequestBody(),
	}
}

type ProjectsGetRequest struct {
	client      *Client
	queryParams *ProjectsGetQueryParams
	pathParams  *ProjectsGetPathParams
	method      string
	headers     http.Header
	requestBody ProjectsGetRequestBody
}

func (c *Client) NewProjectsGetQueryParams() *ProjectsGetQueryParams {
	return &ProjectsGetQueryParams{}
}

type ProjectsGetQueryParams struct {
	CurrentPage int    `schema:"currentpage,omitempty"`
	Status      string `schema:"status,omitempty"`
}

func (p ProjectsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ProjectsGetRequest) QueryParams() *ProjectsGetQueryParams {
	return r.queryParams
}

func (c *Client) NewProjectsGetPathParams() *ProjectsGetPathParams {
	return &ProjectsGetPathParams{}
}

type ProjectsGetPathParams struct {
}

func (p *ProjectsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProjectsGetRequest) PathParams() *ProjectsGetPathParams {
	return r.pathParams
}

func (r *ProjectsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProjectsGetRequest) Method() string {
	return r.method
}

func (s *Client) NewProjectsGetRequestBody() ProjectsGetRequestBody {
	return ProjectsGetRequestBody{}
}

type ProjectsGetRequestBody struct {
}

func (r *ProjectsGetRequest) RequestBody() *ProjectsGetRequestBody {
	return &r.requestBody
}

func (r *ProjectsGetRequest) SetRequestBody(body ProjectsGetRequestBody) {
	r.requestBody = body
}

func (r *ProjectsGetRequest) NewResponseBody() *ProjectsGetResponseBody {
	return &ProjectsGetResponseBody{}
}

type ProjectsGetResponseBody struct {
	Metadata Metadata `json:"metadata"`
	Results  Projects `json:"results"`
}

func (r *ProjectsGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("projects", r.PathParams())
}

func (r *ProjectsGetRequest) Do() (ProjectsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
