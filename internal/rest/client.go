package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Credentials struct {
	AppKey   string
	Basic    string
	Bearer   string
	Endpoint string
}

func (c Credentials) Token() string {
	if strings.TrimSpace(c.Basic) != "" {
		return "Basic " + c.Basic
	}
	return "Bearer " + c.Bearer
}

type Response struct {
	Body []byte
	Code int
}

type RestClient interface {
	FormData(path string) (Response, error)
	Get(path string) (Response, error)
	Post(path string, value interface{}) (Response, error)
}

type restClient struct {
	credentials Credentials
	params      url.Values
}

func NewClient(c Credentials, params url.Values) RestClient {
	return restClient{
		credentials: c,
		params:      params}
}

func (r restClient) FormData(path string) (Response, error) {
	contentType := "application/x-www-form-urlencoded"
	return r.send(http.MethodPost, path, contentType, strings.NewReader(r.params.Encode()))
}

func (r restClient) Get(path string) (Response, error) {
	contentType := "application/json; charset=utf-8"
	path += r.queryString()
	return r.send(http.MethodGet, path, contentType, nil)
}

func (r restClient) Post(path string, value interface{}) (Response, error) {
	body, err := json.Marshal(value)
	if err != nil {
		return Response{}, err
	}
	path += r.queryString()
	contentType := "application/json; charset=utf-8"
	return r.send(http.MethodPost, path, contentType, bytes.NewReader(body))
}

func (r restClient) send(method, path, contentType string, body io.Reader) (Response, error) {
	url := r.credentials.Endpoint + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{}, err
	}

	req.Header.Add("Content-type", contentType)
	req.Header.Add("Authorization", r.credentials.Token())

	httpClient := &http.Client{Timeout: 15 * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Body: content,
		Code: res.StatusCode,
	}, nil
}

func (r restClient) queryString() string {
	qs := r.params.Encode()
	if qs != "" {
		qs += "&"
	}
	return fmt.Sprintf("?%sgw-dev-app-key=%s",
		qs,
		r.credentials.AppKey,
	)
}
