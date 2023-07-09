# starch

Experimental library for declarative go web apps. This had made many people very angry and has been widely regarded as a bad move.

# Install

```go
go get -u github.com/awesolabs/starch
```

# Getting Started

in `main.go`:

```go
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
```

# Documentation

TBD. For now review various example apps under [`examples`](examples/)