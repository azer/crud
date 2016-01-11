package crud_test

import (
	"github.com/azer/crud"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFields(t *testing.T) {
	fields, err := crud.GetFieldsOf(UserProfile{})
	assert.Nil(t, err)
	assert.Equal(t, len(fields), 4)
	assert.Equal(t, fields[0].Name, "Id")
	assert.Equal(t, fields[0].SQL.Name, "id")
	assert.Equal(t, fields[0].SQL.Type, "int")
	assert.Equal(t, fields[0].SQL.AutoIncrement, 1)
	assert.True(t, fields[0].SQL.IsRequired)
	assert.True(t, fields[0].SQL.IsPrimaryKey)
	assert.Equal(t, fields[1].Name, "Name")
	assert.Equal(t, fields[1].SQL.Name, "name")
	assert.Equal(t, fields[1].SQL.Type, "varchar")
	assert.Equal(t, fields[2].Name, "Bio")
	assert.Equal(t, fields[2].SQL.Name, "bio")
	assert.Equal(t, fields[2].SQL.Type, "text")
	assert.Equal(t, fields[3].Name, "Email")
	assert.Equal(t, fields[3].SQL.Name, "email")
	assert.Equal(t, fields[3].SQL.Type, "varchar")
}

func TestHasPK(t *testing.T) {
	fields, err := crud.GetFieldsOf(UserProfile{})
	assert.Nil(t, err)
	assert.True(t, crud.HasPK(fields))

	fields, err = crud.GetFieldsOf(Post{})
	assert.Nil(t, err)
	assert.True(t, crud.HasPK(fields))

	fields, err = crud.GetFieldsOf(Foo{})
	assert.Nil(t, err)
	assert.False(t, crud.HasPK(fields))
}

func TestSetDefaultPK(t *testing.T) {
	fields, err := crud.GetFieldsOf(Foo{})
	crud.SetDefaultPK(fields)
	assert.Nil(t, err)
	assert.True(t, fields[0].SQL.IsPrimaryKey)
}

func TestDefaultPK(t *testing.T) {
	fields, err := crud.GetFieldsOf(Foo{})
	assert.Nil(t, err)

	crud.SetDefaultPK(fields)
	assert.Equal(t, len(fields), 4)
	assert.Equal(t, fields[0].Name, "Id")
	assert.Equal(t, fields[0].SQL.Name, "id")
	assert.Equal(t, fields[0].SQL.Type, "int")
	assert.Equal(t, fields[0].SQL.AutoIncrement, 1)
	assert.True(t, fields[0].SQL.IsRequired)
	assert.True(t, fields[0].SQL.IsPrimaryKey)
	assert.Equal(t, fields[1].Name, "APIKey")
	assert.Equal(t, fields[1].SQL.Name, "api_key")
	assert.Equal(t, fields[1].SQL.Type, "varchar")
	assert.Equal(t, fields[2].Name, "YOLO")
	assert.Equal(t, fields[2].SQL.Name, "yolo")
	assert.Equal(t, fields[2].SQL.Type, "tinyint")
	assert.Equal(t, fields[3].Name, "Beast")
	assert.Equal(t, fields[3].SQL.Name, "beast")
	assert.Equal(t, fields[3].SQL.Type, "varchar")
}
