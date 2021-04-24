package sqlite_test

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	text "github.com/matthewmueller/gotext"
	"github.com/matthewmueller/pogo"
	"github.com/matthewmueller/pogo/internal/sqlite"
	"github.com/matthewmueller/pogo/internal/testutil"
	"github.com/pkg/errors"
	"github.com/tj/assert"
)

func TestSQLite(t *testing.T) {
	uri := os.Getenv("SQLITE_URL")
	assert.NotEmpty(t, uri)
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	tmpdir := filepath.Join(cwd, "tmp")
	assert.NoError(t, os.RemoveAll(tmpdir))
	defer func() {
		if !t.Failed() {
			assert.NoError(t, os.RemoveAll(tmpdir))
		}
	}()

	u, err := url.Parse(uri)
	assert.NoError(t, err)
	path := filepath.Join(tmpdir, u.Path)
	err = os.MkdirAll(filepath.Dir(path), 0755)
	assert.NoError(t, err)
	dbpath := path + "?" + u.Query().Encode()

	for _, test := range tests {
		name := testutil.Name(test)
		t.Run(name, func(t *testing.T) {
			sq, err := sqlite.Open(dbpath)
			assert.NoError(t, err)
			defer sq.Close()

			if test.After != "" {
				_, err = sq.Exec(test.After)
				assert.NoError(t, err)
			}
			if test.Before != "" {
				_, err = sq.Exec(test.Before)
				assert.NoError(t, err, "seeding the db failed")
			}

			testpath := filepath.Join(tmpdir, text.Snake(name))
			err = os.MkdirAll(testpath, 0755)
			assert.NoError(t, err)

			exec := `actual, err = ` + test.Func
			if strings.Contains(test.Func, "=") {
				exec = test.Func
			}

			pogopath := filepath.Join(testpath, "pogo")
			err = pogo.Generate(dbpath, pogopath, test.Schema)
			assert.NoError(t, err)
			imp := testutil.GoImport(t, testpath)
			mainpath := filepath.Join(testpath, "main.go")
			stdout, stderr, remove := testutil.GoRun(t, mainpath, `
				package main

				import (
					"time"
					"database/sql"

					// sqlite db
					_ "github.com/mattn/go-sqlite3"

					pogo `+imp(`pogo`)+`
					`+imp(`pogo/blog`)+`
					`+imp(`pogo/post`)+`
				)

				func main() {
					// pogo.Log = func(sql string, v ...interface{}) {
					// 	io.WriteString(os.Stderr, sql+"\n")
					// }

					now := time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)
					_ = now

					var actual interface{}
					var err error

					// open the database
					db, err := sql.Open("sqlite3", "`+dbpath+`")
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						return
					}

					`+exec+`
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						return
					}

					buf, err := json.Marshal(actual)
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						return
					}

					fmt.Fprintf(os.Stdout, "%s", string(buf))
				}
			`)
			defer func() {
				if !t.Failed() {
					remove()
				}
			}()

			if stderr != "" {
				if test.Error != "" {
					if test.Error == stderr {
						return
					}
					fmt.Println("# expect:")
					fmt.Println(test.Error)
					fmt.Println()
					fmt.Println("# Actual:")
					fmt.Println(stderr)
					fmt.Println()
					t.Fatal(testutil.Diff(test.Error, stderr))
				}
				t.Fatal(errors.New(stderr))
			}

			if test.Expect != stdout {
				fmt.Println("# expect:")
				fmt.Println(test.Expect)
				fmt.Println()
				fmt.Println("# Actual:")
				fmt.Println(stdout)
				fmt.Println()
				t.Fatal(testutil.Diff(test.Expect, stdout))
			}
		})
	}
}

