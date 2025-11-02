package crud_test

import (
	"testing"

	"github.com/azer/crud/v4"
	"github.com/stretchr/testify/assert"
)

func TestSQLTableNameOf(t *testing.T) {
	assert.Equal(t, crud.SQLTableNameOf(getDriver(), &Foo{}), "foos")
	assert.Equal(t, crud.SQLTableNameOf(getDriver(), &FooSlice{}), "foo_slices")
	assert.Equal(t, crud.SQLTableNameOf(getDriver(), FooPTRSlice{}), "foo_ptr_slices")
}
