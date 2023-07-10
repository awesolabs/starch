package starch

import (
	"net/http"

	"github.com/ucarion/urlpath"
)

func NewAppContext() *AppContext {
	return &AppContext{
		Routes: map[*urlpath.Path]Route{},
	}
}

func (t *AppContext) AddRoute(r Route) {
	path := urlpath.New(r.Path)
	t.Routes[&path] = r
}

func (t *AppContext) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for path, route := range t.Routes {
		match, ok := path.Match(r.URL.Path)
		if !ok {
			continue
		}
		hctx := NewHTTPContext().With(func(hc *HTTPContext) {
			hc.params = match.Params
			hc.request = r
			hc.response = w
			hc.route = route
		})
		if err := route.Render(hctx); err != nil {
			panic(err)
		}
		return
	}
	http.DefaultServeMux.ServeHTTP(w, r)
}
