package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewEmployeeContractsPostRequest() EmployeeContractsPostRequest {
	return EmployeeContractsPostRequest{
		client:      c,
		queryParams: c.NewEmployeeContractsPostQueryParams(),
		pathParams:  c.NewEmployeeContractsPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewEmployeeContractsPostRequestBody(),
	}
}

type EmployeeContractsPostRequest struct {
	client      *Client
	queryParams *EmployeeContractsPostQueryParams
	pathParams  *EmployeeContractsPostPathParams
	method      string
	headers     http.Header
	requestBody EmployeeContractsPostRequestBody
}

func (c *Client) NewEmployeeContractsPostQueryParams() *EmployeeContractsPostQueryParams {
	return &EmployeeContractsPostQueryParams{}
}

type EmployeeContractsPostQueryParams struct{}

func (p EmployeeContractsPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *EmployeeContractsPostRequest) QueryParams() *EmployeeContractsPostQueryParams {
	return r.queryParams
}

func (c *Client) NewEmployeeContractsPostPathParams() *EmployeeContractsPostPathParams {
	return &EmployeeContractsPostPathParams{}
}

type EmployeeContractsPostPathParams struct {
}

func (p *EmployeeContractsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeeContractsPostRequest) PathParams() *EmployeeContractsPostPathParams {
	return r.pathParams
}

func (r *EmployeeContractsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeeContractsPostRequest) Method() string {
	return r.method
}

func (s *Client) NewEmployeeContractsPostRequestBody() EmployeeContractsPostRequestBody {
	return EmployeeContractsPostRequestBody{}
}

type EmployeeContractsPostRequestBody struct {
	PersNum    string `json:"pers_num"`
	ContractNr string `json:"contractnr"`
	StartDatum *Date  `json:"startdatum"`
	EindDatum  *Date  `json:"einddatum"`
	Type       string `json:"type"`
	FTE        string `json:"fte"`
	Notitie    string `json:"notitie"`
	WeekAantal string `json:"week_aantal"`
	Maandag1   string `json:"maandag1"`
	Dinsdag1   string `json:"dinsdag1"`
	Woensdag1  string `json:"woensdag1"`
	Donderdag1 string `json:"donderdag1"`
	Vrijdag1   string `json:"vrijdag1"`
	Zaterdag1  string `json:"zaterdag1"`
	Zondag1    string `json:"zondag1"`
}

func (r *EmployeeContractsPostRequest) RequestBody() *EmployeeContractsPostRequestBody {
	return &r.requestBody
}

func (r *EmployeeContractsPostRequest) SetRequestBody(body EmployeeContractsPostRequestBody) {
	r.requestBody = body
}

func (r *EmployeeContractsPostRequest) NewResponseBody() *AppResponse {
	return &AppResponse{}
}

// type EmployeeContractsPostResponseBody struct {
// 	AppResponse
// }

func (r *EmployeeContractsPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("employeeContract", r.PathParams())
}

func (r *EmployeeContractsPostRequest) Do() (AppResponse, error) {
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
