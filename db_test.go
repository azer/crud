package crud_test

import (
	"testing"
)

func BenchmarkExecutingSQL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := DB.Client.Exec("SHOW TABLES LIKE 'shouldnotexist'")
		if err != nil {
			panic(err)
		}
	}
}
