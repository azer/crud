package crud_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustUpdate(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	fuji := UserProfile{}
	err := DB.Read(&fuji, "SELECT * FROM user_profiles WHERE name = 'Fuji'")
	assert.Nil(t, err)

	fuji.Bio = "Hola"
	assert.Nil(t, DB.Update(fuji))

	fujic := UserProfile{}
	err = DB.Read(&fujic, "SELECT * FROM user_profiles WHERE name = 'Fuji'")
	assert.Nil(t, err)
	assert.Equal(t, fujic.Bio, "Hola")
	assert.Equal(t, fujic.Email, fuji.Email)
	assert.Equal(t, fujic.Id, fuji.Id)
}

func TestMustUpdateNotMatching(t *testing.T) {
	assert.NotNil(t, DB.Update(&UserProfile{
		Id:   123,
		Name: "Yolo",
	}))
}

/*func TestUpdate(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	fuji := UserProfile{}
	err := DB.Read(&fuji, "SELECT * FROM user_profile WHERE name = 'Fuji'")
	assert.Nil(t, err)

	fuji.Bio = "Y O L O"
	assert.Nil(t, DB.Update(fuji))

	fujic := UserProfile{}
	err = DB.Read(&fujic, "SELECT * FROM user_profile WHERE name = 'Fuji'")
	assert.Nil(t, err)
	assert.Equal(t, fujic.Bio, "Y O L O")
	assert.Equal(t, fujic.Email, fuji.Email)
	assert.Equal(t, fujic.Id, fuji.Id)
}

func TestUpdateNotMatching(t *testing.T) {
	assert.Nil(t, DB.Update(&UserProfile{
		Id:   123,
		Name: "Yolo",
	}))
}*/
