package crud_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	err := DB.ResetTables(UserProfile{})
	assert.Nil(t, err)

	azer := UserProfile{
		Name:  "John",
		Bio:   "I like photography",
		Email: "azer@mitte.ai",
	}

	err = DB.Create(azer)
	assert.Nil(t, err)

	DB.DropTables(UserProfile{})
}

func TestCreateAndRead(t *testing.T) {
	DB.ResetTables(UserProfile{})

	azer := UserProfile{
		Name:  "John",
		Bio:   "I like photography",
		Email: "azer@mitte.ai",
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
		Name: "John",
	}

	err := DB.Create(azer)
	assert.Nil(t, err)

	DB.DropTables(UserProfile{})
}

func TestEmbedding(t *testing.T) {
	err := DB.ResetTables(EmbeddedFoo{})
	if err != nil {
		panic(err)
	}

	foo := EmbeddedFoo{
		Foo: Foo{
			APIKey: "hi",
			YOLO:   true,
			Beast:  "span eggs",
		},
		Span: 123,
		Eggs: "yoo",
	}

	assert.Equal(t, foo.Id, 0)
	err = DB.CreateAndRead(&foo)
	assert.Nil(t, err)
	assert.Equal(t, foo.Id, 1)
	assert.Equal(t, foo.APIKey, "hi")
	assert.Equal(t, foo.Beast, "span eggs")

	DB.DropTables(EmbeddedFoo{})
}

func TestCreatingRenamedTableRow(t *testing.T) {
	DB.ResetTables(Post{})

	p := Post{
		Title:     "Foo",
		Text:      "bar",
		CreatedAt: time.Now(),
	}

	assert.Equal(t, p.Id, 0)
	err := DB.CreateAndRead(&p)
	assert.Nil(t, err)
	assert.Equal(t, p.Id, 1)

	DB.DropTables(Post{})
}
