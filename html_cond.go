package starch

type AttrIf struct {
	Cond CondFunc
	Then Component
}

func (t AttrIf) Attribute() {}

func (t AttrIf) Render(c Context) error {
	if t.Cond(c) {
		return t.Then.Render(c)
	}
	return nil
}
