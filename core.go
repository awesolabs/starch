package starch

type Component interface {
	Render(Context) error
}

type App []Component

type If struct {
	Cond func(Context) bool
	Then Component
}

type Param []string

type RenderFunc []func(c Context) error

type Route struct {
	Path   string
	Handle Component
}
