package main

import . "github.com/awesolabs/starch"

var app = App{
	Route{
		Path: "/",
		Handle: Layout{
			Content: Div{
				H3{Text{`Home Page`}},
				If{Cond: hasName, Then: Div{
					H4{Param{"name"}},
				}},
			},
		},
	},
	Route{
		Path: "/about",
		Handle: Layout{
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

func (t Layout) Render(c Context) error {
	template := Html{
		Head{},
		Body{
			H1{Text{`Simple Website`}},
			t.Content,
		},
	}
	return c.Next(template)
}

func hasName(c Context) bool {
	if c.GetParam("name") != "" {
		return true
	}
	return false
}
