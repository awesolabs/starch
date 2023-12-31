package tests

import (
	"testing"

	. "github.com/awesolabs/starch"
	"github.com/stretchr/testify/assert"
)

func Test_element_render_smoketest(t *testing.T) {
	html := HTML{
		Head{},
		Body{
			H1{Text{`heading`}},
			H2{Text{`heading`}},
			H3{Text{`heading`}},
			H4{Text{`heading`}},
			H5{Text{`heading`}},
			H6{Text{`heading`}},
			Div{},
		},
	}
	ctx := NewMemoryContext()
	assert.NoError(t, html.Render(ctx))
	result := ctx.Sink.String()
	assert.NotEmpty(t, result)
}

func Benchmark_element_render_smoketest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ctx := NewNoopContext()
		html := HTML{
			Head{},
			Body{
				Div{},
			},
		}
		html.Render(ctx)
	}
	b.ReportAllocs()
}

func Test_element_renders_with_no_children(t *testing.T) {
	html := HTML{}
	ctx := NewMemoryContext()
	assert.NoError(t, html.Render(ctx))
	result := ctx.Sink.String()
	assert.Equal(t, "<html></html>", result)
}

func Test_element_renders_with_attributes(t *testing.T) {
	html := HTML{Attr{"foo", "bar"}}
	ctx := NewMemoryContext()
	assert.NoError(t, html.Render(ctx))
	result := ctx.Sink.String()
	assert.Equal(t, `<html foo="bar"></html>`, result)
}

func Test_element_renders_with_child(t *testing.T) {
	html := HTML{
		Head{},
	}
	ctx := NewMemoryContext()
	assert.NoError(t, html.Render(ctx))
	result := ctx.Sink.String()
	assert.Equal(t, `<html><head></head></html>`, result)
}

func Test_input_renders_attributes(t *testing.T) {
	elem := Input{TypeButton}
	ctx := NewMemoryContext()
	assert.NoError(t, elem.Render(ctx))
	result := ctx.Sink.String()
	assert.Equal(t, `<input type="button">`, result)
}
