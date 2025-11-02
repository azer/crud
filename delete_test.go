package crud_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*func TestDelete(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	fuji := UserProfile{}
	err := DB.Read(&fuji, "SELECT * FROM user_profile WHERE name = 'Fuji'")
	assert.Nil(t, err)

	assert.Nil(t, DB.Delete(fuji))

	fujic := UserProfile{}
	err = DB.Read(&fujic, "SELECT * FROM user_profile WHERE name = 'Fuji'")
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

	fuji := UserProfile{}
	err := DB.Read(&fuji, "SELECT * FROM user_profiles WHERE name = 'Fuji'")
	assert.Nil(t, err)

	assert.Nil(t, DB.Delete(fuji))

	fujic := UserProfile{}
	err = DB.Read(&fujic, "SELECT * FROM user_profiles WHERE name = 'Fuji'")
	assert.NotNil(t, err)
}

func TestMustDeleteNotMatching(t *testing.T) {
	assert.NotNil(t, DB.Delete(&UserProfile{
		Id:   123,
		Name: "Yolo",
	}))
}
