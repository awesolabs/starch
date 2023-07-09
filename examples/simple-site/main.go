package main

import . "github.com/awesolabs/starch"

var app = App{
	Route{
		Path: "/",
		Handle: &Layout{
			Content: Div{
				H3{Text{`Home Page`}},
				If{Cond: hasName, Then: Div{
					H4{CtxParam{"name"}},
				}},
			},
		},
	},
	Route{
		Path: "/about",
		Handle: &Layout{
			Content: Div{
				H3{Text{`About Page`}},
			},
		},
	},
}

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

type Layout struct {
	Content Component
}

func (t *Layout) Render(c Context) error {
	template := HTML{
		Head{},
		Body{
			H1{Text{`Simple Website`}},
			Div{
				A{Href{"/"}, Text{`Home`}},
				A{Href{"/about"}, Text{`About`}},
			},
			t.Content,
		},
	}
	return template.Render(c)
}

func hasName(c Context) bool {
	if c.GetParam("name") != "" {
		return true
	}
	return false
}
