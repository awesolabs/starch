package starch

import "net/http"

func (t App) Run() error {
	ctx := NewAppContext()
	for _, child := range t {
		if route, ok := child.(Route); ok {
			ctx.AddRoute(route)
		}
	}
	return http.ListenAndServe(":8080", ctx)
}
