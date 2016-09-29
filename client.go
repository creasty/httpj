package httpj

import (
	"net/http"
	"time"
)

const (
	HTTP_CLIENT_MAX_CONNECTIONS = 20
	HTTP_CLIENT_REQUEST_TIMEOUT = time.Duration(5)
)

type Client struct {
	*http.Client
	UrlPrefix string
}

func New() *Client {
	return &Client{
		Client: &http.Client{
			Timeout: 0,
		},
	}
}

func Default() *Client {
	return New().
		SetMaxConnection(HTTP_CLIENT_MAX_CONNECTIONS).
		SetTimeout(HTTP_CLIENT_REQUEST_TIMEOUT)
}

func (self *Client) SetTimeout(timeout time.Duration) *Client {
	self.Client.Timeout = timeout
	return self
}

func (self *Client) SetMaxConnection(num int) *Client {
	self.Client.Transport = &http.Transport{
		MaxIdleConnsPerHost: num,
	}
	return self
}

func (self *Client) SetPrefix(prefix string) *Client {
	self.UrlPrefix = prefix
	return self
}

func (self *Client) NewRequest(url string) *Request {
	return NewRequest(self.Client, self.UrlPrefix+url)
}
