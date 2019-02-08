package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewCustomersPostRequest() customersPostRequest {
	return customersPostRequest{
		client:      c,
		queryParams: c.NewCustomersPostQueryParams(),
		pathParams:  c.NewCustomersPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewCustomersPostRequestBody(),
	}
}

type customersPostRequest struct {
	client      *Client
	queryParams *customersPostQueryParams
	pathParams  *customersPostPathParams
	method      string
	headers     http.Header
	requestBody CustomersPostRequestBody
}

func (c *Client) NewCustomersPostQueryParams() *customersPostQueryParams {
	return &customersPostQueryParams{}
}

type customersPostQueryParams struct{}

func (p customersPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *customersPostRequest) QueryParams() *customersPostQueryParams {
	return r.queryParams
}

func (c *Client) NewCustomersPostPathParams() *customersPostPathParams {
	return &customersPostPathParams{}
}

type customersPostPathParams struct {
}

func (p *customersPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *customersPostRequest) PathParams() *customersPostPathParams {
	return r.pathParams
}

func (r *customersPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *customersPostRequest) Method() string {
	return r.method
}

func (s *Client) NewCustomersPostRequestBody() CustomersPostRequestBody {
	return CustomersPostRequestBody{}
}

type CustomersPostRequestBody struct {
	Name                  string  `json:"name"`
	Code                  string  `json:"code"`
	CustomerCategory      string  `json:"customerCategory"`
	Keywords              string  `json:"keywords"`
	Description           string  `json:"description"`
	KVK                   string  `json:"kvk"`
	KVKCity               string  `json:"kvkcity"`
	VAT                   string  `json:"vat"`
	InvoiceSalutation     string  `json:"invoicesalutation"`
	Account               string  `json:"account"`
	Mobile                string  `json:"mobile"`
	Fax                   string  `json:"fax"`
	Website               string  `json:"website"`
	Phone                 string  `json:"phone"`
	Email                 string  `json:"e-mail"`
	BillingEmail          string  `json:"billingEmail"`
	LedgerAccount         string  `json:"ledgeraccount"`
	LedgerAccountTurnover string  `json:"ledgeraccountturnover"`
	Costcenter            string  `json:"costcenter"`
	Costunit              string  `json:"costunit"`
	VisitingAddress       Address `json:"visitingaddress"`
	BillingAddress        Address `json:"billingaddress"`
	MailingAddress        Address `json:"mailingaddress"`
	ContactPerson         Person  `json:"contactperson"`
	AccountManagers       []struct {
		EmployeeNumber string `json:"employeenumber"`
	} `json:"accountmanager"`
	CustomFields CustomFields `json:"customfields"`
}

func (r *customersPostRequest) RequestBody() *CustomersPostRequestBody {
	return &r.requestBody
}

func (r *customersPostRequest) SetRequestBody(body CustomersPostRequestBody) {
	r.requestBody = body
}

func (r *customersPostRequest) NewResponseBody() *AppResponse {
	return &AppResponse{}
}

// type customersPostResponseBody struct {
// 	AppResponse
// }

func (r *customersPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("customers", r.PathParams())
}

func (r *customersPostRequest) Do() (AppResponse, error) {
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

type Address struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postalcode"`
	Note       string `json:"note"`
}

type Person struct {
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Initials   string `json:"initials"`
	Infix      string `json:"infix"`
	Email      string `json:"e-mail"`
	Phone      string `json:"phone"`
	Gender     string `json:"gender"`
	Title      string `json:"title"`
	Function   string `json:"function"`
	Department string `json:"department"`
}
