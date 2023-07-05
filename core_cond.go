package starch

func (t If) Render(c Context) error {
	if t.Cond(c) {
		return t.Then.Render(c)
	}
	return nil
}
