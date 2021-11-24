package crud_test

import (
	"testing"

	"github.com/azer/crud/v2"
	"github.com/stretchr/testify/assert"
)

func TestSQLTableNameOf(t *testing.T) {
	assert.Equal(t, crud.SQLTableNameOf(&Foo{}), "foos")
	assert.Equal(t, crud.SQLTableNameOf(&FooSlice{}), "foo_slices")
	assert.Equal(t, crud.SQLTableNameOf(FooPTRSlice{}), "foo_ptr_slices")
}
