//go:build postgres
// +build postgres

package crud_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresConnection(t *testing.T) {
	assert.Nil(t, DB.Ping())
}

func TestPostgresCreateTable(t *testing.T) {
	err := DB.CreateTables(UserProfile{})
	assert.Nil(t, err)
	assert.True(t, DB.CheckIfTableExists("user_profiles"))
	DB.DropTables(UserProfile{})
}

func TestPostgresInsertReturning(t *testing.T) {
	DB.ResetTables(UserProfile{})

	user := UserProfile{
		Name:  "Test",
		Email: "test@example.com",
		Bio:   "Testing postgres",
	}

	err := DB.CreateAndRead(&user)
	assert.Nil(t, err)
	assert.NotEqual(t, user.Id, 0)

	DB.DropTables(UserProfile{})
}

func TestPostgresSequence(t *testing.T) {
	DB.ResetTables(UserProfile{})

	user1 := UserProfile{Name: "User1", Email: "u1@test.com"}
	user2 := UserProfile{Name: "User2", Email: "u2@test.com"}

	DB.CreateAndRead(&user1)
	DB.CreateAndRead(&user2)

	assert.Equal(t, user1.Id, 1)
	assert.Equal(t, user2.Id, 2)

	DB.DropTables(UserProfile{})
}
