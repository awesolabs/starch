package main

import (
	"log"

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

var (
	TodoAddEndpoint    = "/todo/add"
	TodoDoneEndpoint   = "/todo/done"
	TodoDeleteEndpoint = "/todo/delete"
)

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
			Body{
				Div{Class{"container py-5 h-100"},
					Div{Class{"row d-flex justify-content-center align-items-center h-100"},
						Div{Class{"col col-xl-10"},
							Div{Class{"card"}, StyleAttr{"border-radius: 15px"},
								Div{Class{"card-body p-5"},
									Form{Class{"d-flex justify-content-center align-items-top mb-4"},
										Div{Class{"form-outline flex-fill"},
											Input{
												TypeText,
												Name{"title"},
												Class{"form-control form-control-lg"},
												Placeholder{"enter a title for your todo"},
											},
										},
										Button{
											TypeSubmit,
											FormMethodPOST,
											FormAction{TodoAddEndpoint},
											Class{"btn btn-primary btn-lg ms-2"},
											Text{"Add"},
										},
									},
									Ul{Class{"list-group mb-0"},
										Each[*Todo]{Items: &data.Todos, Thenf: func(t *Todo) Component {
											done := isTodoDone(t, true)
											return Li{Class{"list-group-item d-flex justify-content-between align-items-center border-start-0 border-top-0 border-end-0 border-bottom rounded-0 mb-2"},
												Form{Class{"d-flex align-items-center"}, MethodPOST, Action{TodoDoneEndpoint, "?title=", t.Title},
													Input{
														TypeCheckbox,
														Class{"form-check-input me-2"},
														Attr{"onChange", "this.form.submit()"},
														AttrIf{Cond: done, Then: Attr{"checked"}},
													},
													If{Cond: done,
														Then: S{Text{t.Title}},
														Else: Text{t.Title},
													},
												},
												Form{MethodPOST, Action{TodoDeleteEndpoint, "?title=", t.Title},
													Button{TypeSubmit, StyleAttr{"border:0;background:none;"},
														I{Class{"bx bxs-x-square"}},
													},
												},
											}
										}},
									},
								},
							},
						},
					},
				},
			},
		},
	},
	Route{
		Path: TodoDoneEndpoint,
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
		Path: TodoAddEndpoint,
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
	Route{
		Path: TodoDeleteEndpoint,
		Handle: RenderFunc{func(c Context) error {
			title := c.GetParam("title")
			log.Println(title)
			for i, todo := range data.Todos {
				if todo.Title == title {
					copy(data.Todos[i:], data.Todos[i+1:])
					data.Todos = data.Todos[:len(data.Todos)-1]
				}
			}
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
