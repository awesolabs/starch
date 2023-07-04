package tests

import (
	"testing"

	. "github.com/awesolabs/starch"
	"github.com/stretchr/testify/assert"
)

func Test_attr_has_attribute_interface(t *testing.T) {
	var attr Attribute = Attr{"key", "value"}
	assert.NotNil(t, attr)
}

func Test_attr_renders_nothing_when_empty(t *testing.T) {
	attr := Attr{}
	ctx := NewMemoryContext()
	err := attr.Render(ctx)
	assert.NoError(t, err)
	result := ctx.Sink.String()
	assert.Equal(t, "", result)
}

func Test_attr_renders_single_key(t *testing.T) {
	attr := Attr{"key"}
	ctx := NewMemoryContext()
	err := attr.Render(ctx)
	assert.NoError(t, err)
	result := ctx.Sink.String()
	assert.Equal(t, " key", result)
}

func Test_attr_renders_key_and_value(t *testing.T) {
	attr := Attr{"key", "value"}
	ctx := NewMemoryContext()
	err := attr.Render(ctx)
	assert.NoError(t, err)
	result := ctx.Sink.String()
	assert.Equal(t, ` key="value"`, result)
}

func Test_attr_renders_key_and_multi_value(t *testing.T) {
	attr := Attr{"key", "value1", "value2"}
	ctx := NewMemoryContext()
	err := attr.Render(ctx)
	assert.NoError(t, err)
	result := ctx.Sink.String()
	assert.Equal(t, ` key="value1 value2"`, result)
}

func Test_class_has_attribute_interface(t *testing.T) {
	var class Attribute = Class{"foo", "bar"}
	assert.NotNil(t, class)
}

func Test_class_renders_nothing_when_empty(t *testing.T) {
	attr := Class{}
	ctx := NewMemoryContext()
	err := attr.Render(ctx)
	assert.NoError(t, err)
	result := ctx.Sink.String()
	assert.Equal(t, ``, result)
}

func Test_class_renders_key_and_value(t *testing.T) {
	attr := Class{"value"}
	ctx := NewMemoryContext()
	err := attr.Render(ctx)
	assert.NoError(t, err)
	result := ctx.Sink.String()
	assert.Equal(t, ` class="value"`, result)
}

func Test_class_renders_key_and_multi_value(t *testing.T) {
	attr := Class{"value1", "value2"}
	ctx := NewMemoryContext()
	err := attr.Render(ctx)
	assert.NoError(t, err)
	result := ctx.Sink.String()
	assert.Equal(t, ` class="value1 value2"`, result)
}
