package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewEmployeesPostRequest() EmployeesPostRequest {
	return EmployeesPostRequest{
		client:      c,
		queryParams: c.NewEmployeesPostQueryParams(),
		pathParams:  c.NewEmployeesPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewEmployeesPostRequestBody(),
	}
}

type EmployeesPostRequest struct {
	client      *Client
	queryParams *EmployeesPostQueryParams
	pathParams  *EmployeesPostPathParams
	method      string
	headers     http.Header
	requestBody EmployeesPostRequestBody
}

func (c *Client) NewEmployeesPostQueryParams() *EmployeesPostQueryParams {
	return &EmployeesPostQueryParams{}
}

type EmployeesPostQueryParams struct{}

func (p EmployeesPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *EmployeesPostRequest) QueryParams() *EmployeesPostQueryParams {
	return r.queryParams
}

func (c *Client) NewEmployeesPostPathParams() *EmployeesPostPathParams {
	return &EmployeesPostPathParams{}
}

type EmployeesPostPathParams struct {
}

func (p *EmployeesPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeesPostRequest) PathParams() *EmployeesPostPathParams {
	return r.pathParams
}

func (r *EmployeesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesPostRequest) Method() string {
	return r.method
}

func (s *Client) NewEmployeesPostRequestBody() EmployeesPostRequestBody {
	return EmployeesPostRequestBody{}
}

type EmployeesPostRequestBody struct {
	EmployeeNumber        string       `json:"employeenumber"`
	BSN                   string       `json:"bsn"`
	DateOfBirth           *Date        `json:"dateofbirth"`
	FirstName             string       `json:"firstname"`
	Initials              string       `json:"initials"`
	Infix                 string       `json:"infix"`
	LastName              string       `json:"lastname"`
	Gender                string       `json:"gender"`
	EmployeeAbbr          string       `json:"employeeabbr"`
	JobName               string       `json:"jobname"`
	StartRegistrationDate *Date        `json:"startregistrationdate"`
	StartDate             Date         `json:"startdate"`
	EndDate               *Date        `json:"enddate"`
	CostCenter            string       `json:"costcenter"`
	CostUnit              string       `json:"costunit"`
	AllowOvertime         Bool         `json:"allowovertime"`
	Department            Department   `json:"department"`
	State                 string       `json:"state"`
	User                  *User        `json:"user,omitempty"`
	CustomFields          CustomFields `json:"customfields"`
}

func (r *EmployeesPostRequest) RequestBody() *EmployeesPostRequestBody {
	return &r.requestBody
}

func (r *EmployeesPostRequest) SetRequestBody(body EmployeesPostRequestBody) {
	r.requestBody = body
}

func (r *EmployeesPostRequest) NewResponseBody() *EmployeesPostResponseBody {
	return &EmployeesPostResponseBody{}
}

type EmployeesPostResponseBody struct {
	AppResponse
}

func (r *EmployeesPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("employees", r.PathParams())
}

func (r *EmployeesPostRequest) Do() (EmployeesPostResponseBody, error) {
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
