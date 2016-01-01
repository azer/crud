package crud

import (
	"github.com/azer/crud/sql"
	"github.com/azer/snakecase"
	"reflect"
)

func NewIteration(st interface{}) *Iteration {
	rvalue, rtype := GetReflection(st)
	length := rvalue.NumField()

	return &Iteration{
		Index:        -1,
		Length:       length,
		ReflectType:  rtype,
		ReflectValue: rvalue,
	}
}

type Iteration struct {
	Index        int
	Length       int
	ReflectValue reflect.Value
	ReflectType  reflect.Type
}

func (iteration *Iteration) Next() bool {
	if iteration.Index+1 >= iteration.Length {
		return false
	}

	iteration.Index += 1
	return true
}

func (iteration *Iteration) TypeField() reflect.StructField {
	return iteration.ReflectType.Field(iteration.Index)
}

func (iteration *Iteration) ValueField() reflect.Value {
	return iteration.ReflectValue.Field(iteration.Index)
}

func (iteration *Iteration) SQLOptions() (*sql.Options, error) {
	result, err := sql.NewOptions(iteration.TypeField().Tag.Get("sql"))
	if err != nil {
		return nil, err
	}

	if len(result.Name) == 0 {
		result.Name = snakecase.SnakeCase(iteration.Name())
	}

	if len(result.Type) == 0 {
		sqlType, err := sql.MatchType(iteration.TypeField().Type.String())
		if err != nil {
			return nil, err
		}

		result.Type = sqlType
		result.Length = sql.Types[result.Type]
	}

	return result, nil
}

func (iteration *Iteration) Value() interface{} {
	return iteration.ReflectValue.Field(iteration.Index).Interface()
}

func (iteration *Iteration) Name() string {
	return iteration.TypeField().Name
}

func GetReflection(st interface{}) (v reflect.Value, t reflect.Type) {
	defer func() {
		if r := recover(); r != nil {
			v = reflect.ValueOf(st)
			t = reflect.TypeOf(st)
		}
	}()

	v = reflect.ValueOf(st).Elem()
	t = v.Type()
	return
}
