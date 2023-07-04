package starch

import "strings"

type Attribute interface {
	Component
	Attribute()
}

func (t Attr) Attribute() {
}

func (t Attr) Render(c Context) error {
	if len(t) == 0 {
		return nil
	}
	return renderattr(c, t[0], t[1:]...)
}

func (t Class) Attribute() {
}

func (t Class) Render(c Context) error {
	if len(t) == 0 {
		return nil
	}
	return renderattr(c, "class", t...)
}

func renderattr(c Context, key string, values ...string) error {
	switch len(values) {
	case 0:
		return c.WriteString(` %s`, key)
	case 1:
		return c.WriteString(` %s="%s"`, key, values[0])
	default:
		return c.WriteString(
			` %s="%s"`,
			key,
			strings.Join(values, " "),
		)
	}
}
