package yoobi

import (
	"net/http"
	"net/url"
)

func (c *Client) NewProjectsPostRequest() ProjectsPostRequest {
	return ProjectsPostRequest{
		client:      c,
		queryParams: c.NewProjectsPostQueryParams(),
		pathParams:  c.NewProjectsPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewProjectsPostRequestBody(),
	}
}

type ProjectsPostRequest struct {
	client      *Client
	queryParams *ProjectsPostQueryParams
	pathParams  *ProjectsPostPathParams
	method      string
	headers     http.Header
	requestBody ProjectsPostRequestBody
}

func (c *Client) NewProjectsPostQueryParams() *ProjectsPostQueryParams {
	return &ProjectsPostQueryParams{}
}

type ProjectsPostQueryParams struct{}

func (p ProjectsPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *ProjectsPostRequest) QueryParams() *ProjectsPostQueryParams {
	return r.queryParams
}

func (c *Client) NewProjectsPostPathParams() *ProjectsPostPathParams {
	return &ProjectsPostPathParams{}
}

type ProjectsPostPathParams struct {
}

func (p *ProjectsPostPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ProjectsPostRequest) PathParams() *ProjectsPostPathParams {
	return r.pathParams
}

func (r *ProjectsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *ProjectsPostRequest) Method() string {
	return r.method
}

func (s *Client) NewProjectsPostRequestBody() ProjectsPostRequestBody {
	return ProjectsPostRequestBody{}
}

type ProjectsPostRequestBody struct {
	Name           string `json:"name"`
	Code           string `json:"code"`
	State          string `json:"state"`
	StartDate      Date   `json:"startdate"`
	EndDate        Date   `json:"enddate"`
	Label          string `json:"label"`
	Classification string `json:"classification"`
	Customer       struct {
		Code string `json:"code"`
	}
	Department struct {
		Name string `json:"name"`
	}
	LedgerAccount string       `json:"ledgeraccount"`
	CostCenter    string       `json:"costcenter"`
	CostUnit      string       `json:"costunit"`
	Currency      string       `json:"currency"`
	ProjectRoles  ProjectRoles `json:"projectroles"`
	Activities    Activities   `json:"activities"`
	CustomFields  CustomFields `json:"customfields"`
}

type ProjectRoles []ProjectRole

type ProjectRole struct {
	LastName        string `json:"lastname"`
	ProjectRoleName string `json:"projectrolename"`
	EmployeeNumber  int    `json:"employeenumber"`
	FirstName       string `json:"firstname"`
	Infix           string `json:"infix"`
}

type Activities []Activity

type Activity struct {
	Name            string `json:"name"`
	StartDate       Date   `json:"startdate"`
	Description     string `json:"description"`
	EndDate         Date   `json:"enddate"`
	State           string `json:"state"`
	ExcludeApproval Bool   `json:"exclude_approval"`
	Type            string `json:"type"`
	IsBillable      Bool   `json:"isbillable"`
	Budget          struct {
		BudgetTime   int `json:"budgettime"`
		BudgetNumber int `json:"budgetnumber"`
		BudgetCost   int `json:"budgetcost"`
		TimeActive   int `json:"timeactive"`
		NumberActive int `json:"numberactive"`
		CostActive   int `json:"CostActive"`
	}
	LedgerAccount string `json:"ledgeraccount"`
	CostCenter    string `json:"costcenter"`
	CostUnit      string `json:"costunit"`
}

func (r *ProjectsPostRequest) RequestBody() *ProjectsPostRequestBody {
	return &r.requestBody
}

func (r *ProjectsPostRequest) SetRequestBody(body ProjectsPostRequestBody) {
	r.requestBody = body
}

func (r *ProjectsPostRequest) NewResponseBody() *ProjectsPostResponseBody {
	return &ProjectsPostResponseBody{}
}

type ProjectsPostResponseBody struct {
	Metadata Metadata `json:"metadata"`
	Results  Projects `json:"results"`
}

func (r *ProjectsPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("projects", r.PathParams())
}

func (r *ProjectsPostRequest) Do() (ProjectsPostResponseBody, error) {
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
