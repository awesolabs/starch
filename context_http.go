package starch

import (
	"fmt"
	"net/http"
)

type HTTPContext struct {
	route    Route
	request  *http.Request
	response http.ResponseWriter
	params   map[string]string
	vars     map[string]any
}

func NewHTTPContext() *HTTPContext {
	return &HTTPContext{
		params: map[string]string{},
		vars:   map[string]any{},
	}
}

func (t *HTTPContext) GetVar(key string) any {
	return t.vars[key]
}

func (t *HTTPContext) SetVar(key string, value any) {
	t.vars[key] = value
}

func (t *HTTPContext) Read(buff []byte) (int, error) {
	return t.request.Body.Read(buff)
}

func (t *HTTPContext) Write(buff []byte) (int, error) {
	return t.response.Write(buff)
}

func (t *HTTPContext) WriteHeader(key string, value string) {
	t.response.Header().Add(key, value)
}

func (t *HTTPContext) WriteStatus(code int) {
	t.response.WriteHeader(code)
}

func (t *HTTPContext) WriteString(s string, args ...any) error {
	_, err := t.response.Write([]byte(fmt.Sprintf(s, args...)))
	return err
}

func (t *HTTPContext) With(f func(*HTTPContext)) *HTTPContext {
	f(t)
	return t
}

func (t *HTTPContext) Redirect(url string) {
	http.Redirect(t.response, t.request, url, http.StatusMovedPermanently)
}

func (t *HTTPContext) GetParam(key string) string {
	param, ok := t.params[key]
	if !ok {
		param = t.request.URL.Query().Get(key)
	}
	if param == "" {
		param = t.request.FormValue(key)
	}
	return param
}
