package meta

import (
	"reflect"
	"strings"
)

func Get(any interface{}) (reflect.Value, reflect.Type) {
	v := reflect.Indirect(reflect.ValueOf(any))
	return v, v.Type()
}

func TypeNameOf(any interface{}) string {
	parts := strings.Split(TypeOf(any).String(), ".")
	return parts[len(parts)-1]
}

func TypeOf(any interface{}) reflect.Type {
	return ValueOf(any).Type()
}

func ValueOf(any interface{}) reflect.Value {
	return reflect.Indirect(reflect.ValueOf(any))
}

func DirectTypeOf(any interface{}) reflect.Type {
	return reflect.TypeOf(any)
}

func DirectValueOf(any interface{}) reflect.Value {
	return reflect.ValueOf(any)
}

func IsPointer(any interface{}) bool {
	return DirectValueOf(any).Kind() == reflect.Ptr
}

func IsStruct(any interface{}) bool {
	return DirectValueOf(any).Kind() == reflect.Struct
}

func HasPointers(any interface{}) bool {
	if !IsSlice(any) {
		return IsPointer(any)
	}

	return ElementType(any).Kind() == reflect.Ptr
}

func IsEmpty(any interface{}) bool {
	v := ValueOf(any)

	// I copy pasted this following switch-case code from encoding/json package
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return false
}

func CreateIfNil(any interface{}) interface{} {
	val := DirectValueOf(any)

	if !val.IsNil() {
		return any
	}

	return val.Interface()
}

func HasAnyStruct(any interface{}) bool {
	var t reflect.Type

	if IsSlice(any) {
		t = ElementType(any)
	} else {
		t = TypeOf(any)
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Kind() == reflect.Struct
}
