package starch

import (
	"bytes"
	"io"

	"github.com/ucarion/urlpath"
)

type Context interface {
	io.Writer
	io.Reader
	GetParam(key string) string
	GetVar(key string) any
	Redirect(url string)
	SetVar(key string, value any)
	WriteHeader(string, string)
	WriteStatus(int)
	WriteString(string, ...any) error
}

type AppContext struct {
	Routes map[*urlpath.Path]Route
}

type MemoryContext struct {
	Headers map[string]string
	Params  map[string]string
	Sink    bytes.Buffer
	Source  bytes.Buffer
	Status  int
	Vars    map[string]any
}

type NoopContext struct{}
