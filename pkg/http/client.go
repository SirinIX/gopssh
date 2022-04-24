package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

const (
	userAgent      = "Dtstack DevOps 1.0"
	
	defaultBaseURL = ""
	mediaType      = "application/json"
)

type Client struct {
	client             *http.Client
	BaseURL            *url.URL
	UserAgent          string
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

type Response struct {
	*http.Response
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
	// Error message
	Message string
}

func NewInsecureHTTPClient() (*Client, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	return NewHTTPClient(client)
}

func NewHTTPClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		return nil, errors.New("http client must not be nil")
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, params map[string]string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if len(params) != 0 {
		for param, value := range params {
			q := rel.Query()
			q.Set(param, value)
			rel.RawQuery = q.Encode()
		}

	}
	rel.Path = path.Join(c.BaseURL.Path, rel.Path)
	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func newResponse(r *http.Response) *Response {
	response := Response{Response: r}
	return &response
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; c >= 200 && c < 300 {
		return nil
	}
	return fmt.Errorf("http response code: %v", r.StatusCode)
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, resp)
	}

	defer resp.Body.Close()

	response := newResponse(resp)
	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err := io.Copy(w, resp.Body)
			if err != nil {
				return response, err
			}
		} else {
			err := json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return response, err
			}
		}
	}

	return response, err
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
