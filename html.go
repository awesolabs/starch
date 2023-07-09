package starch

type Attribute interface {
	Component
	Attribute()
}

type AttrValueMerge interface {
	Component
	AttrValueMerge()
}

type SelfClosing interface {
	Component
	SelfClosing()
}

type PhrasingContent interface {
	Component
	PhrasingContent()
}
