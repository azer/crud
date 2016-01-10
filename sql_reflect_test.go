package crud_test

import (
	"github.com/azer/crud"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSQLTableNameOf(t *testing.T) {
	assert.Equal(t, crud.SQLTableNameOf(&Foo{}), "foo")
	assert.Equal(t, crud.SQLTableNameOf(&FooSlice{}), "foo_slice")
	assert.Equal(t, crud.SQLTableNameOf(FooPTRSlice{}), "foo_ptr_slice")
}
