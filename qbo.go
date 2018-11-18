/*
Package qbo implements a simple REST client for the QuickBooks Online API.

The simplest way to get started is:

    you can't!
*/
package qbo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	sandboxBaseURL = "https://sandbox-quickbooks.api.intuit.com/v3"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	Account AccountService
}

type Response struct {
	*http.Response

	Time string
}

type Error struct {
	Message string `json:"Message"`
	Detail  string `json:"Detail"`
	Code    string `json:"code"`
	Element string `json:"element"`
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Fault struct {
		Error []Error `json:"Error"`
		Type  string  `json:"type"`
	} `json:"Fault"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Fault.Type)
}

// NewClient optionally takes in an existing http.Client, and then returns a QuickBooks Online client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(sandboxBaseURL)

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}

	c.Account = &AccountServiceClient{client: c}

	return c
}

// NewRequest prepares a QBO request by encoding it's body and setting up HTTP headers
func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "github.com/clone1018/qbo")
	return req, nil
}

func performRequestWithClient(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	req = req.WithContext(ctx)
	return client.Do(req)
}

// Send actually submits the request to the destination http server, checking the response for an error
func (c *Client) Send(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := performRequestWithClient(ctx, c.client, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	response := (func(r *http.Response) *Response {
		response := Response{Response: r}

		return &response
	})(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return response, err
}

// CheckResponse reads the http status code and unmarshal-ability of the http response, and optionally returns an error
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Fault.Error = append(errorResponse.Fault.Error, Error{
				Message: string(data),
			})
		}
	}

	return errorResponse
}
