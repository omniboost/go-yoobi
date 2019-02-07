package yoobi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"text/template"
	"time"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-yoobi/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
	apiVersion     = "v1"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "{{.organisation}}.yoobi.nl",
		Path:   "api/{{.version}}/",
	}
)

// NewClient returns a new InvoiceXpress Client client
func NewClient(httpClient *http.Client, organisation string, username string, password string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http:              httpClient,
		requestTimestamps: []time.Time{},
	}

	client.SetOrganisation(organisation)
	client.SetUsername(username)
	client.SetPassword(password)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with InvoiceXpress Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	organisation string
	username     string
	password     string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback

	requestTimestamps []time.Time
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) Organisation() string {
	return c.organisation
}

func (c *Client) SetOrganisation(organisation string) {
	c.organisation = organisation
}

func (c *Client) Username() string {
	return c.username
}

func (c *Client) SetUsername(username string) {
	c.username = username
}

func (c *Client) Password() string {
	return c.password
}

func (c *Client) SetPassword(password string) {
	c.password = password
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) GetEndpointURL(path string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	clientURL.Host = strings.Replace(clientURL.Host, "{{.organisation}}", c.Organisation(), -1)
	clientURL.Path = clientURL.Path + path

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	params["version"] = apiVersion
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, method string, URL url.URL, body interface{}) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	req, err := http.NewRequest(method, URL.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set credentials
	req.SetBasicAuth(c.Username(), c.Password())

	// optionally pass along context
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	// set other headers
	req.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	req.Header.Add("Accept", c.MediaType())
	req.Header.Add("User-Agent", c.UserAgent())

	return req, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	c.SleepUntilRequestRate()
	c.RegisterRequestTimestamp(time.Now())

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	errorResponse := &ErrorResponse{Response: httpResp}

	err = c.Unmarshal(httpResp.Body, &responseBody, &errorResponse)
	if err != nil {
		return httpResp, err
	}

	appResponse, ok := responseBody.(*AppResponse)
	if ok {
		if len(appResponse.Errors) > 0 {
			return httpResp, appResponse.Errors
		}
	}

	if errorResponse.Err.Message != "" {
		return httpResp, errorResponse
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	wg := sync.WaitGroup{}
	wg.Add(len(vv))
	errs := []error{}
	writers := make([]io.Writer, len(vv))

	for i, v := range vv {
		pr, pw := io.Pipe()
		writers[i] = pw

		go func(i int, v interface{}, pr *io.PipeReader, pw *io.PipeWriter) {
			dec := json.NewDecoder(pr)
			if c.disallowUnknownFields {
				dec.DisallowUnknownFields()
			}

			err := dec.Decode(v)
			if err != nil {
				errs = append(errs, err)
			}

			// mark routine as done
			wg.Done()

			// Drain reader
			io.Copy(ioutil.Discard, pr)

			// close reader
			// pr.CloseWithError(err)
			pr.Close()
		}(i, v, pr, pw)
	}

	// copy the data in a multiwriter
	mw := io.MultiWriter(writers...)
	_, err := io.Copy(mw, r)
	if err != nil {
		return err
	}

	wg.Wait()
	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}
	return nil
}

func (c *Client) RegisterRequestTimestamp(t time.Time) {
	if len(c.requestTimestamps) >= 5 {
		c.requestTimestamps = c.requestTimestamps[1:5]
	}
	c.requestTimestamps = append(c.requestTimestamps, t)
}

func (c *Client) SleepUntilRequestRate() {
	// Requestrate is 5r/1s

	// if there are less then 5 registered requests: execute the request
	// immediately
	if len(c.requestTimestamps) < 4 {
		return
	}

	// is the first item within 1 second? If it's > 1 second the request can be
	// executed imediately
	diff := time.Now().Sub(c.requestTimestamps[0])
	if diff >= time.Second {
		return
	}

	// Sleep for the time it takes for the first item to be > 1 second old
	time.Sleep(time.Second - diff)
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return errorResponse
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Err struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Type    string `json:"type"`
	} `json:"error"`
}

type AppResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Errors     Errors            `json:"errors"`
	ResultID   string            `json:"resultid"`
	ReturnVars interface{}       `json:"returnvars"`
	Object     string            `json:"object"`
	Warnings   map[string]string `json:"warnings"`
}

type Errors map[string]string

func (errs Errors) Error() string {
	err := []string{}
	for k, v := range errs {
		err = append(err, fmt.Sprintf("%s: %s", k, v))
	}

	return strings.Join(err, ", ")
}

func (r *AppResponse) UnmarshalJSON(data []byte) error {
	type Alias AppResponse
	tmp := []Alias{}

	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	if len(tmp) == 1 {
		*r = AppResponse(tmp[0])
		return nil
	}

	if len(tmp) > 1 {
		return errors.New("Too many elements in AppResponse")
	}
	return errors.New("Too little elements in AppResponse")
}

// {"error":{"message":"key [currency] doesn't exist","code":"0","type":"expression"}}
// [{"errors":{},"resultid":"63759415-E3B7-408D-A648-A4F2036E51E3","returnvars":{},"object":"projectactiviteit","warnings":{}}]

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%s: %s (%s)", r.Err.Code, r.Err.Message, r.Err.Type)
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

type PathParams interface {
	Params() map[string]string
}

type FilterParams map[string]string

// type FilterParams []Filter

// type Filter struct {
// 	Field    string
// 	Operator string
// 	Value    string
// }

type Metadata struct {
	Last        int `json:"last"`
	Results     Int `json:"results"`
	Currentpage Int `json:"currentpage"`
}
