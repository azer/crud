package crud_test

import (
	"github.com/azer/crud"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var DB *crud.DB

type UserProfile struct {
	Id    int    `json:"id" sql:"auto-increment primary-key required"`
	Name  string `json:"name" sql:"unique required"`
	Bio   string `json:"bio" sql:"type=text"`
	Email string `json:"e-mail" sql:"name=email"`
}

type Post struct {
	Id        int       `json:"id" sql:"auto-increment primary-key required"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	var err error
	DB, err = crud.Connect(os.Getenv("DATABASE_DRIVER"), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}

func TestPing(t *testing.T) {
	assert.Nil(t, DB.Ping())
}

func TestExecuteSQL(t *testing.T) {
	result, err := DB.Client.Exec("SHOW TABLES LIKE 'shouldnotexist'")
	assert.Nil(t, err)

	l, err := result.LastInsertId()
	assert.Equal(t, err, nil)
	assert.Equal(t, l, int64(0))

	a, err := result.RowsAffected()
	assert.Equal(t, err, nil)
	assert.Equal(t, a, int64(0))
}

func TestCreateTable(t *testing.T) {
	err := DB.CreateTable(UserProfile{}, Post{})
	assert.Nil(t, err)
	assert.True(t, DB.CheckIfTableExists("user_profile"))
	assert.True(t, DB.CheckIfTableExists("post"))
}

func TestDropTable(t *testing.T) {
	err := DB.DropTable(UserProfile{}, Post{})
	assert.Nil(t, err)
	assert.False(t, DB.CheckIfTableExists("user_profile"))
	assert.False(t, DB.CheckIfTableExists("post"))
}
