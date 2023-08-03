package starch

import (
	"bytes"
	"fmt"
)

type MemoryContext struct {
	Headers map[string]string
	Params  map[string]string
	Sink    bytes.Buffer
	Source  bytes.Buffer
	Status  int
	Vars    map[string]any
}

func NewMemoryContext() *MemoryContext {
	return &MemoryContext{
		Headers: map[string]string{},
		Params:  map[string]string{},
		Vars:    map[string]any{},
	}
}

func (t *MemoryContext) Var(key string) any {
	return t.Vars[key]
}

func (t *MemoryContext) SetVar(key string, value any) {
	t.Vars[key] = value
}

func (t *MemoryContext) Read(buff []byte) (int, error) {
	return t.Source.Read(buff)
}

func (t *MemoryContext) Write(buff []byte) (int, error) {
	return t.Sink.Write(buff)
}

func (t *MemoryContext) WriteHeader(key string, value string) error {
	t.Headers[key] = value
	return nil
}

func (t *MemoryContext) WriteStatus(code int) error {
	t.Status = code
	return nil
}

func (t *MemoryContext) Render(c Component) error {
	return c.Render(t)
}

func (t *MemoryContext) WriteString(s string, args ...any) error {
	_, err := t.Write([]byte(fmt.Sprintf(s, args...)))
	return err
}

func (t *MemoryContext) Redirect(string) error {
	return nil
}

func (t *MemoryContext) Param(key string) string {
	return t.Params[key]
}
