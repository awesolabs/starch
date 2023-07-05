package starch

import (
	"fmt"
)

func NewMemoryContext() *MemoryContext {
	return &MemoryContext{
		Headers: map[string]string{},
		Params:  map[string]string{},
		Vars:    map[string]any{},
	}
}

func (t *MemoryContext) GetVar(key string) any {
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

func (t *MemoryContext) WriteHeader(key string, value string) {
	t.Headers[key] = value
}

func (t *MemoryContext) WriteStatus(code int) {
	t.Status = code
}

func (t *MemoryContext) Next(c Component) error {
	return c.Render(t)
}

func (t *MemoryContext) WriteString(s string, args ...any) error {
	_, err := t.Sink.Write([]byte(fmt.Sprintf(s, args...)))
	return err
}

func (t *MemoryContext) GetParam(key string) string {
	return t.Params[key]
}
