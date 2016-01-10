package reflect

import (
	lib "reflect"
)

func IsSlice(any interface{}) bool {
	return ValueOf(any).Kind() == lib.Slice
}

func ElementType(any interface{}) lib.Type {
	return TypeOf(any).Elem()
}

func Push(sl lib.Value, el lib.Value) {
	if DirectValueOf(el).Kind() == lib.Ptr {
		el = el.Addr()
	}

	sl.Set(lib.Append(sl, el))
}

func CreateElement(any interface{}) lib.Value {
	t := ElementType(any)

	if t.Kind() != lib.Ptr {
		return lib.New(ElementType(any)).Elem()
	}

	return lib.New(ElementType(any).Elem()).Elem().Addr()
}

func Addressable(any interface{}) lib.Value {
	return DirectValueOf(any).Elem()
}

func SliceOf(any interface{}) lib.Value {
	val := lib.New(lib.SliceOf(DirectTypeOf(any))).Elem()
	Push(val, DirectValueOf(any))
	return val
}
