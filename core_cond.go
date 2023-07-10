package starch

type CondFunc func(c Context) bool

type If struct {
	Cond CondFunc
	Then Component
	Else Component
}

func (t If) Render(c Context) error {
	if t.Then != nil && t.Cond(c) {
		return t.Then.Render(c)
	} else if t.Else != nil {
		return t.Else.Render(c)
	}
	return nil
}
