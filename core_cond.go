package starch

type CondFunc func(c Context) bool

type If struct {
	Cond CondFunc
	Then Component
}

func (t If) Render(c Context) error {
	if t.Cond(c) {
		return t.Then.Render(c)
	}
	return nil
}
