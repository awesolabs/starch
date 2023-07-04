package starch

import "runtime/debug"

type If struct {
	Cond func(Context) bool
	Then Component
}

type Default struct {
	Then Component
}

func (t If) Render(c Context) error {
	if t.Cond(c) {
		debug.PrintStack()
		return c.Next(t.Then)
	}
	return nil
}
func (t Default) Render(c Context) error { return nil }
