package tests

import (
	"testing"

	. "github.com/awesolabs/starch"
)

func Test_example(t *testing.T) {
	_ = App{
		Route{
			Path: "/",
			Handle: HTML{
				Head{
					Script{Attr{"type", "text/javascript"}, Text{`
						console.log("Hello, World!");
					`}},
				},
				Body{
					If{Cond: func(Context) bool { return true }, Then: Div{Class{"one", "two"},
						Text{`Condition is true`},
					}},
				},
			},
		},
	}
}
