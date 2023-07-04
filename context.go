package starch

import "io"

type Context interface {
	io.Writer
	io.Reader
	GetParam(key string) string
	GetVar(key string) any
	Next(Component) error
	SetVar(key string, value any)
	WriteHeader(string, string)
	WriteStatus(int)
	WriteString(string, ...any) error
}

func GetVar[T any](c Context, key string, defaultv ...T) T {
	v := c.GetVar(key)
	if v == nil && len(defaultv) > 0 {
		return defaultv[0]
	}
	return v.(T)
}
