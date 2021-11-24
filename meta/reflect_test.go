package meta_test

import (
	"fmt"
	"testing"

	"github.com/azer/crud/v2/meta"
	"github.com/stretchr/testify/assert"
)

type Foo struct {
	Id int
}

type EmbeddedFoo struct {
	Foo
	Span bool
	Eggs int
}

type FooSlice []Foo
type FooPTRSlice []*Foo

func TestGetTypeNameOf(t *testing.T) {
	assert.Equal(t, meta.TypeNameOf(Foo{}), "Foo")
	assert.Equal(t, meta.TypeNameOf(FooSlice{}), "FooSlice")
	assert.Equal(t, meta.TypeNameOf(&FooPTRSlice{}), "FooPTRSlice")
}

func TestElementType(t *testing.T) {
	assert.Equal(t, meta.ElementType([]Foo{}).String(), "meta_test.Foo")
	assert.Equal(t, meta.ElementType([]*Foo{}).String(), "*meta_test.Foo")
	assert.Equal(t, meta.ElementType(&[]*Foo{}).String(), "*meta_test.Foo")
	assert.Equal(t, meta.ElementType(FooSlice{}).String(), "meta_test.Foo")
	assert.Equal(t, meta.ElementType(&FooSlice{}).String(), "meta_test.Foo")
	assert.Equal(t, meta.ElementType(FooPTRSlice{}).String(), "*meta_test.Foo")
	assert.Equal(t, meta.ElementType(&FooPTRSlice{}).String(), "*meta_test.Foo")
}

func TestCreateElement(t *testing.T) {
	assert.Equal(t, meta.DirectTypeOf(meta.CreateElement([]Foo{}).Interface()).String(), "meta_test.Foo")
	assert.Equal(t, meta.DirectTypeOf(meta.CreateElement([]*Foo{}).Interface()).String(), "*meta_test.Foo")
	assert.Equal(t, meta.DirectTypeOf(meta.CreateElement(&[]Foo{}).Interface()).String(), "meta_test.Foo")
	assert.Equal(t, meta.DirectTypeOf(meta.CreateElement(&[]*Foo{}).Interface()).String(), "*meta_test.Foo")
	assert.Equal(t, meta.DirectTypeOf(meta.CreateElement(FooPTRSlice{}).Interface()).String(), "*meta_test.Foo")
	assert.Equal(t, meta.DirectTypeOf(meta.CreateElement(&FooPTRSlice{}).Interface()).String(), "*meta_test.Foo")
	assert.Equal(t, meta.DirectTypeOf(meta.CreateElement([]string{}).Interface()).String(), "string")
}

func TestIsPointer(t *testing.T) {
	assert.True(t, meta.IsPointer(&Foo{}))
	assert.True(t, meta.IsPointer(&[]Foo{}))
	assert.False(t, meta.IsPointer(Foo{}))
	assert.False(t, meta.IsPointer(*&Foo{}))
}

func TestIsStruct(t *testing.T) {
	ef := EmbeddedFoo{}

	assert.True(t, meta.IsStruct(ef))
	assert.True(t, meta.IsStruct(ef.Foo))
	assert.False(t, meta.IsStruct(ef.Span))
	assert.False(t, meta.IsStruct(ef.Eggs))
}

func TestHasPointers(t *testing.T) {
	assert.True(t, meta.HasPointers([]*Foo{}))
	assert.True(t, meta.HasPointers(&[]*Foo{}))
	assert.False(t, meta.HasPointers([]Foo{}))
	assert.False(t, meta.HasPointers(&[]Foo{}))
	assert.True(t, meta.HasPointers(&Foo{}))
	assert.False(t, meta.HasPointers(Foo{}))
}

func TestHasAnyStruct(t *testing.T) {
	var u Foo
	var up Foo
	var sup []*Foo

	assert.True(t, meta.HasAnyStruct(Foo{}))
	assert.True(t, meta.HasAnyStruct(&Foo{}))
	assert.True(t, meta.HasAnyStruct([]Foo{}))
	assert.True(t, meta.HasAnyStruct([]*Foo{}))
	assert.True(t, meta.HasAnyStruct(&[]Foo{}))
	assert.True(t, meta.HasAnyStruct(&[]*Foo{}))
	assert.True(t, meta.HasAnyStruct(u))
	assert.True(t, meta.HasAnyStruct(&up))
	assert.True(t, meta.HasAnyStruct(sup))

	var a string
	var b int
	var c bool

	assert.False(t, meta.HasAnyStruct(123))
	assert.False(t, meta.HasAnyStruct(true))
	assert.False(t, meta.HasAnyStruct(""))
	assert.False(t, meta.HasAnyStruct(a))
	assert.False(t, meta.HasAnyStruct(b))
	assert.False(t, meta.HasAnyStruct(c))
}

/*func TestIsEmpty(t *testing.T) {
	var f Foo
	var fptr *Foo
	var sup []*Foo
	var s string
	var ls []string
	var n int

	assert.True(t, meta.IsEmpty(f))
	assert.True(t, meta.IsEmpty(fptr))
	assert.True(t, meta.IsEmpty(sup))
	assert.True(t, meta.IsEmpty(s))
	assert.True(t, meta.IsEmpty(ls))
	assert.True(t, meta.IsEmpty(n))
}*/

func TestCreateIfNil(t *testing.T) {
	var f *Foo
	var sf []*Foo

	fmt.Println(f)
	fmt.Println(sf)

	fmt.Println("#", meta.CreateIfNil(f))
	fmt.Println(meta.CreateIfNil(sf))
}
