package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	client  http.Client
	headers map[string]interface{}
}

type RequestType string

const (
	REQUEST_GET_TYPE       = RequestType("Get")
	REQUEST_POST_JSON_TYPE = RequestType("PostJson")
	REQUEST_POST_FORM_TYPE = RequestType("PostForm")
)

func (r *Request) WithHeaders(headers map[string]interface{}) *Request {
	r.headers = headers
	return r
}

func (r *Request) Get(url string, params map[string]interface{}) (error, []byte) {

	// create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, nil
	}

	// build query
	if params != nil {
		query := req.URL.Query()
		for k, v := range params {
			query.Set(k, v.(string))
		}
		req.URL.RawQuery = query.Encode()
	}

	// add header
	for key, value := range r.headers {
		req.Header.Set(key, value.(string))
	}

	// send http request
	resp, err := r.client.Do(req)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}

func (r *Request) PostJson(url string, form interface{}) (error, []byte) {

	jsonStr, err := json.Marshal(form)

	if err != nil {
		return err, nil
	}

	// create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err, nil
	}

	// add header
	for key, value := range r.headers {
		req.Header.Set(key, value.(string))
	}

	req.Header.Set("Content-Type", "application/json")

	// send http request
	resp, err := r.client.Do(req)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	return nil, body
}

func (r *Request) PostForm(url string, form map[string]interface{}) (error, []byte) {

	// create request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err, nil
	}

	for k, v := range form {
		req.Form.Add(k, v.(string))
	}

	// add header
	for key, value := range r.headers {
		req.Header.Set(key, value.(string))
	}

	// send http request
	resp, err := r.client.Do(req)
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	return nil, body
}

func (r *Request) Send(requestType RequestType, url string, params map[string]interface{}, data interface{}) (error, []byte) {
	var err error
	var body []byte
	switch requestType {
	case REQUEST_POST_FORM_TYPE:
		err, body = r.PostForm(url, params)
	case REQUEST_POST_JSON_TYPE:
		err, body = r.PostJson(url, params)
	case REQUEST_GET_TYPE:
		fallthrough
	default:
		err, body = r.Get(url, params)
		break
	}
	if err != nil {
		return err, nil
	}
	if data != nil {
		json.Unmarshal(body, data)
	}
	return nil, body
}

func NewHttpClient() *Request {
	return &Request{http.Client{}, make(map[string]interface{})}
}
