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
				Style{Text{`
					.app-container {
						display: flex;
						flex-direction: column;
  						justify-content: center;
  						align-items: center;						
					}
				`}},
			},
			Body{Class{"app-container"},
				Div{
					H1{Text{`Todo Application`}},
				},
				Div{
					Form{MethodPOST, Action{"/add-todo"},
						Input{TypeText, Name{"title"}},
						Input{TypeSubmit, Value{`Add`}},
					},
				},
				Each[*Todo]{Items: &data.Todos, Then: func(i *Todo) Component {
					return Div{
						H2{
							Form{MethodPOST, Action{"/mark-done?title=", i.Title},
								If{Cond: isTodoDone(i, true), Then: S{Text{i.Title}}},
								If{Cond: isTodoDone(i, false), Then: Text{i.Title}},
								Input{TypeSubmit, Value{`Done`}},
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
