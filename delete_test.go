package crud_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*func TestDelete(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	nova := UserProfile{}
	err := DB.Read(&nova, "SELECT * FROM user_profile WHERE name = 'Nova'")
	assert.Nil(t, err)

	assert.Nil(t, DB.Delete(nova))

	novac := UserProfile{}
	err = DB.Read(&novac, "SELECT * FROM user_profile WHERE name = 'Nova'")
	assert.NotNil(t, err)
}

func TestDeleteNotMatching(t *testing.T) {
	assert.Nil(t, DB.Delete(&UserProfile{
		Id:   123,
		Name: "Yolo",
	}))
}*/

func TestMustDelete(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	nova := UserProfile{}
	err := DB.Read(&nova, "SELECT * FROM user_profiles WHERE name = 'Nova'")
	assert.Nil(t, err)

	assert.Nil(t, DB.Delete(nova))

	novac := UserProfile{}
	err = DB.Read(&novac, "SELECT * FROM user_profiles WHERE name = 'Nova'")
	assert.NotNil(t, err)
}

func TestMustDeleteNotMatching(t *testing.T) {
	assert.NotNil(t, DB.Delete(&UserProfile{
		Id:   123,
		Name: "Yolo",
	}))
}
