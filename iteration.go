package crud

import (
	"github.com/azer/crud/reflect"
	"github.com/azer/crud/sql"
	"github.com/azer/snakecase"
	reflectlib "reflect"
)

func NewIteration(st interface{}) *Iteration {
	rvalue, rtype := reflect.Get(st)
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
	ReflectValue reflectlib.Value
	ReflectType  reflectlib.Type
}

func (iteration *Iteration) Next() bool {
	if iteration.Index+1 >= iteration.Length {
		return false
	}

	iteration.Index += 1
	return true
}

func (iteration *Iteration) TypeField() reflectlib.StructField {
	return iteration.ReflectType.Field(iteration.Index)
}

func (iteration *Iteration) ValueField() reflectlib.Value {
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
