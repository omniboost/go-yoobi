package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewCustomersGetRequest() CustomersGetRequest {
	return CustomersGetRequest{
		client:      c,
		queryParams: c.NewCustomersGetQueryParams(),
		pathParams:  c.NewCustomersGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCustomersGetRequestBody(),
	}
}

type CustomersGetRequest struct {
	client      *Client
	queryParams *CustomersGetQueryParams
	pathParams  *CustomersGetPathParams
	method      string
	headers     http.Header
	requestBody CustomersGetRequestBody
}

func (c *Client) NewCustomersGetQueryParams() *CustomersGetQueryParams {
	return &CustomersGetQueryParams{}
}

type CustomersGetQueryParams struct {
	CurrentPage int    `schema:"currentpage"`
	Status      string `schema:"status,omitempty"`
}

func (p CustomersGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomersGetRequest) QueryParams() *CustomersGetQueryParams {
	return r.queryParams
}

func (c *Client) NewCustomersGetPathParams() *CustomersGetPathParams {
	return &CustomersGetPathParams{}
}

type CustomersGetPathParams struct {
}

func (p *CustomersGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CustomersGetRequest) PathParams() *CustomersGetPathParams {
	return r.pathParams
}

func (r *CustomersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomersGetRequest) Method() string {
	return r.method
}

func (s *Client) NewCustomersGetRequestBody() CustomersGetRequestBody {
	return CustomersGetRequestBody{}
}

type CustomersGetRequestBody struct {
}

func (r *CustomersGetRequest) RequestBody() *CustomersGetRequestBody {
	return &r.requestBody
}

func (r *CustomersGetRequest) SetRequestBody(body CustomersGetRequestBody) {
	r.requestBody = body
}

func (r *CustomersGetRequest) NewResponseBody() *CustomersGetResponseBody {
	return &CustomersGetResponseBody{}
}

type CustomersGetResponseBody struct {
	Metadata Metadata  `json:"metadata"`
	Results  Customers `json:"results"`
}

func (r *CustomersGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("customers", r.PathParams())
}

func (r *CustomersGetRequest) Do() (CustomersGetResponseBody, error) {
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
