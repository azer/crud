package meta

import (
	"reflect"
)

func IsSlice(any interface{}) bool {
	return ValueOf(any).Kind() == reflect.Slice
}

func ElementType(any interface{}) reflect.Type {
	return TypeOf(any).Elem()
}

func Push(sl reflect.Value, el reflect.Value) {
	if DirectValueOf(el).Kind() == reflect.Ptr {
		el = el.Addr()
	}

	sl.Set(reflect.Append(sl, el))
}

func CreateElement(any interface{}) reflect.Value {
	t := ElementType(any)

	if t.Kind() != reflect.Ptr {
		return reflect.New(ElementType(any)).Elem()
	}

	return reflect.New(ElementType(any).Elem()).Elem().Addr()
}

func Addressable(any interface{}) reflect.Value {
	return DirectValueOf(any).Elem()
}

func SliceOf(any interface{}) reflect.Value {
	val := reflect.New(reflect.SliceOf(DirectTypeOf(any))).Elem()
	Push(val, DirectValueOf(any))
	return val
}