var tests = []testutil.Test{
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					name text not null
				);
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.Insert(db, blog.New().Name("a"))`,
		Expect: `1`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					name text not null
				);
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.InsertMany(db, blog.New().Name("a"), blog.New().Name("b"))`,
		Expect: `[1,2]`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					name text not null
				);
				insert into blogs (name) values ('a');
				insert into blogs (name) values ('b');
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.Find(db)`,
		Expect: `{"name":"a","rowid":1}`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					id integer primary key not null,
					name text not null
				);
				insert into blogs (name) values ('a');
				insert into blogs (name) values ('b');
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.Find(db)`,
		Expect: `{"id":1,"name":"a"}`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					name text not null
				);
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:  `blog.Find(db)`,
		Error: `blog not found`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					name text not null
				);
				insert into blogs (name) values ('a');
				insert into blogs (name) values ('b');
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.Find(db, blog.NewFilter().Name("b"))`,
		Expect: `{"name":"b","rowid":2}`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					id integer primary key not null,
					name text not null
				);
				insert into blogs (name) values ('a');
				insert into blogs (name) values ('b');
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.Find(db, blog.NewFilter().Name("b"))`,
		Expect: `{"id":2,"name":"b"}`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					name text not null
				);
				insert into blogs (name) values ('a');
				insert into blogs (name) values ('b');
			`,
		After: `
				drop table if exists blogs;
				drop table if exists posts;
			`,
		Func:   `blog.FindByRowID(db, 2)`,
		Expect: `{"name":"b","rowid":2}`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					id integer primary key not null,
					name text not null,
					url text not null,
					unique(url)
				);
				insert into blogs (name, url) values ('a', 'http://a.com');
				insert into blogs (name, url) values ('b', 'http://b.com');
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.FindByURL(db, "http://b.com")`,
		Expect: `{"id":2,"name":"b","url":"http://b.com"}`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					name text not null,
					url text not null,
					unique(url)
				);
				insert into blogs (name, url) values ('a', 'http://a.com');
				insert into blogs (name, url) values ('b', 'http://b.com');
			`,
		After: `
				drop table if exists blogs;
			`,
		Func:   `blog.FindByURL(db, "http://b.com")`,
		Expect: `{"name":"b","rowid":2,"url":"http://b.com"}`,
	},
	{
		Before: `
				pragma foreign_keys = 1;
				create table if not exists blogs (
					id integer primary key not null,
					name text not null
				);
				create table if not exists posts (
					id integer primary key not null,
					blog_id integer not null references blogs (id) on delete cascade on update cascade,
					title text not null,
					slug text not null,
					unique(blog_id, slug)
				);
				insert into blogs (name) values ('a');
				insert into blogs (name) values ('b');
				insert into posts (blog_id, title, slug) values (1, 'b', 's');
				insert into posts (blog_id, title, slug) values (2, 'b', 's');
			`,
		After: `
				drop table if exists blogs;
				drop table if exists posts;
			`,
		Func:   `post.FindByBlogIDAndSlug(db, 2, "s")`,
		Expect: `{"blog_id":2,"id":2,"slug":"s","title":"b"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			create table if not exists posts (
				id integer primary key not null,
				blog_id integer not null references blogs (id) on delete cascade on update cascade,
				title text not null,
				slug text not null,
				unique(blog_id, slug)
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
			drop table if exists posts;
		`,
		Func:   `blog.FindMany(db)`,
		Expect: `[{"id":1,"name":"a"},{"id":2,"name":"b"}]`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			create table if not exists posts (
				id integer primary key not null,
				blog_id integer not null references blogs (id) on delete cascade on update cascade,
				title text not null,
				slug text not null,
				unique(blog_id, slug)
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
			drop table if exists posts;
		`,
		Func:   `blog.FindMany(db, blog.NewFilter().Name("a"))`,
		Expect: `[{"id":1,"name":"a"}]`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Update(db, blog.New().Name("c"))`,
		Expect: `2`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Update(db, blog.New().Name("c"), blog.NewFilter().Name("b"))`,
		Expect: `1`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:  `blog.Update(db, blog.New().Name("c"), blog.NewFilter().Name("c"))`,
		Error: `blog not found`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `err = blog.UpdateByID(db, 1, blog.New().Name("c"))`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:  `err = blog.UpdateByID(db, 3, blog.New().Name("c"))`,
		Error: `blog not found`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `err = blog.UpdateByURL(db, "http://b.com", blog.New().Name("c"))`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:  `err = blog.UpdateByURL(db, "http://c.com", blog.New().Name("c"))`,
		Error: `blog not found`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			create table if not exists posts (
				id integer primary key not null,
				blog_id integer not null references blogs (id) on delete cascade on update cascade,
				title text not null,
				slug text not null,
				unique(blog_id, slug)
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
			insert into posts (blog_id, title, slug) values (1, 'b', 's');
			insert into posts (blog_id, title, slug) values (2, 'b', 's');
		`,
		After: `
			drop table if exists blogs;
			drop table if exists posts;
		`,
		Func:   `err = post.UpdateByBlogIDAndSlug(db, 2, "s", post.New().Title("c"))`,
		Expect: `null`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			create table if not exists posts (
				id integer primary key not null,
				blog_id integer not null references blogs (id) on delete cascade on update cascade,
				title text not null,
				slug text not null,
				unique(blog_id, slug)
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
			insert into posts (blog_id, title, slug) values (1, 'b', 's');
			insert into posts (blog_id, title, slug) values (2, 'b', 's');
		`,
		After: `
			drop table if exists blogs;
			drop table if exists posts;
		`,
		Func:  `err = post.UpdateByBlogIDAndSlug(db, 2, "z", post.New().Title("c"))`,
		Error: `post not found`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Delete(db)`,
		Expect: `2`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Delete(db, blog.NewFilter().URL("http://b.com"))`,
		Expect: `1`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:  `blog.Delete(db)`,
		Error: `blog not found`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `err = blog.DeleteByID(db, 1)`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:  `err = blog.DeleteByID(db, 10)`,
		Error: `blog not found`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `err = blog.DeleteByURL(db, "http://b.com")`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:  `err = blog.DeleteByURL(db, "http://c.com")`,
		Error: `blog not found`,
	},

	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Upsert(db, blog.New().ID(1).Name("c"))`,
		Expect: `0`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Upsert(db, blog.New().ID(2).Name("c"))`,
		Expect: `0`,
	},
	// TODO: handle edge cases like this
	// {
	// 	Before: `
	// 		create table if not exists blogs (
	// 			id integer primary key not null,
	// 			name text not null
	// 		);
	// 		insert into blogs (name) values ('a');
	// 		insert into blogs (name) values ('b');
	// 	`,
	// 	After: `
	// 		drop table if exists blogs;
	// 	`,
	// 	Func:   `blog.UpsertByID(db, -10, blog.New().Name("c"))`,
	// 	Expect: `0`,
	// },
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.UpsertByURL(db, "http://c.com", blog.New().Name("c"))`,
		Expect: `3`,
	},
	{
		Before: `
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				url text not null,
				unique(url)
			);
			insert into blogs (name, url) values ('a', 'http://a.com');
			insert into blogs (name, url) values ('b', 'http://b.com');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.UpsertByURL(db, "http://b.com", blog.New().Name("c"))`,
		Expect: `0`,
	},
	{
		Before: `
			create table if not exists migrate (
				version bigint not null primary key
			);
		`,
		After: `
			drop table if exists migrate;
		`,
		Func:   `migrate.Insert(db, migrate.New().Version(0))`,
		Expect: `1`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
			insert into variables (name, key, value, email) values ("a", "b", "c", "d");
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `variable.FindByKeyAndName(db, "b", "a")`,
		Expect: `{"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
			insert into variables (name, key, value, email) values ("a", "b", "c", "d");
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `variable.FindByKeyAndName(db, "b", "a")`,
		Expect: `{"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `variable.FindByEmailAndValue(db, "d", "c")`,
		Expect: `{"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
			insert into variables (name, key, value, email) values ("a", "b", "c", "d");
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `err = variable.UpdateByKeyAndName(db, "b", "a", variable.New().Email("e"))`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				count integer not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email, count) values ('a', 'b', 'c', 'd', 0);
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `err = variable.UpdateByEmailAndValue(db, "d", "c", variable.New().Count(1))`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
			insert into variables (name, key, value, email) values ("a", "b", "c", "d");
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `err = variable.DeleteByKeyAndName(db, "b", "a")`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `err = variable.DeleteByEmailAndValue(db, "d", "c")`,
		Expect: `null`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `variable.Upsert(db, variable.New().Key("a").Name("b").Value("c").Email("d"))`,
		Expect: `1`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				count integer not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email, count) values ('a', 'b', 'c', 'd', 0);
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `variable.UpsertByEmailAndValue(db, "d", "c", variable.New().Name("h").Key("g").Count(1))`,
		Expect: `0`,
	},
	{
		Before: `
			create table if not exists variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				count integer not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email, count) values ('a', 'b', 'c', 'd', 0);
		`,
		After: `
			drop table if exists variables;
		`,
		Func:   `variable.UpsertByEmailAndValue(db, "r", "c", variable.New().Name("h").Key("g").Count(1))`,
		Expect: `2`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameNot("abc"))`,
		Expect: `{"id":2,"name":"bad"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameContains("d"))`,
		Expect: `{"id":2,"name":"bad"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameNotContains("d"))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameStartsWith("ab"))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameNotStartsWith("ba"))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameEndsWith("bc"))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameNotEndsWith("ad"))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameLt("abd"))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameLte("abc"))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameGt("abd"))`,
		Expect: `{"id":2,"name":"bad"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().NameGte("bad"))`,
		Expect: `{"id":2,"name":"bad"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().ID(1))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().IDNot(2))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().IDLt(2))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().IDLte(1))`,
		Expect: `{"id":1,"name":"abc"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().IDGt(1))`,
		Expect: `{"id":2,"name":"bad"}`,
	},
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null
			);
			insert into blogs (name) values ('abc');
			insert into blogs (name) values ('bad');
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().IDGte(2))`,
		Expect: `{"id":2,"name":"bad"}`,
	},
	// TODO: fix floats
	// {
	// 	Before: `
	// 		pragma foreign_keys = 1;
	// 		create table if not exists blogs (
	// 			id integer primary key not null,
	// 			name text not null,
	// 			amount real not null
	// 		);
	// 		insert into blogs (name, amount) values ('abc', 1.03);
	// 		insert into blogs (name, amount) values ('bad', 2.05);
	// 	`,
	// 	After: `
	// 		drop table if exists blogs;
	// 	`,
	// 	Func:   `blog.Find(db, blog.NewFilter().Amount(1.03))`,
	// 	Expect: `{"id":1,"name":"abc","amount":1.03}`,
	// },
	{
		Before: `
			pragma foreign_keys = 1;
			create table if not exists blogs (
				id integer primary key not null,
				name text not null,
				active boolean not null
			);
			insert into blogs (name, active) values ('abc', true);
			insert into blogs (name, active) values ('bad', false);
		`,
		After: `
			drop table if exists blogs;
		`,
		Func:   `blog.Find(db, blog.NewFilter().Active(true))`,
		Expect: `{"active":true,"id":1,"name":"abc"}`,
	},
	// TODO: fix datetime filters
	// {
	// 	Before: `
	// 		pragma foreign_keys = 1;
	// 		create table if not exists blogs (
	// 			id integer primary key not null,
	// 			name text not null,
	// 			created_at text not null default date('now', 'start of year')
	// 		);
	// 		insert into blogs (name) values ('abc');
	// 		insert into blogs (name) values ('bad');
	// 	`,
	// 	After: `
	// 		drop table if exists blogs;
	// 	`,
	// 	Func:   `blog.Find(db, blog.NewFilter().CreatedAt(time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)))`,
	// 	Expect: `{"id":2,"name":"bad"}`,
	// },
	// {
	// 	Before: `
	// 		pragma foreign_keys = 1;
	// 		create table if not exists blogs (
	// 			id integer primary key not null,
	// 			name text not null,
	// 			active boolean not null
	// 		);
	// 		insert into blogs (name, active) values ('abc', true);
	// 		insert into blogs (name, active) values ('bad', false);
	// 	`,
	// 	After: `
	// 		drop table if exists blogs;
	// 	`,
	// 	Func:   `blog.Find(db, blog.NewFilter().Active(false).Name("bad"))`,
	// 	Expect: `{"id":2,"name":"bad"}`,
	// },
}
