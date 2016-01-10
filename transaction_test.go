package crud_test

/*
FIXME Transaction tests are not complete yet

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestSuccessfulCommit(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	tx, err := DB.Begin()
	assert.Nil(t, err)

	n := UserProfile{}
	err = DB.Read(&n, "WHERE id = ?", 2)
	assert.Nil(t, err)

	n.Bio = "let's go somewhere"

	assert.Nil(t, tx.Update(&n))

	azer := UserProfile{}
	err = DB.Read(&azer, "WHERE id = ?", 2)
	assert.Nil(t, err)
	assert.Equal(t, azer.Bio, "Engineer")

	assert.Nil(t, tx.Commit())

	time.Sleep(time.Second * 3)

	azerc := UserProfile{}
	err = DB.Read(&azerc, "WHERE id = ?", 2)
	assert.Nil(t, err)
	assert.Equal(t, azer.Bio, "let's go somewhere")
}
*/
