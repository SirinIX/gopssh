package http

import (
	"gopssh/log"
	"fmt"
	"net/http"
)

type HttpRequest interface{}

type HttpResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data HttpResponseData `json:"data"`
}

type HttpResponseData interface{}

func (c *Client) GetRequest(uri string, headers, cookies map[string]string) (HttpResponseData, error) {
	// Build request
	req, err := c.NewRequest("GET", uri, nil, nil)
	if err != nil {
		log.Error("failed to build GET request %v, error: %v", uri, err)
		return nil, err
	}

	// Set headers
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	for key, val := range cookies {
		req.AddCookie(&http.Cookie{
			Name:  key,
			Value: val,
		})
	}

	log.Info("do get request, uri: %v", uri)
	// Do request
	resBody := new(HttpResponse)
	resp, err := c.Do(req, resBody)
	if err != nil {
		log.Error("failed to do get request %v, error: %v", uri, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resBody.Code != 0 {
		log.Error("failed to do get request %v, code: %v, msg: %v\n data: %v", uri, resBody.Code, resBody.Msg, resBody.Data)
		return nil, fmt.Errorf("response code: %v", resBody.Code)
	}
	log.Info("successfully do get request %v, message: %v, code: %v", uri, resBody.Msg, resBody.Code)

	return resBody.Data, nil
}

func (c *Client) PostRequest(uri string, request HttpRequest, headers, cookies map[string]string) (HttpResponseData, error) {
	// Build request
	req, err := c.NewRequest("POST", uri, nil, request)
	if err != nil {
		log.Error("failed to build POST request", err)
		return nil, err
	}

	// Set headers
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	for key, val := range cookies {
		req.AddCookie(&http.Cookie{
			Name:  key,
			Value: val,
		})
	}

	log.Info("do post request, uri: %v, body: %v", uri, request)
	// Do request
	resBody := new(HttpResponse)
	resp, err := c.Do(req, resBody)
	if err != nil {
		log.Error("failed to do post request %v, error: %v", uri, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resBody.Code != 0 {
		log.Error("failed to do post request %v, code: %v, msg: %v\n data: %v", uri, resBody.Code, resBody.Msg, resBody.Data)
		return nil, fmt.Errorf("response code: %v", resBody.Code)
	}
	log.Info("successfully do post request %v, message: %v, code: %v", uri, resBody.Msg, resBody.Code)

	return resBody.Data, nil
}
