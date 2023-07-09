package starch

type Component interface {
	Render(Context) error
}

type App []Component

type CtxParam []string

type RenderFunc []func(c Context) error

type Route struct {
	Path   string
	Handle Component
}
