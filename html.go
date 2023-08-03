package starch

type SelfClosing interface {
	Component
	SelfClosing()
}

type FormatBlock interface {
	FormatBlock()
}

type PhrasingContent interface {
	Component
	PhrasingContent()
}
