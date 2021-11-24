package sql_test

import (
	"strings"
	"testing"

	"github.com/azer/crud/v2/sql"
	"github.com/stretchr/testify/assert"
)

func TestNewFieldQuery(t *testing.T) {
	f, err := sql.NewOptions("varchar(30) required default='yolo' name=key")
	assert.Nil(t, err)
	query := sql.NewFieldQuery(f)
	assert.Equal(t, query, "  `key` varchar(30) NOT NULL DEFAULT 'yolo'")

	f, err = sql.NewOptions("name=id type=int(11) auto-increment unique unsigned")
	assert.Nil(t, err)
	query = sql.NewFieldQuery(f)
	assert.Equal(t, query, "  `id` int(11) UNSIGNED UNIQUE AUTO_INCREMENT")

	f, err = sql.NewOptions("text name=bio")
	assert.Nil(t, err)
	query = sql.NewFieldQuery(f)
	assert.Equal(t, query, "  `bio` text")
}

func TestNewFieldQueries(t *testing.T) {
	id, err := sql.NewOptions("name=id type=int(11) auto-increment unique unsigned")
	assert.Nil(t, err)

	key, err := sql.NewOptions("varchar(30) required default='yolo' name=key")
	assert.Nil(t, err)

	bio, err := sql.NewOptions("text name=bio")
	assert.Nil(t, err)

	ignored, err := sql.NewOptions("-")
	assert.Nil(t, err)

	queries := sql.NewFieldQueries([]*sql.Options{id, key, bio, ignored})
	assert.Nil(t, err)

	lines := strings.Split(queries, "\n")
	assert.Equal(t, len(lines), 3)
	assert.Equal(t, lines[0], "  `id` int(11) UNSIGNED UNIQUE AUTO_INCREMENT,")
	assert.Equal(t, lines[1], "  `key` varchar(30) NOT NULL DEFAULT 'yolo',")
	assert.Equal(t, lines[2], "  `bio` text")
}

func TestNewTableQuery(t *testing.T) {
	id, err := sql.NewOptions("name=id type=int(11) auto-increment=100 unique unsigned primary-key")
	assert.Nil(t, err)

	key, err := sql.NewOptions("varchar(30) required default='yolo' name=key")
	assert.Nil(t, err)

	bio, err := sql.NewOptions("text name=bio")
	assert.Nil(t, err)

	ignored, err := sql.NewOptions("-")
	assert.Nil(t, err)

	query := sql.NewTableQuery("user", []*sql.Options{id, key, bio, ignored}, false)
	lines := strings.Split(query, "\n")
	assert.Equal(t, len(lines), 6)
	assert.Equal(t, lines[0], "CREATE TABLE `user` (")
	assert.Equal(t, lines[1], "  `id` int(11) UNSIGNED UNIQUE AUTO_INCREMENT,")
	assert.Equal(t, lines[2], "  `key` varchar(30) NOT NULL DEFAULT 'yolo',")
	assert.Equal(t, lines[3], "  `bio` text,")
	assert.Equal(t, lines[4], "  PRIMARY KEY (`id`)")
	assert.Equal(t, lines[5], ") AUTO_INCREMENT=100;")

	query = sql.NewTableQuery("user", []*sql.Options{id, key, bio, ignored}, true)
	lines = strings.Split(query, "\n")
	assert.Equal(t, lines[0], "CREATE TABLE IF NOT EXISTS `user` (")
}

func TestDropTableQuery(t *testing.T) {
	assert.Equal(t, sql.DropTableQuery("yolo", false), "DROP TABLE yolo")
	assert.Equal(t, sql.DropTableQuery("yolo", true), "DROP TABLE IF EXISTS yolo")
}

func TestSelectQuery(t *testing.T) {
	assert.Equal(t, sql.SelectQuery("yolo", []string{"foo", "bar"}), "SELECT foo,bar FROM yolo")
	assert.Equal(t, sql.SelectQuery("yolo", []string{}), "SELECT * FROM yolo")
}

func TestInsertQuery(t *testing.T) {
	assert.Equal(t, sql.InsertQuery("yolo", []string{"name", "email", "age"}), "INSERT INTO yolo (`name`,`email`,`age`) VALUES (?,?,?)")
}

func TestUpdateQuery(t *testing.T) {
	assert.Equal(t, sql.UpdateQuery("yolo", "id", []string{"name", "email", "age"}), "UPDATE yolo SET `name`=?, `email`=?, `age`=? WHERE id=?")
}

func TestDeleteQuery(t *testing.T) {
	assert.Equal(t, sql.DeleteQuery("yolo", "id"), "DELETE FROM yolo WHERE id=?")
}
