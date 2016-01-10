package crud

import (
	"github.com/azer/crud/reflect"
	"github.com/azer/snakecase"
)

func SQLTableNameOf(st interface{}) string {
	return snakecase.SnakeCase(reflect.TypeNameOf(st))
}
