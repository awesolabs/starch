package starch

type AttrIf struct {
	Cond CondFunc
	Then Attribute
	Else Attribute
}

func (t AttrIf) Attribute() {}

func (t AttrIf) Render(c Context) error {
	if t.Then != nil && t.Cond(c) {
		return t.Then.Render(c)
	} else if t.Else != nil {
		return t.Else.Render(c)
	}
	return nil
}
