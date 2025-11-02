package crud

import (
	"reflect"

	"github.com/azer/crud/v3/meta"
	"github.com/azer/crud/v3/pg"
	"github.com/azer/crud/v3/types"
	"github.com/azer/snakecase"
)

// Take any kind of struct and return a FieldIteration instance
// which helps walking the fields of the given struct one by one
// reading its name, value and SQL options
func NewFieldIteration(driver string, st interface{}) *FieldIteration {
	rvalue, rtype := meta.Get(st)

	length := rvalue.NumField()

	return &FieldIteration{
		Driver:       driver,
		Index:        -1,
		Length:       length,
		ReflectType:  rtype,
		ReflectValue: rvalue,
	}
}

type FieldIteration struct {
	Driver       string
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

func (iteration *FieldIteration) SQLOptions() (*types.ColumnOptions, error) {
	result, err := types.NewColumnOptions(iteration.TypeField().Tag.Get("sql"))
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
		goTypeName := iteration.TypeField().Type.String()

		sqlType, sqlLength, err := types.MatchType(iteration.Driver, goTypeName)
		if err != nil {
			return nil, err
		}

		if pg.IsPostgres(iteration.Driver) && (result.Serial || result.IsAutoIncrementing) {
			sqlType = "SERIAL"
			sqlLength = -1
		}

		result.Type = sqlType
		result.Length = sqlLength
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
	if _, _, err := types.MatchType(iteration.Driver, iteration.TypeField().Type.String()); err == nil {
		return false
	}

	/*if _, ok := sql.TypeDict[iteration.TypeField().Type.String()]; ok {
		return false
	}*/

	return iteration.ReflectValue.Field(iteration.Index).Kind() == reflect.Struct
}
