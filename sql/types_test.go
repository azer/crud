package sql_test

import (
	"testing"

	"github.com/azer/crud/v2/sql"
	"github.com/stretchr/testify/assert"
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

	m, err = sql.MatchType("float32")
	assert.Nil(t, err)
	assert.Equal(t, m, "float")

	m, err = sql.MatchType("[]byte")
	assert.Nil(t, err)
	assert.Equal(t, m, "blob")

	m, err = sql.MatchType("sql.NullFloat32")
	assert.Nil(t, err)
	assert.Equal(t, m, "float")

	m, err = sql.MatchType("float64")
	assert.Nil(t, err)
	assert.Equal(t, m, "float")

	m, err = sql.MatchType("sql.NullFloat64")
	assert.Nil(t, err)
	assert.Equal(t, m, "float")

	m, err = sql.MatchType("bool")
	assert.Nil(t, err)
	assert.Equal(t, m, "tinyint")

	m, err = sql.MatchType("sql.NullBool")
	assert.Nil(t, err)
	assert.Equal(t, m, "tinyint")

	m, err = sql.MatchType("sql.NullInt64")
	assert.Nil(t, err)
	assert.Equal(t, m, "bigint")

	m, err = sql.MatchType("sql.NullString")
	assert.Nil(t, err)
	assert.Equal(t, m, "varchar")

	m, err = sql.MatchType("Yolo")
	assert.NotNil(t, err)
}
