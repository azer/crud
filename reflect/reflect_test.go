package reflect_test

import (
	"fmt"
	"github.com/azer/crud/reflect"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Foo struct {
	Id int
}

type FooSlice []Foo
type FooPTRSlice []*Foo

func TestGetTypeNameOf(t *testing.T) {
	assert.Equal(t, reflect.TypeNameOf(Foo{}), "Foo")
	assert.Equal(t, reflect.TypeNameOf(FooSlice{}), "FooSlice")
	assert.Equal(t, reflect.TypeNameOf(&FooPTRSlice{}), "FooPTRSlice")
}

func TestElementType(t *testing.T) {
	assert.Equal(t, reflect.ElementType([]Foo{}).String(), "reflect_test.Foo")
	assert.Equal(t, reflect.ElementType([]*Foo{}).String(), "*reflect_test.Foo")
	assert.Equal(t, reflect.ElementType(&[]*Foo{}).String(), "*reflect_test.Foo")
	assert.Equal(t, reflect.ElementType(FooSlice{}).String(), "reflect_test.Foo")
	assert.Equal(t, reflect.ElementType(&FooSlice{}).String(), "reflect_test.Foo")
	assert.Equal(t, reflect.ElementType(FooPTRSlice{}).String(), "*reflect_test.Foo")
	assert.Equal(t, reflect.ElementType(&FooPTRSlice{}).String(), "*reflect_test.Foo")
}

func TestCreateElement(t *testing.T) {
	assert.Equal(t, reflect.DirectTypeOf(reflect.CreateElement([]Foo{}).Interface()).String(), "reflect_test.Foo")
	assert.Equal(t, reflect.DirectTypeOf(reflect.CreateElement([]*Foo{}).Interface()).String(), "*reflect_test.Foo")
	assert.Equal(t, reflect.DirectTypeOf(reflect.CreateElement(&[]Foo{}).Interface()).String(), "reflect_test.Foo")
	assert.Equal(t, reflect.DirectTypeOf(reflect.CreateElement(&[]*Foo{}).Interface()).String(), "*reflect_test.Foo")
	assert.Equal(t, reflect.DirectTypeOf(reflect.CreateElement(FooPTRSlice{}).Interface()).String(), "*reflect_test.Foo")
	assert.Equal(t, reflect.DirectTypeOf(reflect.CreateElement(&FooPTRSlice{}).Interface()).String(), "*reflect_test.Foo")
	assert.Equal(t, reflect.DirectTypeOf(reflect.CreateElement([]string{}).Interface()).String(), "string")
}

func TestIsPointer(t *testing.T) {
	assert.True(t, reflect.IsPointer(&Foo{}))
	assert.True(t, reflect.IsPointer(&[]Foo{}))
	assert.False(t, reflect.IsPointer(Foo{}))
	assert.False(t, reflect.IsPointer(*&Foo{}))
}

func TestHasPointers(t *testing.T) {
	assert.True(t, reflect.HasPointers([]*Foo{}))
	assert.True(t, reflect.HasPointers(&[]*Foo{}))
	assert.False(t, reflect.HasPointers([]Foo{}))
	assert.False(t, reflect.HasPointers(&[]Foo{}))
	assert.True(t, reflect.HasPointers(&Foo{}))
	assert.False(t, reflect.HasPointers(Foo{}))
}

func TestHasAnyStruct(t *testing.T) {
	var u Foo
	//var up *Foo
	var sup []*Foo

	assert.True(t, reflect.HasAnyStruct(Foo{}))
	assert.True(t, reflect.HasAnyStruct(&Foo{}))
	assert.True(t, reflect.HasAnyStruct([]Foo{}))
	assert.True(t, reflect.HasAnyStruct([]*Foo{}))
	assert.True(t, reflect.HasAnyStruct(&[]Foo{}))
	assert.True(t, reflect.HasAnyStruct(&[]*Foo{}))
	assert.True(t, reflect.HasAnyStruct(u))
	//assert.True(t, reflect.HasAnyStruct(up))
	assert.True(t, reflect.HasAnyStruct(sup))

	var a string
	var b int
	var c bool

	assert.False(t, reflect.HasAnyStruct(123))
	assert.False(t, reflect.HasAnyStruct(true))
	assert.False(t, reflect.HasAnyStruct(""))
	assert.False(t, reflect.HasAnyStruct(a))
	assert.False(t, reflect.HasAnyStruct(b))
	assert.False(t, reflect.HasAnyStruct(c))
}

func TestCreateIfNil(t *testing.T) {
	var f *Foo
	var sf []*Foo

	fmt.Println(f)
	fmt.Println(sf)

	fmt.Println("#", reflect.CreateIfNil(f))
	fmt.Println(reflect.CreateIfNil(sf))
}
