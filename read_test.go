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
	assert.Equal(t, result[0].Name, "Fuji")
	assert.Equal(t, result[0].Bio, "Engineer")
	assert.Equal(t, result[0].Email, "fuji@mitte.ai")
	assert.Equal(t, result[1].Name, "John")
	assert.Equal(t, result[1].Bio, "Engineer")
	assert.Equal(t, result[1].Email, "azer@mitte.ai")
	assert.Equal(t, string(result[1].Attachment), "{ \"azer\": \"bar\" }")

	resultptr := []*UserProfile{}
	err = DB.Read(&resultptr, "SELECT * FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, len(resultptr), 3)
	assert.Equal(t, resultptr[0].Name, "Fuji")
	assert.Equal(t, resultptr[0].Bio, "Engineer")
	assert.Equal(t, resultptr[0].Email, "fuji@mitte.ai")
	assert.Equal(t, resultptr[1].Name, "John")
	assert.Equal(t, resultptr[1].Bio, "Engineer")
	assert.Equal(t, resultptr[1].Email, "azer@mitte.ai")

	var results []*UserProfile
	err = DB.Read(&results, "SELECT * FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, len(results), 3)
	assert.Equal(t, results[0].Name, "Fuji")
	assert.Equal(t, results[0].Bio, "Engineer")
	assert.Equal(t, results[0].Email, "fuji@mitte.ai")
	assert.Equal(t, results[1].Name, "John")
	assert.Equal(t, results[1].Bio, "Engineer")
	assert.Equal(t, results[1].Email, "azer@mitte.ai")

	var notmatching []*UserProfile
	err = DB.Read(&notmatching, "SELECT * FROM user_profiles WHERE name='not matching'")
	assert.Nil(t, err)
	assert.Equal(t, len(notmatching), 0)
}

func TestReadingSingleRow(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	fuji := UserProfile{}
	err := DB.Read(&fuji, "SELECT * FROM user_profiles WHERE name = ?", "Fuji")
	assert.Nil(t, err)
	assert.Equal(t, fuji.Id, 1)
	assert.Equal(t, fuji.Name, "Fuji")
	assert.Equal(t, fuji.Bio, "Engineer")
	assert.Equal(t, fuji.Email, "fuji@mitte.ai")

	var azer *UserProfile = &UserProfile{}
	err = DB.Read(azer, "SELECT * FROM user_profiles WHERE name = ?", "John")
	assert.Nil(t, err)
	assert.Equal(t, azer.Id, 2)
	assert.Equal(t, azer.Name, "John")
	assert.Equal(t, azer.Bio, "Engineer")
	assert.Equal(t, azer.Email, "azer@mitte.ai")

	var az UserProfile
	err = DB.Read(&az, "SELECT * FROM user_profiles WHERE name = ?", "John")
	assert.Nil(t, err)
	assert.Equal(t, az.Id, 2)
	assert.Equal(t, az.Name, "John")
	assert.Equal(t, az.Bio, "Engineer")
	assert.Equal(t, az.Email, "azer@mitte.ai")

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
	assert.Equal(t, result[0].Name, "Fuji")
	assert.Equal(t, result[0].Bio, "Engineer")
	assert.Equal(t, result[0].Email, "fuji@mitte.ai")
	assert.Equal(t, result[1].Name, "John")
	assert.Equal(t, result[1].Bio, "Engineer")
	assert.Equal(t, result[1].Email, "azer@mitte.ai")
	assert.Equal(t, result[2].Name, "Hola")
	assert.Equal(t, result[2].Bio, "")
	assert.Equal(t, result[2].Email, "hola@mitte.ai")

	fuji := UserProfile{}
	err = DB.Read(&fuji, "SELECT * FROM user_profiles WHERE name=?", "Fuji")
	assert.Nil(t, err)
	assert.Equal(t, fuji.Name, "Fuji")
	assert.Equal(t, fuji.Bio, "Engineer")
	assert.Equal(t, fuji.Email, "fuji@mitte.ai")

	DB.DropTables(UserProfile{})
}

func TestScanningToCustomValues(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	names := []string{}
	err := DB.Read(&names, "SELECT name FROM user_profiles ORDER BY id ASC")
	assert.Nil(t, err)
	assert.Equal(t, len(names), 3)
	assert.Equal(t, names[0], "Fuji")
	assert.Equal(t, names[1], "John")

	name := ""
	err = DB.Read(&name, "SELECT name FROM user_profiles WHERE id=1")
	assert.Nil(t, err)
	assert.Equal(t, name, "Fuji")

	count := 0
	err = DB.Read(&count, "SELECT COUNT(id) FROM user_profiles")
	assert.Nil(t, err)
	assert.Equal(t, count, 3)

	DB.DropTables(UserProfile{})
}

func TestScanningToNullTypes(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	fuji := UserProfileNull{}
	err := DB.Read(&fuji, "SELECT * FROM user_profiles WHERE name = ?", "Fuji")
	assert.Nil(t, err)

	assert.Equal(t, fuji.Id.Int64, int64(1))
	assert.Equal(t, fuji.Name.String, "Fuji")
	assert.Equal(t, fuji.Bio.String, "Engineer")
	assert.Equal(t, fuji.Email.String, "fuji@mitte.ai")

	DB.DropTables(UserProfile{})
}

func TestUnexistingFields(t *testing.T) {
	assert.Nil(t, CreateUserProfiles())

	fuji := UserProfile{}
	// Test that extra columns in SELECT don't break scanning
	// Using a literal value instead of COUNT to avoid GROUP BY issues
	err := DB.Read(&fuji, "SELECT u.*, 1 as extra_column FROM user_profiles u WHERE name=?", "Fuji")
	assert.Nil(t, err)
	assert.Equal(t, fuji.Name, "Fuji")
	assert.Equal(t, fuji.Bio, "Engineer")
	assert.Equal(t, fuji.Email, "fuji@mitte.ai")

	DB.DropTables(UserProfile{})
}

func CreateUserProfiles() error {
	if err := DB.ResetTables(UserProfile{}); err != nil {
		return err
	}

	if err := DB.Create(UserProfile{
		Name:       "Fuji",
		Bio:        "Engineer",
		Email:      "fuji@mitte.ai",
		Attachment: []byte("{ \"fuji\": \"bar\" }"),
	}); err != nil {
		return err
	}

	if err := DB.Create(UserProfile{
		Name:       "John",
		Bio:        "Engineer",
		Email:      "azer@mitte.ai",
		Attachment: []byte("{ \"azer\": \"bar\" }"),
	}); err != nil {
		return err
	}

	if err := DB.Create(UserProfile{
		Name:       "Hola",
		Email:      "hola@mitte.ai",
		Attachment: []byte("{ \"hola\": \"bar\" }"),
	}); err != nil {
		return err
	}

	return nil
}
