package factual

import (
	"encoding/json"
	"errors"
	"strconv"
)

type errorResp struct {
	Message string `json:"message"`
}

func ErrInvalidUrl(url string) error {
	url += ": invalid url"
	return errors.New(url)
}

func ErrHttpBody(url string) error {
	url += ": response body was malformed"
	return errors.New(url)
}

func ErrHttpResponse(url string, code int, resp []byte) error {
	e := errorResp{}
	json.Unmarshal(resp, &e)
	m := url + " (" + strconv.Itoa(code) + "): " + e.Message
	return errors.New(m)
}
