package starch

import "strings"

type (
	Action    []string
	Attr      []string
	Class     []string
	Href      []string
	ID        []string
	Method    []string
	Name      []string
	Rel       []string
	StyleAttr []string
	Target    []string
	Type      []string
	Value     []string
)

var (
	TargetBlank       = Target{"__blank"}
	TargetParent      = Target{"__parent"}
	TargetSelf        = Target{"__self"}
	TargetTop         = Target{"__top"}
	TypeButton        = Type{"button"}
	TypeCheckbox      = Type{"checkbox"}
	TypeColor         = Type{"color"}
	TypeDate          = Type{"date"}
	TypeDateTime      = Type{"datetime"}
	TypeDateTimeLocal = Type{"datetime-local"}
	TypeEmail         = Type{"email"}
	TypeFile          = Type{"file"}
	TypeHidden        = Type{"hidden"}
	TypeImage         = Type{"image"}
	TypeJavascript    = Type{"text/javascript"}
	TypeMonth         = Type{"month"}
	TypeNumber        = Type{"number"}
	TypePassword      = Type{"password"}
	TypeRadio         = Type{"radio"}
	TypeRange         = Type{"range"}
	TypeReset         = Type{"reset"}
	TypeSearch        = Type{"search"}
	TypeSubmit        = Type{"submit"}
	TypeTel           = Type{"tel"}
	TypeText          = Type{"text"}
	TypeTime          = Type{"time"}
	TypeURL           = Type{"url"}
	TypeWeek          = Type{"week"}
	RelExternal       = Rel{"external"}
	RelHelp           = Rel{"help"}
	RelLicense        = Rel{"license"}
	RelNext           = Rel{"next"}
	RelNoFollow       = Rel{"nofollow"}
	RelNoOpener       = Rel{"noopener"}
	RelNoReferrer     = Rel{"noreferrer"}
	RelOpener         = Rel{"opener"}
	RelPrev           = Rel{"prev"}
	RelSearch         = Rel{"search"}
	MethodGET         = Method{"get"}
	MethodPOST        = Method{"post"}
)

func (t Action) Attribute()    {}
func (t Attr) Attribute()      {}
func (t Class) Attribute()     {}
func (t Href) Attribute()      {}
func (t ID) Attribute()        {}
func (t Method) Attribute()    {}
func (t Name) Attribute()      {}
func (t Rel) Attribute()       {}
func (t StyleAttr) Attribute() {}
func (t Target) Attribute()    {}
func (t Type) Attribute()      {}
func (t Value) Attribute()     {}

func (t Action) AttrValueMerge() {}
func (t Href) AttrValueMerge()   {}
func (t Value) AttrValueMerge()  {}

func (t Action) Render(c Context) error    { return renderattr(c, "action", t, t...) }
func (t Attr) Render(c Context) error      { return renderattr(c, t[0], t, t[1:]...) }
func (t Class) Render(c Context) error     { return renderattr(c, "class", t, t...) }
func (t Href) Render(c Context) error      { return renderattr(c, "href", t, t...) }
func (t ID) Render(c Context) error        { return renderattr(c, "id", t, t...) }
func (t Method) Render(c Context) error    { return renderattr(c, "method", t, t...) }
func (t Name) Render(c Context) error      { return renderattr(c, "name", t, t...) }
func (t Rel) Render(c Context) error       { return renderattr(c, "rel", t, t...) }
func (t StyleAttr) Render(c Context) error { return renderattr(c, "style", t, t...) }
func (t Target) Render(c Context) error    { return renderattr(c, "target", t, t...) }
func (t Type) Render(c Context) error      { return renderattr(c, "type", t, t...) }
func (t Value) Render(c Context) error     { return renderattr(c, "value", t, t...) }

func renderattr(c Context, key string, target Component, values ...string) error {
	switch len(values) {
	case 0:
		return c.WriteString(` %s`, key)
	case 1:
		return c.WriteString(` %s="%s"`, key, values[0])
	default:
		switch target.(type) {
		case AttrValueMerge:
			return c.WriteString(
				` %s="%s"`,
				key,
				strings.Join(values, ""),
			)
		default:
			return c.WriteString(
				` %s="%s"`,
				key,
				strings.Join(values, " "),
			)
		}
	}
}
