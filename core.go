package starch

type Component interface {
	Render(Context) error
}

type CtxParam []string
