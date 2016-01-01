package sql_test

import (
	"github.com/azer/crud/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchType(t *testing.T) {
	m, err := sql.MatchType("int")
	assert.Nil(t, err)
	assert.Equal(t, m, "int")

	m, err = sql.MatchType("string")
	assert.Nil(t, err)
	assert.Equal(t, m, "varchar")

	m, err = sql.MatchType("time.Time")
	assert.Nil(t, err)
	assert.Equal(t, m, "timestamp")

	m, err = sql.MatchType("Yolo")
	assert.NotNil(t, err)
}
