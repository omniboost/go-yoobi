package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewEmployeeContractsGetRequest() EmployeeContractsGetRequest {
	return EmployeeContractsGetRequest{
		client:      c,
		queryParams: c.NewEmployeeContractsGetQueryParams(),
		pathParams:  c.NewEmployeeContractsGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewEmployeeContractsGetRequestBody(),
	}
}

type EmployeeContractsGetRequest struct {
	client      *Client
	queryParams *EmployeeContractsGetQueryParams
	pathParams  *EmployeeContractsGetPathParams
	method      string
	headers     http.Header
	requestBody EmployeeContractsGetRequestBody
}

func (c *Client) NewEmployeeContractsGetQueryParams() *EmployeeContractsGetQueryParams {
	return &EmployeeContractsGetQueryParams{}
}

type EmployeeContractsGetQueryParams struct {
	CurrentPage int    `schema:"currentpage,omitempty"`
	Status      string `schema:"status,omitempty"`
}

func (p EmployeeContractsGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *EmployeeContractsGetRequest) QueryParams() *EmployeeContractsGetQueryParams {
	return r.queryParams
}

func (c *Client) NewEmployeeContractsGetPathParams() *EmployeeContractsGetPathParams {
	return &EmployeeContractsGetPathParams{}
}

type EmployeeContractsGetPathParams struct {
}

func (p *EmployeeContractsGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeeContractsGetRequest) PathParams() *EmployeeContractsGetPathParams {
	return r.pathParams
}

func (r *EmployeeContractsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeeContractsGetRequest) Method() string {
	return r.method
}

func (s *Client) NewEmployeeContractsGetRequestBody() EmployeeContractsGetRequestBody {
	return EmployeeContractsGetRequestBody{}
}

type EmployeeContractsGetRequestBody struct {
}

func (r *EmployeeContractsGetRequest) RequestBody() *EmployeeContractsGetRequestBody {
	return &r.requestBody
}

func (r *EmployeeContractsGetRequest) SetRequestBody(body EmployeeContractsGetRequestBody) {
	r.requestBody = body
}

func (r *EmployeeContractsGetRequest) NewResponseBody() *EmployeeContractsGetResponseBody {
	return &EmployeeContractsGetResponseBody{}
}

type EmployeeContractsGetResponseBody struct {
	Metadata Metadata          `json:"metadata"`
	Results  EmployeeContracts `json:"results"`
}

func (r *EmployeeContractsGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("employees", r.PathParams())
}

func (r *EmployeeContractsGetRequest) Do() (EmployeeContractsGetResponseBody, error) {
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
