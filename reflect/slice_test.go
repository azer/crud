package reflect_test

import (
	"github.com/azer/crud/reflect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSlice(t *testing.T) {
	assert.False(t, reflect.IsSlice(123))
	assert.False(t, reflect.IsSlice("yo"))
	assert.False(t, reflect.IsSlice(Foo{}))
	assert.True(t, reflect.IsSlice(FooSlice{}))
	assert.True(t, reflect.IsSlice(&FooSlice{}))
	assert.True(t, reflect.IsSlice(FooPTRSlice{}))
	assert.True(t, reflect.IsSlice(&FooPTRSlice{}))
}

func TestPushToSlice(t *testing.T) {
	sl := []Foo{}
	slval := reflect.Addressable(&sl)
	reflect.Push(slval, reflect.ValueOf(Foo{123}))
	reflect.Push(slval, reflect.ValueOf(Foo{456}))

	assert.Equal(t, len(sl), 2)
	assert.Equal(t, sl[0].Id, 123)
	assert.Equal(t, sl[1].Id, 456)
}

func TestPushToPTRSlice(t *testing.T) {
	sl := []*Foo{}
	slval := reflect.Addressable(&sl)
	reflect.Push(slval, reflect.DirectValueOf(&Foo{123}))
	reflect.Push(slval, reflect.DirectValueOf(&Foo{456}))

	assert.Equal(t, len(sl), 2)
	assert.Equal(t, sl[0].Id, 123)
	assert.Equal(t, sl[1].Id, 456)
}

func TestSliceOf(t *testing.T) {
	a := reflect.SliceOf(123).Interface().([]int)
	assert.Equal(t, len(a), 1)
	assert.Equal(t, a[0], 123)

	b := reflect.SliceOf(Foo{}).Interface().([]Foo)
	assert.Equal(t, len(b), 1)
	assert.Equal(t, reflect.DirectTypeOf(b[0]).String(), "reflect_test.Foo")

	r := &Foo{}
	c := reflect.SliceOf(r).Interface().([]*Foo)
	assert.Equal(t, len(c), 1)
	assert.Equal(t, reflect.DirectTypeOf(c[0]).String(), "*reflect_test.Foo")
	assert.Equal(t, c[0], r)

	var e *Foo
	d := reflect.SliceOf(e).Interface().([]*Foo)
	assert.Equal(t, len(d), 1)
	assert.Equal(t, reflect.DirectTypeOf(d[0]).String(), "*reflect_test.Foo")
	assert.Equal(t, d[0], e)
}
