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
