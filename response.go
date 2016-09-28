package httpj

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

func (self *Response) IsSuccess() bool {
	code := self.StatusCode
	return 200 <= code && code < 400
}

func (self *Response) Bind(obj interface{}) error {
	body, err := ioutil.ReadAll(self.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}
