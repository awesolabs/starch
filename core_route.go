package starch

type Route struct {
	Path   string
	Handle Component
}

func (t Route) Render(c Context) error {
	return t.Handle.Render(c)
}
