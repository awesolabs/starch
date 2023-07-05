package starch

func (t RenderFunc) Render(c Context) error {
	for _, f := range t {
		if e := f(c); e != nil {
			return e
		}
	}
	return nil
}
