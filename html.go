package starch

type SelfClosing interface {
	Component
	SelfClosing()
}

type PhrasingContent interface {
	Component
	PhrasingContent()
}
