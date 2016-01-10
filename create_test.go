package crud_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	DB.ResetTables(UserProfile{})

	azer := UserProfile{
		Name:  "Azer",
		Bio:   "I like photography",
		Email: "azer@roadbeats.com",
	}

	err := DB.Create(azer)
	assert.Nil(t, err)

	DB.DropTables(UserProfile{})
}
