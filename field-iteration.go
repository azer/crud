package crud

import (
	"reflect"

	"github.com/azer/crud/v2/meta"
	"github.com/azer/crud/v2/sql"
	"github.com/azer/snakecase"
)

// Take any kind of struct and return a FieldIteration instance
// which helps walking the fields of the given struct one by one
// reading its name, value and SQL options
func NewFieldIteration(st interface{}) *FieldIteration {
	rvalue, rtype := meta.Get(st)

	length := rvalue.NumField()

	return &FieldIteration{
		Index:        -1,
		Length:       length,
		ReflectType:  rtype,
		ReflectValue: rvalue,
	}
}

type FieldIteration struct {
	Index        int
	Length       int
	ReflectValue reflect.Value
	ReflectType  reflect.Type
}

func (iteration *FieldIteration) Next() bool {
	if iteration.Index+1 >= iteration.Length {
		return false
	}

	iteration.Index += 1
	return true
}

func (iteration *FieldIteration) TypeField() reflect.StructField {
	return iteration.ReflectType.Field(iteration.Index)
}

func (iteration *FieldIteration) ValueField() reflect.Value {
	return iteration.ReflectValue.Field(iteration.Index)
}

func (iteration *FieldIteration) SQLOptions() (*sql.Options, error) {
	result, err := sql.NewOptions(iteration.TypeField().Tag.Get("sql"))
	if err != nil {
		return nil, err
	}

	if result.Ignore {
		return result, nil
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

func (iteration *FieldIteration) Value() interface{} {
	return iteration.ReflectValue.Field(iteration.Index).Interface()
}

func (iteration *FieldIteration) Name() string {
	return iteration.TypeField().Name
}

func (iteration *FieldIteration) IsEmbeddedStruct() bool {
	if _, ok := sql.TypeDict[iteration.TypeField().Type.String()]; ok {
		return false
	}

	return iteration.ReflectValue.Field(iteration.Index).Kind() == reflect.Struct
}
