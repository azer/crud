package crud_test

import (
	"testing"

	"github.com/azer/crud/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewTable(t *testing.T) {
	table, err := crud.NewTable(UserProfile{})
	assert.Nil(t, err)
	assert.Equal(t, table.Name, "UserProfile")
	assert.Equal(t, table.SQLName, "user_profiles")
	assert.Equal(t, len(table.Fields), 6)
	assert.Equal(t, table.Fields[0].Name, "Id")
	assert.Equal(t, table.Fields[0].SQL.Name, "id")
	assert.Equal(t, table.Fields[0].SQL.Type, "int")
	assert.Equal(t, table.Fields[0].SQL.Length, 11)
	assert.Equal(t, table.Fields[0].SQL.AutoIncrement, 1)
	assert.Equal(t, table.Fields[0].SQL.IsPrimaryKey, true)
	assert.Equal(t, table.Fields[0].SQL.IsRequired, true)
	assert.Equal(t, table.Fields[1].Name, "Name")
	assert.Equal(t, table.Fields[1].SQL.Name, "name")
	assert.Equal(t, table.Fields[1].SQL.Type, "varchar")
	assert.Equal(t, table.Fields[1].SQL.Length, 255)
	assert.Equal(t, table.Fields[1].SQL.IsUnique, false)
	assert.Equal(t, table.Fields[1].SQL.IsRequired, true)
	assert.Equal(t, table.Fields[2].Name, "Bio")
	assert.Equal(t, table.Fields[2].SQL.Name, "bio")
	assert.Equal(t, table.Fields[2].SQL.Type, "text")
	assert.Equal(t, table.Fields[2].SQL.Length, -1)
	assert.Equal(t, table.Fields[3].Name, "Email")
	assert.Equal(t, table.Fields[3].SQL.Name, "email")
	assert.Equal(t, table.Fields[3].SQL.Type, "varchar")
	assert.Equal(t, table.Fields[3].SQL.Length, 255)
	assert.Equal(t, table.Fields[4].SQL.Name, "attachment")
	assert.Equal(t, table.Fields[4].SQL.Type, "blob")
	assert.Equal(t, table.Fields[5].Name, "Modified")
	assert.Equal(t, table.Fields[5].SQL.Name, "modified_col")
	assert.Equal(t, table.Fields[5].SQL.Type, "bigint")
	assert.Equal(t, table.Fields[5].SQL.Length, 20)
}

func TestColumnDict(t *testing.T) {
	table, err := crud.NewTable(Foo{})
	assert.Nil(t, err)

	columnDict := table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")

	table, err = crud.NewTable(&Foo{})
	assert.Nil(t, err)

	columnDict = table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")

	/*var f *Foo
	table, err = crud.NewTable(f)
	assert.Nil(t, err)

	columnDict = table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")*/
}

func TestColumnDictOfSlices(t *testing.T) {
	table, err := crud.NewTable([]Foo{})
	assert.Nil(t, err)

	columnDict := table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")

	table, err = crud.NewTable(FooSlice{})
	assert.Nil(t, err)

	columnDict = table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")

	table, err = crud.NewTable(&FooSlice{})
	assert.Nil(t, err)

	columnDict = table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")

	table, err = crud.NewTable(FooPTRSlice{})
	assert.Nil(t, err)

	columnDict = table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")

	table, err = crud.NewTable(&FooPTRSlice{})
	assert.Nil(t, err)

	columnDict = table.SQLColumnDict()
	assert.Equal(t, columnDict["id"], "Id")
	assert.Equal(t, columnDict["api_key"], "APIKey")
	assert.Equal(t, columnDict["yolo"], "YOLO")
	assert.Equal(t, columnDict["beast"], "Beast")
}

func TestReadingCustomTableName(t *testing.T) {
	table, err := crud.NewTable(CustomTableName{})
	assert.Nil(t, err)
	assert.Equal(t, table.Name, "CustomTableName")
	assert.Equal(t, table.SQLName, "yolo")
}

func TestReadingCustomTableNameFromList(t *testing.T) {
	table, err := crud.NewTable([]*CustomTableName{})
	assert.Nil(t, err)
	assert.Equal(t, table.Name, "CustomTableName")
	assert.Equal(t, table.SQLName, "yolo")
}

func TestReadTableName(t *testing.T) {
	// Pointer to slice
	name, sqlName := crud.ReadTableName(&[]*CustomTableName{})
	assert.Equal(t, name, "CustomTableName")
	assert.Equal(t, sqlName, "yolo")

	// Slice
	name, sqlName = crud.ReadTableName([]*CustomTableName{})
	assert.Equal(t, name, "CustomTableName")
	assert.Equal(t, sqlName, "yolo")

	// Pointer
	name, sqlName = crud.ReadTableName(&CustomTableName{})
	assert.Equal(t, name, "CustomTableName")
	assert.Equal(t, sqlName, "yolo")

	// Struct
	name, sqlName = crud.ReadTableName(CustomTableName{})
	assert.Equal(t, name, "CustomTableName")
	assert.Equal(t, sqlName, "yolo")
}

func TestMixed(t *testing.T) {
	table, err := crud.NewTable(Mixed{})
	assert.Nil(t, err)
	assert.Equal(t, table.Name, "Mixed")
	assert.Equal(t, table.SQLName, "__mixed__")
}

func TestReadingTableColumns(t *testing.T) {
	columns, err := crud.ReadTableColumns(UserProfile{})

	assert.Nil(t, err)
	assert.Equal(t, columns[0], "id")
	assert.Equal(t, columns[1], "name")
	assert.Equal(t, columns[2], "bio")
	assert.Equal(t, columns[3], "email")
	assert.Equal(t, columns[4], "attachment")
	assert.Equal(t, columns[5], "modified_col")
}

func TestReadingTableColumnsFromPointer(t *testing.T) {
	columns, err := crud.ReadTableColumns(&UserProfile{})

	assert.Nil(t, err)
	assert.Equal(t, columns[0], "id")
	assert.Equal(t, columns[1], "name")
	assert.Equal(t, columns[2], "bio")
	assert.Equal(t, columns[3], "email")
	assert.Equal(t, columns[4], "attachment")
	assert.Equal(t, columns[5], "modified_col")
}

func TestReadingTableColumnsFromList(t *testing.T) {
	list := []*UserProfile{}

	columns, err := crud.ReadTableColumns(&list)

	assert.Nil(t, err)
	assert.Equal(t, columns[0], "id")
	assert.Equal(t, columns[1], "name")
	assert.Equal(t, columns[2], "bio")
	assert.Equal(t, columns[3], "email")
	assert.Equal(t, columns[4], "attachment")
	assert.Equal(t, columns[5], "modified_col")
}
