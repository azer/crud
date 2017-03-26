package crud_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	err := DB.ResetTables(UserProfile{})
	assert.Nil(t, err)

	azer := UserProfile{
		Name:  "Azer",
		Bio:   "I like photography",
		Email: "azer@roadbeats.com",
	}

	err = DB.Create(azer)
	assert.Nil(t, err)

	DB.DropTables(UserProfile{})
}

func TestCreateAndRead(t *testing.T) {
	DB.ResetTables(UserProfile{})

	azer := UserProfile{
		Name:  "Azer",
		Bio:   "I like photography",
		Email: "azer@roadbeats.com",
	}

	assert.Equal(t, azer.Id, 0)
	err := DB.CreateAndRead(&azer)
	assert.Nil(t, err)
	assert.Equal(t, azer.Id, 1)

	DB.DropTables(UserProfile{})
}

func TestCreateEmpty(t *testing.T) {
	DB.ResetTables(UserProfile{})

	azer := UserProfile{
		Name: "Azer",
	}

	err := DB.Create(azer)
	assert.Nil(t, err)

	DB.DropTables(UserProfile{})
}
