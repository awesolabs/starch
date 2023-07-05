package starch

func (t Route) Render(c Context) error {
	return t.Handle.Render(c)
}
