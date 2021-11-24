package crud_test

import (
	"testing"

	"github.com/azer/crud/v2"
	"github.com/stretchr/testify/assert"
)

func TestResolveReadParams(t *testing.T) {
	query, params, err := crud.ResolveReadParams([]interface{}{})
	assert.Equal(t, query, "")
	assert.Equal(t, len(params), 0)
	assert.Nil(t, err)

	query, params, err = crud.ResolveReadParams([]interface{}{123, 456})
	assert.Equal(t, query, "")
	assert.Nil(t, params)
	assert.NotNil(t, err)

	query, params, err = crud.ResolveReadParams([]interface{}{"yolo", 456, 123})
	assert.Equal(t, query, "yolo")
	assert.Equal(t, len(params), 2)
	assert.Nil(t, err)
}

func TestReadingMultipleRows(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	result := []UserProfile{}
	err := DB.Read(&result, "SELECT * FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0].Name, "Nova")
	assert.Equal(t, result[0].Bio, "Photographer")
	assert.Equal(t, result[0].Email, "nova@roadbeats.com")
	assert.Equal(t, result[1].Name, "Azer")
	assert.Equal(t, result[1].Bio, "Engineer")
	assert.Equal(t, result[1].Email, "azer@roadbeats.com")
	assert.Equal(t, string(result[1].Attachment), "{ \"azer\": \"bar\" }")

	resultptr := []*UserProfile{}
	err = DB.Read(&resultptr, "SELECT * FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, len(resultptr), 3)
	assert.Equal(t, resultptr[0].Name, "Nova")
	assert.Equal(t, resultptr[0].Bio, "Photographer")
	assert.Equal(t, resultptr[0].Email, "nova@roadbeats.com")
	assert.Equal(t, resultptr[1].Name, "Azer")
	assert.Equal(t, resultptr[1].Bio, "Engineer")
	assert.Equal(t, resultptr[1].Email, "azer@roadbeats.com")

	var results []*UserProfile
	err = DB.Read(&results, "SELECT * FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, len(results), 3)
	assert.Equal(t, results[0].Name, "Nova")
	assert.Equal(t, results[0].Bio, "Photographer")
	assert.Equal(t, results[0].Email, "nova@roadbeats.com")
	assert.Equal(t, results[1].Name, "Azer")
	assert.Equal(t, results[1].Bio, "Engineer")
	assert.Equal(t, results[1].Email, "azer@roadbeats.com")

	var notmatching []*UserProfile
	err = DB.Read(&notmatching, "SELECT * FROM user_profiles WHERE name='not matching'")
	assert.Nil(t, err)
	assert.Equal(t, len(notmatching), 0)
}

func TestReadingSingleRow(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	nova := UserProfile{}
	err := DB.Read(&nova, "SELECT * FROM user_profiles WHERE name = ?", "Nova")
	assert.Nil(t, err)
	assert.Equal(t, nova.Id, 1)
	assert.Equal(t, nova.Name, "Nova")
	assert.Equal(t, nova.Bio, "Photographer")
	assert.Equal(t, nova.Email, "nova@roadbeats.com")

	var azer *UserProfile = &UserProfile{}
	err = DB.Read(azer, "SELECT * FROM user_profiles WHERE name = ?", "Azer")
	assert.Nil(t, err)
	assert.Equal(t, azer.Id, 2)
	assert.Equal(t, azer.Name, "Azer")
	assert.Equal(t, azer.Bio, "Engineer")
	assert.Equal(t, azer.Email, "azer@roadbeats.com")

	var az UserProfile
	err = DB.Read(&az, "SELECT * FROM user_profiles WHERE name = ?", "Azer")
	assert.Nil(t, err)
	assert.Equal(t, az.Id, 2)
	assert.Equal(t, az.Name, "Azer")
	assert.Equal(t, az.Bio, "Engineer")
	assert.Equal(t, az.Email, "azer@roadbeats.com")

	no := UserProfile{}
	err = DB.Read(&no, "SELECT * FROM user_profiles WHERE name = ?", "Not matching")
	assert.NotNil(t, err)

	DB.DropTables(UserProfile{})
}

func TestGeneratingQueries(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	result := []UserProfile{}
	err := DB.Read(&result, "SELECT * FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, len(result), 3)
	assert.Equal(t, result[0].Name, "Nova")
	assert.Equal(t, result[0].Bio, "Photographer")
	assert.Equal(t, result[0].Email, "nova@roadbeats.com")
	assert.Equal(t, result[1].Name, "Azer")
	assert.Equal(t, result[1].Bio, "Engineer")
	assert.Equal(t, result[1].Email, "azer@roadbeats.com")
	assert.Equal(t, result[2].Name, "Hola")
	assert.Equal(t, result[2].Bio, "")
	assert.Equal(t, result[2].Email, "hola@roadbeats.com")

	nova := UserProfile{}
	err = DB.Read(&nova, "SELECT * FROM user_profiles WHERE name=?", "Nova")
	assert.Nil(t, err)
	assert.Equal(t, nova.Name, "Nova")
	assert.Equal(t, nova.Bio, "Photographer")
	assert.Equal(t, nova.Email, "nova@roadbeats.com")

	DB.DropTables(UserProfile{})
}

func TestScanningToCustomValues(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	names := []string{}
	err := DB.Read(&names, "SELECT name FROM user_profiles ORDER BY id ASC")
	assert.Nil(t, err)
	assert.Equal(t, len(names), 3)
	assert.Equal(t, names[0], "Nova")
	assert.Equal(t, names[1], "Azer")

	name := ""
	err = DB.Read(&name, "SELECT name FROM user_profiles WHERE id=1")
	assert.Nil(t, err)
	assert.Equal(t, name, "Nova")

	count := 0
	err = DB.Read(&count, "SELECT COUNT(id) FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, count, 3)

	DB.DropTables(UserProfile{})
}

func TestScanningToNullTypes(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	nova := UserProfileNull{}
	err := DB.Read(&nova, "SELECT * FROM user_profiles WHERE name = ?", "Nova")
	assert.Nil(t, err)

	assert.Equal(t, nova.Id.Int64, int64(1))
	assert.Equal(t, nova.Name.String, "Nova")
	assert.Equal(t, nova.Bio.String, "Photographer")
	assert.Equal(t, nova.Email.String, "nova@roadbeats.com")

	DB.DropTables(UserProfile{})
}

func TestUnexistingFields(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	nova := UserProfile{}
	err := DB.Read(&nova, "SELECT u.*, COUNT(u.id) as ucount FROM user_profiles u WHERE name=?", "Nova")
	assert.Nil(t, err)
	assert.Equal(t, nova.Name, "Nova")
	assert.Equal(t, nova.Bio, "Photographer")
	assert.Equal(t, nova.Email, "nova@roadbeats.com")

	DB.DropTables(UserProfile{})
}

func CreateUserProfiles() error {
	if err := DB.ResetTables(UserProfile{}); err != nil {
		return err
	}

	if err := DB.Create(UserProfile{
		Name:       "Nova",
		Bio:        "Photographer",
		Email:      "nova@roadbeats.com",
		Attachment: []byte("{ \"nova\": \"bar\" }"),
	}); err != nil {
		return err
	}

	if err := DB.Create(UserProfile{
		Name:       "Azer",
		Bio:        "Engineer",
		Email:      "azer@roadbeats.com",
		Attachment: []byte("{ \"azer\": \"bar\" }"),
	}); err != nil {
		return err
	}

	if err := DB.Create(UserProfile{
		Name:       "Hola",
		Email:      "hola@roadbeats.com",
		Attachment: []byte("{ \"hola\": \"bar\" }"),
	}); err != nil {
		return err
	}

	return nil
}
