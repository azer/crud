package crud

import (
	"github.com/azer/crud/meta"
	"github.com/azer/snakecase"
)

func SQLTableNameOf(st interface{}) string {
	return snakecase.SnakeCase(meta.TypeNameOf(st))
}
