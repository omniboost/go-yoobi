package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewEmployeesGetRequest() EmployeesGetRequest {
	return EmployeesGetRequest{
		client:      c,
		queryParams: c.NewEmployeesGetQueryParams(),
		pathParams:  c.NewEmployeesGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewEmployeesGetRequestBody(),
	}
}

type EmployeesGetRequest struct {
	client      *Client
	queryParams *EmployeesGetQueryParams
	pathParams  *EmployeesGetPathParams
	method      string
	headers     http.Header
	requestBody EmployeesGetRequestBody
}

func (c *Client) NewEmployeesGetQueryParams() *EmployeesGetQueryParams {
	return &EmployeesGetQueryParams{}
}

type EmployeesGetQueryParams struct {
	CurrentPage int    `schema:"currentpage"`
	Status      string `schema:"status,omitempty"`
}

func (p EmployeesGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *EmployeesGetRequest) QueryParams() *EmployeesGetQueryParams {
	return r.queryParams
}

func (c *Client) NewEmployeesGetPathParams() *EmployeesGetPathParams {
	return &EmployeesGetPathParams{}
}

type EmployeesGetPathParams struct {
}

func (p *EmployeesGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeesGetRequest) PathParams() *EmployeesGetPathParams {
	return r.pathParams
}

func (r *EmployeesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesGetRequest) Method() string {
	return r.method
}

func (s *Client) NewEmployeesGetRequestBody() EmployeesGetRequestBody {
	return EmployeesGetRequestBody{}
}

type EmployeesGetRequestBody struct {
}

func (r *EmployeesGetRequest) RequestBody() *EmployeesGetRequestBody {
	return &r.requestBody
}

func (r *EmployeesGetRequest) SetRequestBody(body EmployeesGetRequestBody) {
	r.requestBody = body
}

func (r *EmployeesGetRequest) NewResponseBody() *EmployeesGetResponseBody {
	return &EmployeesGetResponseBody{}
}

type EmployeesGetResponseBody struct {
	Metadata Metadata  `json:"metadata"`
	Results  Employees `json:"results"`
}

func (r *EmployeesGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("employees", r.PathParams())
}

func (r *EmployeesGetRequest) Do() (EmployeesGetResponseBody, error) {
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
