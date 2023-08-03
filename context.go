package starch

import (
	"io"
)

type Context interface {
	io.Writer
	io.Reader
	Param(key string) string
	Redirect(url string) error
	Render(Component) error
	SetVar(key string, value any)
	Var(key string) any
	WriteHeader(string, string) error
	WriteStatus(int) error
	WriteString(string, ...any) error
}
