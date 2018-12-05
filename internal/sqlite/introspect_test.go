package sqlite_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/matthewmueller/pogo/internal/sqlite"
	"github.com/tj/assert"

	// sqlite db
	_ "github.com/mattn/go-sqlite3"
)

// Open an sqlite database
func Open(t testing.TB, uri string) (*sqlite.DB, func()) {
	db, err := sqlite.Open(uri)
	assert.NoError(t, err)

	return db, func() {
		assert.NoError(t, db.Close())
		assert.NoError(t, os.Remove(uri))
	}
}

func Exec(t testing.TB, db *sql.DB, schema string) {
	tx, err := db.Begin()
	assert.NoError(t, err)
	_, err = tx.Exec(schema)
	assert.NoError(t, err)
	err = tx.Commit()
	assert.NoError(t, err)
}

func TestIntrospect(t *testing.T) {
	db, close := Open(t, "./test.db")
	defer close()

	_, err := db.Exec(`
		create table if not exists blogs (
			name text not null
		);

		create table if not exists posts (
			blog_id integer not null references blogs(rowid) on delete cascade on update cascade,
			is_draft integer not null default true,
			title text not null,
			slug text not null,
			unique(blog_id, slug)
		);
	`)
	assert.NoError(t, err)

	schema, err := db.Introspect("public")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(schema.Tables))

	// tables
	blogs := schema.Tables[0]
	assert.Equal(t, "blogs", blogs.Name)
	posts := schema.Tables[1]
	assert.Equal(t, "posts", posts.Name)

	// columns
	assert.Equal(t, 2, len(blogs.Columns))
	assert.Equal(t, "rowid", blogs.Columns[0].Name)
	assert.Equal(t, "INTEGER", blogs.Columns[0].DataType)
	assert.Equal(t, true, blogs.Columns[0].NotNull)
	assert.Nil(t, blogs.Columns[0].DefaultValue)
	assert.Equal(t, true, blogs.Columns[0].IsPrimaryKey)

	assert.Equal(t, "name", blogs.Columns[1].Name)
	assert.Equal(t, "TEXT", blogs.Columns[1].DataType)
	assert.Equal(t, true, blogs.Columns[1].NotNull)
	assert.Nil(t, blogs.Columns[1].DefaultValue)
	assert.Equal(t, false, blogs.Columns[1].IsPrimaryKey)

	assert.Equal(t, "rowid", posts.Columns[0].Name)
	assert.Equal(t, "INTEGER", posts.Columns[0].DataType)
	assert.Equal(t, true, posts.Columns[0].NotNull)
	assert.Nil(t, posts.Columns[0].DefaultValue)
	assert.Equal(t, true, posts.Columns[0].IsPrimaryKey)

	assert.Equal(t, "blog_id", posts.Columns[1].Name)
	assert.Equal(t, "INTEGER", posts.Columns[1].DataType)
	assert.Equal(t, true, posts.Columns[1].NotNull)
	assert.Nil(t, posts.Columns[1].DefaultValue)
	assert.Equal(t, false, posts.Columns[1].IsPrimaryKey)

	assert.Equal(t, 5, len(posts.Columns))
	assert.Equal(t, "is_draft", posts.Columns[2].Name)
	assert.Equal(t, "INTEGER", posts.Columns[2].DataType)
	assert.Equal(t, true, posts.Columns[2].NotNull)
	assert.Equal(t, "true", *posts.Columns[2].DefaultValue)
	assert.Equal(t, false, posts.Columns[2].IsPrimaryKey)

	assert.Equal(t, "title", posts.Columns[3].Name)
	assert.Equal(t, "TEXT", posts.Columns[3].DataType)
	assert.Equal(t, true, posts.Columns[3].NotNull)
	assert.Nil(t, posts.Columns[3].DefaultValue)
	assert.Equal(t, false, posts.Columns[3].IsPrimaryKey)

	assert.Equal(t, "slug", posts.Columns[4].Name)
	assert.Equal(t, "TEXT", posts.Columns[4].DataType)
	assert.Equal(t, true, posts.Columns[4].NotNull)
	assert.Nil(t, posts.Columns[4].DefaultValue)
	assert.Equal(t, false, posts.Columns[4].IsPrimaryKey)

	// foreign keys
	assert.Equal(t, 0, len(blogs.ForeignKeys))
	assert.Equal(t, 1, len(posts.ForeignKeys))

	assert.Equal(t, "blog_id", posts.ForeignKeys[0].Name)
	assert.Equal(t, "INTEGER", posts.ForeignKeys[0].DataType)
	assert.Equal(t, "", posts.ForeignKeys[0].RefIndexName)
	assert.Equal(t, "blogs", posts.ForeignKeys[0].RefTableName)
	assert.Equal(t, "rowid", posts.ForeignKeys[0].RefColumnName)
	assert.Equal(t, 0, posts.ForeignKeys[0].KeyID)
	assert.Equal(t, 0, posts.ForeignKeys[0].SeqNo)
	assert.Equal(t, "CASCADE", posts.ForeignKeys[0].OnUpdate)
	assert.Equal(t, "CASCADE", posts.ForeignKeys[0].OnDelete)
	assert.Equal(t, "NONE", posts.ForeignKeys[0].Match)

	// indexes
	assert.Equal(t, 0, len(blogs.Indexes))

	assert.Equal(t, 1, len(posts.Indexes))
	assert.Equal(t, "sqlite_autoindex_posts_1", posts.Indexes[0].Name)
	assert.Equal(t, true, posts.Indexes[0].IsUnique)
	assert.Equal(t, false, posts.Indexes[0].IsPrimary)
	assert.Equal(t, 0, posts.Indexes[0].SeqNo)
	assert.Equal(t, "u", posts.Indexes[0].Origin)
	assert.Equal(t, false, posts.Indexes[0].IsPartial)

	// index columns
	assert.Equal(t, 2, len(posts.Indexes[0].Columns))

	assert.Equal(t, 0, posts.Indexes[0].Columns[0].SeqNo)
	assert.Equal(t, 0, posts.Indexes[0].Columns[0].Cid)
	assert.Equal(t, "blog_id", posts.Indexes[0].Columns[0].Name)
	assert.Equal(t, true, posts.Indexes[0].Columns[0].NotNull)
	assert.Equal(t, "INTEGER", posts.Indexes[0].Columns[0].DataType)

	assert.Equal(t, 1, posts.Indexes[0].Columns[1].SeqNo)
	assert.Equal(t, 3, posts.Indexes[0].Columns[1].Cid)
	assert.Equal(t, "slug", posts.Indexes[0].Columns[1].Name)
	assert.Equal(t, true, posts.Indexes[0].Columns[1].NotNull)
	assert.Equal(t, `TEXT`, posts.Indexes[0].Columns[1].DataType)
}
