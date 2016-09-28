package httpj

import (
	"net/http"
	"time"
)

const (
	HTTP_CLIENT_MAX_CONNECTIONS = 20
	HTTP_CLIENT_REQUEST_TIMEOUT = 5
)

type Client struct {
	*http.Client
}

func New(maxConn int, timeout time.Duration) *Client {
	c := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxConn,
		},
		Timeout: timeout * time.Second,
	}
	return &Client{c}
}

func Default(maxConn int, timeout time.Duration) *Client {
	return New(HTTP_CLIENT_MAX_CONNECTIONS, time.Duration(HTTP_CLIENT_REQUEST_TIMEOUT))
}

func (self *Client) NewRequest(url string) *Request {
	return NewRequest(self.Client, url)
}
