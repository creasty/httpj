package httpj

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Request struct {
	Headers map[string]string
	Method  string
	Url     string
}

func NewRequest(url string) *Request {
	req := &Request{
		Url:     url,
		Headers: make(map[string]string),
	}
	req.setDefault()
	return req
}

func (self *Request) setDefault() {
	self.SetHeader("Content-Type", "application/json")
}

func (self *Request) SetHeader(key, value string) *Request {
	self.Headers[key] = value
	return self
}

func (self *Request) Get(body interface{}) (*Response, error) {
	self.Method = "GET"
	return self.send(body)
}
func (self *Request) Post(body interface{}) (*Response, error) {
	self.Method = "POST"
	return self.send(body)
}
func (self *Request) Patch(body interface{}) (*Response, error) {
	self.Method = "PATCH"
	return self.send(body)
}
func (self *Request) Put(body interface{}) (*Response, error) {
	self.Method = "PUT"
	return self.send(body)
}
func (self *Request) Delete(body interface{}) (*Response, error) {
	self.Method = "DELETE"
	return self.send(body)
}
func (self *Request) Head(body interface{}) (*Response, error) {
	self.Method = "HEAD"
	return self.send(body)
}

func (self *Request) send(body interface{}) (*Response, error) {
	var reqBody io.Reader = nil

	if body != nil {
		switch body.(type) {
		case string:
			reqBody = bytes.NewBuffer([]byte(body.(string)))
			break
		default:
			s, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			reqBody = bytes.NewBuffer(s)
		}
	}

	req, err := http.NewRequest(self.Method, self.Url, reqBody)
	if err != nil {
		return nil, err
	}

	for k, v := range self.Headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return &Response{resp}, nil
}
