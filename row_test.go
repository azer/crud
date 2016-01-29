package crud_test

import (
	"github.com/azer/crud"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGettingRowValues(t *testing.T) {
	rows, err := crud.GetRowValuesOf(UserProfile{})
	assert.Nil(t, err)
	assert.Equal(t, len(rows), 4)

	rows, err = crud.GetRowValuesOf(UserProfile{
		Name:  "Azer",
		Email: "azer@roadbeats.com",
		Modified: 9223372036854775807,
	})

	assert.Nil(t, err)
	assert.Equal(t, len(rows), 4)
	assert.Equal(t, rows[0].SQLColumn, "name")
	assert.Equal(t, rows[0].Value.(string), "Azer")
	assert.Equal(t, rows[1].SQLColumn, "bio")
	assert.Equal(t, rows[1].Value.(string), "")
	assert.Equal(t, rows[2].SQLColumn, "email")
	assert.Equal(t, rows[2].Value.(string), "azer@roadbeats.com")
	assert.Equal(t, rows[3].SQLColumn, "modified")
	assert.Equal(t, rows[3].Value.(int64), int64(9223372036854775807))

	rows, err = crud.GetRowValuesOf(Post{})
	assert.Nil(t, err)
	assert.Equal(t, len(rows), 3)

	rows, err = crud.GetRowValuesOf(Post{
		Title: "Hello World",
		Text:  "It's still 2015 here",
	})

	assert.Nil(t, err)
	assert.Equal(t, rows[0].SQLColumn, "title")
	assert.Equal(t, rows[0].Value.(string), "Hello World")
	assert.Equal(t, rows[1].SQLColumn, "text")
	assert.Equal(t, rows[1].Value.(string), "It's still 2015 here")
}
