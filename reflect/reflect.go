package reflect

import (
	"fmt"
	lib "reflect"
	"strings"
)

func Get(any interface{}) (lib.Value, lib.Type) {
	v := lib.Indirect(lib.ValueOf(any))
	return v, v.Type()
}

func TypeNameOf(any interface{}) string {
	parts := strings.Split(TypeOf(any).String(), ".")
	return parts[len(parts)-1]
}

func TypeOf(any interface{}) lib.Type {
	return ValueOf(any).Type()
}

func ValueOf(any interface{}) lib.Value {
	return lib.Indirect(lib.ValueOf(any))
}

func DirectTypeOf(any interface{}) lib.Type {
	return lib.TypeOf(any)
}

func DirectValueOf(any interface{}) lib.Value {
	return lib.ValueOf(any)
}

func IsPointer(any interface{}) bool {
	return DirectValueOf(any).Kind() == lib.Ptr
}

func HasPointers(any interface{}) bool {
	if !IsSlice(any) {
		return IsPointer(any)
	}

	return ElementType(any).Kind() == lib.Ptr
}

func CreateIfNil(any interface{}) interface{} {
	val := DirectValueOf(any)

	fmt.Println("?", val)

	if !val.IsNil() {
		return any
	}

	return val.Interface()
}

func HasAnyStruct(any interface{}) bool {
	var t lib.Type

	if IsSlice(any) {
		t = ElementType(any)
	} else {
		t = TypeOf(any)
	}

	if t.Kind() == lib.Ptr {
		t = t.Elem()
	}

	return t.Kind() == lib.Struct
}
