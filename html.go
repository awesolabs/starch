package starch

import "strings"

type (
	Attr   []string
	Body   []Component
	Class  []string
	Div    []Component
	H1     []Component
	H2     []Component
	H3     []Component
	H4     []Component
	H5     []Component
	H6     []Component
	Head   []Component
	Html   []Component
	Script []Component
	Text   []string
)

type SelfClosing interface {
	Component
	SelfClosing()
}

type PhrasingContent interface {
	Component
	PhrasingContent()
}

func (t H1) Render(c Context) error     { return renderelem(c, "h1", t, t) }
func (t H2) Render(c Context) error     { return renderelem(c, "h2", t, t) }
func (t H3) Render(c Context) error     { return renderelem(c, "h3", t, t) }
func (t H4) Render(c Context) error     { return renderelem(c, "h4", t, t) }
func (t H5) Render(c Context) error     { return renderelem(c, "h5", t, t) }
func (t H6) Render(c Context) error     { return renderelem(c, "h6", t, t) }
func (t Head) Render(c Context) error   { return renderelem(c, "head", t, t) }
func (t Script) Render(c Context) error { return renderelem(c, "script", t, t) }
func (t Div) Render(c Context) error    { return renderelem(c, "div", t, t) }
func (t Body) Render(c Context) error   { return renderelem(c, "body", t, t) }

func (t Html) Render(c Context) error {
	c.WriteHeader("Content-Type", "text/html")
	return renderelem(c, "html", t, t)
}

func (t Text) Render(c Context) error {
	return c.WriteString(strings.Join(t, ""))
}

func renderelem(c Context, tag string, target Component, children []Component) error {
	indent := GetVar(c, VarIndent, 0)
	indentstr := strings.Repeat("\t", indent)
	c.WriteString("%s<%s", indentstr, tag)
	attrclosed := false
	childcount := 0
	_, selfclose := target.(SelfClosing)
	if selfclose {
		children = nil
	}
	for a := 0; a < len(children); a++ {
		child := children[a]
		_, isattr := child.(Attribute)
		if !isattr && !attrclosed {
			attrclosed = true
			if err := c.WriteString(">"); err != nil {
				return err
			}
			if len(children) > a {
				if err := c.WriteString("\n"); err != nil {
					return err
				}
			}
		}
		if !isattr {
			c.SetVar(VarIndent, indent+1)
			childcount += 1
		}
		if err := child.Render(c); err != nil {
			return err
		}
		if !isattr {
			c.SetVar(VarIndent, indent)
		}
	}
	if !attrclosed {
		if err := c.WriteString(">"); err != nil {
			return err
		}
	}
	if !selfclose {
		if err := c.WriteString("%s</%s>", indentstr, tag); err != nil {
			return err
		}
		if childcount > 0 {
			if err := c.WriteString("\n"); err != nil {
				return err
			}
		}
	}
	return nil
}
