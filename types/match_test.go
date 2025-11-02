package types_test

import (
	"testing"

	"github.com/azer/crud/v3/types"
	"github.com/stretchr/testify/assert"
)

// TestMatchTypeMySQL tests type matching for MySQL driver
func TestMatchTypeMySQL(t *testing.T) {
	// Integer types
	sqlType, length, err := types.MatchType("mysql", "int")
	assert.Nil(t, err)
	assert.Equal(t, "int", sqlType)
	assert.Equal(t, 11, length)

	sqlType, length, err = types.MatchType("mysql", "uint")
	assert.Nil(t, err)
	assert.Equal(t, "int", sqlType)
	assert.Equal(t, 11, length)

	sqlType, length, err = types.MatchType("mysql", "int64")
	assert.Nil(t, err)
	assert.Equal(t, "bigint", sqlType)
	assert.Equal(t, 20, length)

	sqlType, length, err = types.MatchType("mysql", "uint64")
	assert.Nil(t, err)
	assert.Equal(t, "bigint", sqlType)
	assert.Equal(t, 20, length)

	// String type
	sqlType, length, err = types.MatchType("mysql", "string")
	assert.Nil(t, err)
	assert.Equal(t, "varchar", sqlType)
	assert.Equal(t, 255, length)

	// Time type
	sqlType, length, err = types.MatchType("mysql", "time.Time")
	assert.Nil(t, err)
	assert.Equal(t, "timestamp", sqlType)
	assert.Equal(t, -1, length)

	// Float types
	sqlType, length, err = types.MatchType("mysql", "float32")
	assert.Nil(t, err)
	assert.Equal(t, "float", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("mysql", "float64")
	assert.Nil(t, err)
	assert.Equal(t, "float", sqlType)
	assert.Equal(t, -1, length)

	// Byte array types
	sqlType, length, err = types.MatchType("mysql", "[]byte")
	assert.Nil(t, err)
	assert.Equal(t, "blob", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("mysql", "[]uint8")
	assert.Nil(t, err)
	assert.Equal(t, "blob", sqlType)
	assert.Equal(t, -1, length)

	// Boolean type
	sqlType, length, err = types.MatchType("mysql", "bool")
	assert.Nil(t, err)
	assert.Equal(t, "tinyint", sqlType)
	assert.Equal(t, -1, length)

	// SQL null types
	sqlType, length, err = types.MatchType("mysql", "sql.NullFloat32")
	assert.Nil(t, err)
	assert.Equal(t, "float", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("mysql", "sql.NullFloat64")
	assert.Nil(t, err)
	assert.Equal(t, "float", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("mysql", "sql.NullInt64")
	assert.Nil(t, err)
	assert.Equal(t, "bigint", sqlType)
	assert.Equal(t, 20, length)

	sqlType, length, err = types.MatchType("mysql", "sql.NullString")
	assert.Nil(t, err)
	assert.Equal(t, "varchar", sqlType)
	assert.Equal(t, 255, length)

	sqlType, length, err = types.MatchType("mysql", "sql.NullBool")
	assert.Nil(t, err)
	assert.Equal(t, "tinyint", sqlType)
	assert.Equal(t, -1, length)
}

// TestMatchTypePostgreSQL tests type matching for PostgreSQL driver
func TestMatchTypePostgreSQL(t *testing.T) {
	// Integer types
	sqlType, length, err := types.MatchType("postgres", "int")
	assert.Nil(t, err)
	assert.Equal(t, "INTEGER", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "uint")
	assert.Nil(t, err)
	assert.Equal(t, "INTEGER", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "int64")
	assert.Nil(t, err)
	assert.Equal(t, "BIGINT", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "uint64")
	assert.Nil(t, err)
	assert.Equal(t, "BIGINT", sqlType)
	assert.Equal(t, -1, length)

	// String type
	sqlType, length, err = types.MatchType("postgres", "string")
	assert.Nil(t, err)
	assert.Equal(t, "VARCHAR", sqlType)
	assert.Equal(t, 255, length)

	// Time type
	sqlType, length, err = types.MatchType("postgres", "time.Time")
	assert.Nil(t, err)
	assert.Equal(t, "TIMESTAMP", sqlType)
	assert.Equal(t, -1, length)

	// Float types
	sqlType, length, err = types.MatchType("postgres", "float32")
	assert.Nil(t, err)
	assert.Equal(t, "DOUBLE PRECISION", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "float64")
	assert.Nil(t, err)
	assert.Equal(t, "DOUBLE PRECISION", sqlType)
	assert.Equal(t, -1, length)

	// Byte array types
	sqlType, length, err = types.MatchType("postgres", "[]byte")
	assert.Nil(t, err)
	assert.Equal(t, "BYTEA", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "[]uint8")
	assert.Nil(t, err)
	assert.Equal(t, "BYTEA", sqlType)
	assert.Equal(t, -1, length)

	// Boolean type
	sqlType, length, err = types.MatchType("postgres", "bool")
	assert.Nil(t, err)
	assert.Equal(t, "BOOLEAN", sqlType)
	assert.Equal(t, -1, length)

	// SQL null types
	sqlType, length, err = types.MatchType("postgres", "sql.NullFloat32")
	assert.Nil(t, err)
	assert.Equal(t, "DOUBLE PRECISION", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "sql.NullFloat64")
	assert.Nil(t, err)
	assert.Equal(t, "DOUBLE PRECISION", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "sql.NullInt64")
	assert.Nil(t, err)
	assert.Equal(t, "BIGINT", sqlType)
	assert.Equal(t, -1, length)

	sqlType, length, err = types.MatchType("postgres", "sql.NullString")
	assert.Nil(t, err)
	assert.Equal(t, "VARCHAR", sqlType)
	assert.Equal(t, 255, length)

	sqlType, length, err = types.MatchType("postgres", "sql.NullBool")
	assert.Nil(t, err)
	assert.Equal(t, "BOOLEAN", sqlType)
	assert.Equal(t, -1, length)
}

// TestMatchTypePostgreSQLAliases tests various PostgreSQL driver name aliases
func TestMatchTypePostgreSQLAliases(t *testing.T) {
	// Test "postgresql" alias
	sqlType, length, err := types.MatchType("postgresql", "string")
	assert.Nil(t, err)
	assert.Equal(t, "VARCHAR", sqlType)
	assert.Equal(t, 255, length)

	// Test "pgx" alias
	sqlType, length, err = types.MatchType("pgx", "int")
	assert.Nil(t, err)
	assert.Equal(t, "INTEGER", sqlType)
	assert.Equal(t, -1, length)
}

// TestMatchTypeUnknownType tests error handling for unknown Go types
func TestMatchTypeUnknownType(t *testing.T) {
	// MySQL
	_, _, err := types.MatchType("mysql", "UnknownType")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Can't match Go type 'UnknownType'")

	// PostgreSQL
	_, _, err = types.MatchType("postgres", "UnknownType")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Can't match Go type 'UnknownType'")
}

// TestNormalizeDriverName tests driver name normalization
func TestNormalizeDriverName(t *testing.T) {
	assert.Equal(t, "postgres", types.NormalizeDriverName("postgres"))
	assert.Equal(t, "postgres", types.NormalizeDriverName("postgresql"))
	assert.Equal(t, "postgres", types.NormalizeDriverName("pgx"))
	assert.Equal(t, "mysql", types.NormalizeDriverName("mysql"))
	assert.Equal(t, "mysql", types.NormalizeDriverName("unknown"))
	assert.Equal(t, "mysql", types.NormalizeDriverName(""))
}

// TestMatchTypeLengthDefaults tests that default lengths are correctly assigned
func TestMatchTypeLengthDefaults(t *testing.T) {
	// MySQL - types with default lengths
	_, length, _ := types.MatchType("mysql", "int")
	assert.Equal(t, 11, length, "MySQL int should have default length 11")

	_, length, _ = types.MatchType("mysql", "string")
	assert.Equal(t, 255, length, "MySQL varchar should have default length 255")

	// MySQL - types without default lengths
	_, length, _ = types.MatchType("mysql", "bool")
	assert.Equal(t, -1, length, "MySQL tinyint should have length -1")

	_, length, _ = types.MatchType("mysql", "[]byte")
	assert.Equal(t, -1, length, "MySQL blob should have length -1")

	// PostgreSQL - types with default lengths
	_, length, _ = types.MatchType("postgres", "string")
	assert.Equal(t, 255, length, "PostgreSQL VARCHAR should have default length 255")

	// PostgreSQL - types without default lengths
	_, length, _ = types.MatchType("postgres", "int")
	assert.Equal(t, -1, length, "PostgreSQL INTEGER should have length -1")

	_, length, _ = types.MatchType("postgres", "bool")
	assert.Equal(t, -1, length, "PostgreSQL BOOLEAN should have length -1")
}

// TestMatchTypeComprehensive tests all type mappings comprehensively
func TestMatchTypeComprehensive(t *testing.T) {
	testCases := []struct {
		driver      string
		goType      string
		expectedSQL string
		expectedLen int
	}{
		// MySQL mappings
		{"mysql", "int", "int", 11},
		{"mysql", "uint", "int", 11},
		{"mysql", "int64", "bigint", 20},
		{"mysql", "uint64", "bigint", 20},
		{"mysql", "string", "varchar", 255},
		{"mysql", "float32", "float", -1},
		{"mysql", "float64", "float", -1},
		{"mysql", "bool", "tinyint", -1},
		{"mysql", "time.Time", "timestamp", -1},
		{"mysql", "[]byte", "blob", -1},
		{"mysql", "[]uint8", "blob", -1},
		{"mysql", "sql.NullInt64", "bigint", 20},
		{"mysql", "sql.NullString", "varchar", 255},
		{"mysql", "sql.NullFloat32", "float", -1},
		{"mysql", "sql.NullFloat64", "float", -1},
		{"mysql", "sql.NullBool", "tinyint", -1},

		// PostgreSQL mappings
		{"postgres", "int", "INTEGER", -1},
		{"postgres", "uint", "INTEGER", -1},
		{"postgres", "int64", "BIGINT", -1},
		{"postgres", "uint64", "BIGINT", -1},
		{"postgres", "string", "VARCHAR", 255},
		{"postgres", "float32", "DOUBLE PRECISION", -1},
		{"postgres", "float64", "DOUBLE PRECISION", -1},
		{"postgres", "bool", "BOOLEAN", -1},
		{"postgres", "time.Time", "TIMESTAMP", -1},
		{"postgres", "[]byte", "BYTEA", -1},
		{"postgres", "[]uint8", "BYTEA", -1},
		{"postgres", "sql.NullInt64", "BIGINT", -1},
		{"postgres", "sql.NullString", "VARCHAR", 255},
		{"postgres", "sql.NullFloat32", "DOUBLE PRECISION", -1},
		{"postgres", "sql.NullFloat64", "DOUBLE PRECISION", -1},
		{"postgres", "sql.NullBool", "BOOLEAN", -1},
	}

	for _, tc := range testCases {
		t.Run(tc.driver+"_"+tc.goType, func(t *testing.T) {
			sqlType, length, err := types.MatchType(tc.driver, tc.goType)
			assert.Nil(t, err, "Should not error for %s type %s", tc.driver, tc.goType)
			assert.Equal(t, tc.expectedSQL, sqlType, "SQL type mismatch for %s %s", tc.driver, tc.goType)
			assert.Equal(t, tc.expectedLen, length, "Length mismatch for %s %s", tc.driver, tc.goType)
		})
	}
}
