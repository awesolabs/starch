package starch

type CSSComponent interface {
	Component
	CSSComponent()
}

type CSSRule struct {
	AlignItems      string
	BackgroundColor string
	Color           string
	Display         string
	JustifyContent  string
	Left            string
	Margin          string
	PlaceItems      string
	Position        string
	Top             string
	Transform       string
	Width           string
}

type CSSSelector map[string]CSSRule

func (t CSSSelector) CSSComponent() {}

func (t CSSSelector) Render(c Context) error { return nil }

func (t CSSRule) CSSComponent() {}
