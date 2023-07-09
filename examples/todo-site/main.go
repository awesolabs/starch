package main

import (
	"log"
	_ "net/http/pprof"

	. "github.com/awesolabs/starch"
)

type AppData struct {
	Todos []*Todo
}

type Todo struct {
	Title string
	Done  bool
}

var data = &AppData{
	Todos: []*Todo{
		{Title: "Do the dishes"},
		{Title: "Take out trash"},
	},
}

var app = App{
	Route{
		Path: "/",
		Handle: HTML{
			Head{
				Meta{CharsetUTF8},
				Meta{Name{"viewport"}, Attr{"content", "width=device-width, initial-scale=1.0"}},
				Link{
					RelStylesheet,
					Href{"https://unpkg.com/boxicons@2.1.4/css/boxicons.min.css"},
				},
				Link{
					RelStylesheet,
					CorsAnonymous,
					Href{"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css"},
					Integrity{"sha384-9ndCyUaIbzAi2FUVXJi0CjmCapSmO7SnpJef0486qhLnuZ2cdeRhO02iuK6FUUVM"},
				},
				Script{
					TypeJavascript,
					CorsAnonymous,
					Src{"https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"},
					Integrity{"sha384-geWF76RCwLtnZ8qwWowPQNguL3RmwHVBC9FhGdlKrxdiJJigb/j/68SIy3Te4Bkz"},
				},
			},
			Body{Class{"container px-4 text-center"},
				Div{
					H1{Text{`Todo Application`}},
				},
				Div{Class{"container"},
					Form{MethodPOST, Action{"/add-todo"},
						Input{TypeText, Name{"title"}},
						Input{TypeSubmit, Value{`Add`}},
					},
				},
				Each[*Todo]{Items: &data.Todos, Thenf: func(i *Todo) Component {
					return Div{Class{"container"},
						Div{
							Form{MethodPOST, Action{"/mark-done?title=", i.Title},
								If{Cond: isTodoDone(i, true), Then: Elems{
									S{Text{i.Title}},
									Button{TypeSubmit, Class{"btn"},
										I{Class{"bx bx-checkbox-checked"}},
									},
								}},
								If{Cond: isTodoDone(i, false), Then: Elems{
									Text{i.Title},
									Button{TypeSubmit, Class{"btn"},
										I{Class{"bx bx-checkbox"}},
									},
								}},
							},
						},
					}
				}},
			},
		},
	},
	Route{
		Path: "/mark-done",
		Handle: RenderFunc{func(c Context) error {
			title := c.GetParam("title")
			for _, todo := range data.Todos {
				if todo.Title == title {
					todo.Done = !todo.Done
				}
			}
			c.Redirect("/")
			return nil
		}},
	},
	Route{
		Path: "/add-todo",
		Handle: RenderFunc{func(c Context) error {
			title := c.GetParam("title")
			log.Println(title)
			data.Todos = append(data.Todos, &Todo{
				Title: title,
			})
			c.Redirect("/")
			return nil
		}},
	},
}

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func isTodoDone(t *Todo, check bool) CondFunc {
	return func(c Context) bool {
		return t.Done == check
	}
}
