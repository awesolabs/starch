package starch

import (
	"html"
	"strings"
)

type (
	A          []Component
	Abbr       []Component
	Acronym    []Component
	Address    []Component
	Applet     []Component
	Area       []Component
	Article    []Component
	Aside      []Component
	Audio      []Component
	B          []Component
	Base       []Component
	Basefont   []Component
	Bdi        []Component
	Bdo        []Component
	Bgsound    []Component
	Big        []Component
	Blink      []Component
	Blockquote []Component
	Body       []Component
	Br         []Component
	Button     []Component
	Canvas     []Component
	Caption    []Component
	Center     []Component
	Cite       []Component
	Code       []Component
	Col        []Component
	Colgroup   []Component
	Command    []Component
	Data       []Component
	Datalist   []Component
	Dd         []Component
	Del        []Component
	Details    []Component
	Dfn        []Component
	Dialog     []Component
	Dir        []Component
	Div        []Component
	Dl         []Component
	Doctype    []Component
	Dt         []Component
	Em         []Component
	Embed      []Component
	Fieldset   []Component
	Figcaption []Component
	Figure     []Component
	Font       []Component
	Footer     []Component
	Form       []Component
	Frame      []Component
	Frameset   []Component
	H1         []Component
	H2         []Component
	H3         []Component
	H4         []Component
	H5         []Component
	H6         []Component
	Head       []Component
	Header     []Component
	Hr         []Component
	HTML       []Component
	I          []Component
	Iframe     []Component
	Img        []Component
	Input      []Component
	Ins        []Component
	Kbd        []Component
	Keygen     []Component
	Label      []Component
	Legend     []Component
	Li         []Component
	Link       []Component
	Main       []Component
	Map        []Component
	Mark       []Component
	Marquee    []Component
	Menu       []Component
	Menuitem   []Component
	Meta       []Component
	Meter      []Component
	Nav        []Component
	Noframes   []Component
	Noscript   []Component
	Object     []Component
	Ol         []Component
	Optgroup   []Component
	Option     []Component
	Output     []Component
	P          []Component
	Param      []Component
	Picture    []Component
	Pre        []Component
	Progress   []Component
	Q          []Component
	RawHTML    []string
	Rp         []Component
	Rt         []Component
	Ruby       []Component
	S          []Component
	Samp       []Component
	Script     []Component
	Section    []Component
	Select     []Component
	Small      []Component
	Source     []Component
	Span       []Component
	Strike     []Component
	Strong     []Component
	Style      []CSSComponent
	Sub        []Component
	Summary    []Component
	Sup        []Component
	Svg        []Component
	Table      []Component
	Tbody      []Component
	Td         []Component
	Template   []Component
	Text       []string
	Textarea   []Component
	Tfoot      []Component
	Th         []Component
	Thead      []Component
	Time       []Component
	Title      []Component
	Tr         []Component
	Track      []Component
	Tt         []Component
	U          []Component
	Ul         []Component
	Var        []Component
	Video      []Component
	Wbr        []Component
)

func (t Area) SelfClosing()    {}
func (t Base) SelfClosing()    {}
func (t Br) SelfClosing()      {}
func (t Col) SelfClosing()     {}
func (t Command) SelfClosing() {}
func (t Embed) SelfClosing()   {}
func (t Hr) SelfClosing()      {}
func (t Img) SelfClosing()     {}
func (t Input) SelfClosing()   {}
func (t Keygen) SelfClosing()  {}
func (t Link) SelfClosing()    {}
func (t Meta) SelfClosing()    {}
func (t Param) SelfClosing()   {}
func (t Source) SelfClosing()  {}
func (t Track) SelfClosing()   {}
func (t Wbr) SelfClosing()     {}

func (t Meta) FormatBlock() {}
func (t Link) FormatBlock() {}

func (t A) Render(c Context) error        { return renderelem(c, "a", t, t) }
func (t B) Render(c Context) error        { return renderelem(c, "b", t, t) }
func (t Body) Render(c Context) error     { return renderelem(c, "body", t, t) }
func (t Button) Render(c Context) error   { return renderelem(c, "button", t, t) }
func (t Div) Render(c Context) error      { return renderelem(c, "div", t, t) }
func (t Form) Render(c Context) error     { return renderelem(c, "form", t, t) }
func (t H1) Render(c Context) error       { return renderelem(c, "h1", t, t) }
func (t H2) Render(c Context) error       { return renderelem(c, "h2", t, t) }
func (t H3) Render(c Context) error       { return renderelem(c, "h3", t, t) }
func (t H4) Render(c Context) error       { return renderelem(c, "h4", t, t) }
func (t H5) Render(c Context) error       { return renderelem(c, "h5", t, t) }
func (t H6) Render(c Context) error       { return renderelem(c, "h6", t, t) }
func (t Head) Render(c Context) error     { return renderelem(c, "head", t, t) }
func (t I) Render(c Context) error        { return renderelem(c, "i", t, t) }
func (t Input) Render(c Context) error    { return renderelem(c, "input", t, t) }
func (t Label) Render(c Context) error    { return renderelem(c, "label", t, t) }
func (t Li) Render(c Context) error       { return renderelem(c, "li", t, t) }
func (t Link) Render(c Context) error     { return renderelem(c, "link", t, t) }
func (t Meta) Render(c Context) error     { return renderelem(c, "meta", t, t) }
func (t S) Render(c Context) error        { return renderelem(c, "s", t, t) }
func (t Script) Render(c Context) error   { return renderelem(c, "script", t, t) }
func (t Section) Render(c Context) error  { return renderelem(c, "section", t, t) }
func (t Style) Render(c Context) error    { return renderelem(c, "style", t, t) }
func (t Textarea) Render(c Context) error { return renderelem(c, "textarea", t, t) }
func (t Ul) Render(c Context) error       { return renderelem(c, "ul", t, t) }

func (t HTML) Render(c Context) error {
	c.WriteHeader("Content-Type", "text/html")
	c.WriteString("<!doctype html>\n")
	return renderelem(c, "html", t, t)
}

func (t Text) Render(c Context) error {
	return c.WriteString(
		html.EscapeString(
			strings.Join(t, ""),
		),
	)
}

func (t RawHTML) Render(c Context) error {
	return c.WriteString(strings.Join(t, ""))
}

func (t Text) CSSComponent() {}

func renderelem[T Component](c Context, tag string, target Component, children []T) error {
	indent := GetVar(c, VarIndent, 0)
	indentstr := strings.Repeat("\t", indent)
	_, selfclose := target.(SelfClosing)
	var attributes []Component
	var elements []Component
	for _, child := range children {
		var child Component = child
		if _, isattr := child.(Attribute); isattr {
			attributes = append(attributes, child)
			continue
		}
		elements = append(elements, child)
	}
	c.WriteString("%s<%s", indentstr, tag)
	for _, attr := range attributes {
		if err := attr.Render(c); err != nil {
			return err
		}
	}
	if err := c.WriteString(">"); err != nil {
		return err
	}
	if !selfclose {
		c.SetVar(VarIndent, indent+1)
		for idx, elem := range elements {
			_, fmtblock := elem.(FormatBlock)
			if idx == 0 || fmtblock {
				c.WriteString("\n")
			}
			if err := elem.Render(c); err != nil {
				return err
			}
		}
		c.SetVar(VarIndent, indent)
	}
	if !selfclose {
		if err := c.WriteString("%s</%s>", indentstr, tag); err != nil {
			return err
		}
		if len(elements) > 0 {
			if err := c.WriteString("\n"); err != nil {
				return err
			}
		}
	}
	return nil
}
