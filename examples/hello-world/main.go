package main

import . "github.com/awesolabs/starch"

var app = App{
	Route{
		Path:   "/",
		Handle: Text{`Hello, World!`},
	},
}

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
