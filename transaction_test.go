package crud_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulCommit(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	tx, err := DB.Begin(context.Background())
	assert.Nil(t, err)

	n := UserProfile{}
	err = tx.Read(&n, "SELECT * from user_profiles WHERE id = ?", 2)
	assert.Nil(t, err)

	n.Bio = "let's go somewhere"

	assert.Nil(t, tx.Update(&n))

	azer := UserProfile{}
	err = DB.Read(&azer, "SELECT * from user_profiles WHERE id = ?", 2)
	assert.Nil(t, err)
	assert.Equal(t, azer.Bio, "Engineer")

	assert.Nil(t, tx.Commit())

	time.Sleep(time.Second * 1)

	azerc := UserProfile{}
	err = DB.Read(&azerc, "SELECT * from user_profiles WHERE id = ?", 2)
	assert.Nil(t, err)
	assert.Equal(t, azerc.Bio, "let's go somewhere")

	DB.DropTables(UserProfile{})
}

func TestRollback(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	tx, err := DB.Begin(context.Background())
	assert.Nil(t, err)

	err = tx.Create(&UserProfile{
		Email: "row1@rows.com",
		Name:  "Row1",
		Bio:   "testing transactions",
	})

	assert.Nil(t, err)

	err = tx.Create(&UserProfile{
		Email: "row2@rows.com",
		Name:  "Row2",
		Bio:   "testing transactions",
	})

	assert.Nil(t, err)

	err = tx.Create(&UserProfile{
		Email: "row1@rows.com",
		Name:  "Row3",
		Bio:   "testing transactions, should fail",
	})

	assert.Error(t, err)
	assert.Nil(t, tx.Rollback())

	shouldNotExist := UserProfile{}
	err = DB.Read(&shouldNotExist, "SELECT * from user_profiles WHERE email = ?", "row1@rows.com")
	assert.Error(t, err)
	assert.True(t, err == sql.ErrNoRows)

	DB.DropTables(UserProfile{})
}
