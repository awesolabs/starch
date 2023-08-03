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

func (t *HTTPContext) Var(key string) any {
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

func (t *HTTPContext) WriteHeader(key string, value string) error {
	t.response.Header().Add(key, value)
	return nil
}

func (t *HTTPContext) WriteStatus(code int) error {
	t.response.WriteHeader(code)
	return nil
}

func (t *HTTPContext) WriteString(s string, args ...any) error {
	_, err := t.Write([]byte(fmt.Sprintf(s, args...)))
	return err
}

func (t *HTTPContext) With(f func(*HTTPContext)) *HTTPContext {
	f(t)
	return t
}

func (t *HTTPContext) Redirect(url string) error {
	http.Redirect(t.response, t.request, url, http.StatusMovedPermanently)
	return nil
}

func (t *HTTPContext) Render(c Component) error {
	return c.Render(t)
}

func (t *HTTPContext) Param(key string) string {
	param, ok := t.params[key]
	if !ok {
		param = t.request.URL.Query().Get(key)
	}
	if param == "" {
		param = t.request.FormValue(key)
	}
	return param
}
