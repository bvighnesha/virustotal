package model

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Headers    *http.Header
	Parameters *url.Values
	Body       io.Reader
}

type Response struct {
	Error *Error `json:"error,omitempty"`
	Data  *Data  `json:"data,omitempty"`
}
