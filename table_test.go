package crud_test

import (
	"github.com/azer/crud"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTable(t *testing.T) {
	table, err := crud.NewTable(UserProfile{})
	assert.Nil(t, err)
	assert.Equal(t, table.Name, "UserProfile")
	assert.Equal(t, table.SQLName, "user_profile")
	assert.Equal(t, len(table.Fields), 4)
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
	assert.Equal(t, table.Fields[1].SQL.IsUnique, true)
	assert.Equal(t, table.Fields[1].SQL.IsRequired, true)
	assert.Equal(t, table.Fields[2].Name, "Bio")
	assert.Equal(t, table.Fields[2].SQL.Name, "bio")
	assert.Equal(t, table.Fields[2].SQL.Type, "text")
	assert.Equal(t, table.Fields[2].SQL.Length, -1)
	assert.Equal(t, table.Fields[3].Name, "Email")
	assert.Equal(t, table.Fields[3].SQL.Name, "email")
	assert.Equal(t, table.Fields[3].SQL.Type, "varchar")
	assert.Equal(t, table.Fields[3].SQL.Length, 255)
}

func TestCreate(t *testing.T) {
	DB.CreateTable(UserProfile{})

	azer := UserProfile{
		Name:  "Azer",
		Bio:   "I like photography",
		Email: "azer@roadbeats.com",
	}

	err := DB.Create(azer)
	assert.Nil(t, err)

	DB.DropTable(UserProfile{})
}
