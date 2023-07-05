package starch

type Attribute interface {
	Component
	Attribute()
}

type SelfClosing interface {
	Component
	SelfClosing()
}

type PhrasingContent interface {
	Component
	PhrasingContent()
}

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
