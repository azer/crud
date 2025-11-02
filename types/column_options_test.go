package types_test

import (
	"testing"

	"github.com/azer/crud/v3/types"
	"github.com/stretchr/testify/assert"
)

func TestAllOptions(t *testing.T) {
	o, err := types.NewColumnOptions("varchar(255) auto-increment=25 primary-key required default='yolo' name=id table-name='hello-kitty'")
	assert.Nil(t, err)
	assert.Equal(t, o.Name, "id")
	assert.Equal(t, o.Type, "varchar")
	assert.Equal(t, o.Length, 255)
	assert.Equal(t, o.AutoIncrement, 25)
	assert.True(t, o.IsAutoIncrementing)
	assert.Equal(t, o.IsPrimaryKey, true)
	assert.Equal(t, o.IsRequired, true)
	assert.Equal(t, o.DefaultValue, "'yolo'")
	assert.Equal(t, o.TableName, "'hello-kitty'")
}

func TestIgnoring(t *testing.T) {
	o, err := types.NewColumnOptions("-")
	assert.Nil(t, err)
	assert.Equal(t, o.Ignore, true)
}

func TestDefaultValues(t *testing.T) {
	o, err := types.NewColumnOptions("int autoincrement")
	assert.Nil(t, err)
	assert.Equal(t, o.AutoIncrement, 1)
	assert.True(t, o.IsAutoIncrementing)
	assert.Equal(t, o.Type, "int")
	assert.Equal(t, o.Length, 11)
}

func TestCustomTypes(t *testing.T) {
	o, err := types.NewColumnOptions("text")
	assert.Nil(t, err)
	assert.Equal(t, o.Type, "text")
	assert.False(t, o.IsAutoIncrementing)

	o, err = types.NewColumnOptions("auto_increment type=bigint(16)")
	assert.Nil(t, err)
	assert.Equal(t, o.Type, "bigint")
	assert.Equal(t, o.Length, 16)
	assert.Equal(t, o.AutoIncrement, 1)
	assert.True(t, o.IsAutoIncrementing)
}

func TestFailing(t *testing.T) {
	_, err := types.NewColumnOptions("yolo")
	assert.NotNil(t, err)
}
