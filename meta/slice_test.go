package meta_test

import (
	"testing"

	"github.com/azer/crud/v2/meta"
	"github.com/stretchr/testify/assert"
)

func TestIsSlice(t *testing.T) {
	assert.False(t, meta.IsSlice(123))
	assert.False(t, meta.IsSlice("yo"))
	assert.False(t, meta.IsSlice(Foo{}))
	assert.True(t, meta.IsSlice(FooSlice{}))
	assert.True(t, meta.IsSlice(&FooSlice{}))
	assert.True(t, meta.IsSlice(FooPTRSlice{}))
	assert.True(t, meta.IsSlice(&FooPTRSlice{}))
}

func TestPushToSlice(t *testing.T) {
	sl := []Foo{}
	slval := meta.Addressable(&sl)
	meta.Push(slval, meta.ValueOf(Foo{123}))
	meta.Push(slval, meta.ValueOf(Foo{456}))

	assert.Equal(t, len(sl), 2)
	assert.Equal(t, sl[0].Id, 123)
	assert.Equal(t, sl[1].Id, 456)
}

func TestPushToPTRSlice(t *testing.T) {
	sl := []*Foo{}
	slval := meta.Addressable(&sl)
	meta.Push(slval, meta.DirectValueOf(&Foo{123}))
	meta.Push(slval, meta.DirectValueOf(&Foo{456}))

	assert.Equal(t, len(sl), 2)
	assert.Equal(t, sl[0].Id, 123)
	assert.Equal(t, sl[1].Id, 456)
}

func TestSliceOf(t *testing.T) {
	a := meta.SliceOf(123).Interface().([]int)
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0], 123)

	b := meta.SliceOf(Foo{}).Interface().([]Foo)
	assert.Equal(t, len(b), 1)
	assert.Equal(t, meta.DirectTypeOf(b[0]).String(), "meta_test.Foo")

	r := &Foo{}
	c := meta.SliceOf(r).Interface().([]*Foo)
	assert.Equal(t, len(c), 1)
	assert.Equal(t, meta.DirectTypeOf(c[0]).String(), "*meta_test.Foo")
	assert.Equal(t, c[0], r)

	var e *Foo
	d := meta.SliceOf(e).Interface().([]*Foo)
	assert.Equal(t, len(d), 1)
	assert.Equal(t, meta.DirectTypeOf(d[0]).String(), "*meta_test.Foo")
	assert.Equal(t, d[0], e)
}
