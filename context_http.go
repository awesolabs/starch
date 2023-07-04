package starch

import (
	"fmt"
	"net/http"
)

type HttpContext struct {
	route    Route
	request  *http.Request
	response http.ResponseWriter
	params   map[string]string
	vars     map[string]any
}

func NewHttpContext() *HttpContext {
	return &HttpContext{
		params: map[string]string{},
		vars:   map[string]any{},
	}
}

func (t *HttpContext) GetVar(key string) any {
	return t.vars[key]
}

func (t *HttpContext) SetVar(key string, value any) {
	t.vars[key] = value
}

func (t *HttpContext) Read(buff []byte) (int, error) {
	return t.request.Body.Read(buff)
}

func (t *HttpContext) Write(buff []byte) (int, error) {
	return t.response.Write(buff)
}

func (t *HttpContext) WriteHeader(key string, value string) {
	t.response.Header().Add(key, value)
}

func (t *HttpContext) WriteStatus(code int) {
	t.response.WriteHeader(code)
}

func (t *HttpContext) WriteString(s string, args ...any) error {
	_, err := t.response.Write([]byte(fmt.Sprintf(s, args...)))
	return err
}

func (t *HttpContext) Next(c Component) error {
	return c.Render(t)
}

func (t *HttpContext) With(f func(*HttpContext)) *HttpContext {
	f(t)
	return t
}

func (t *HttpContext) GetParam(key string) string {
	param, ok := t.params[key]
	if !ok {
		param = t.request.URL.Query().Get(key)
	}
	return param
}
