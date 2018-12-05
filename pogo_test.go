package pogo_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matthewmueller/pogo/internal/vfs"

	text "github.com/matthewmueller/go-text"
	"github.com/matthewmueller/pogo/internal/postgres"
	"github.com/matthewmueller/pogo/internal/sqlite"
	"github.com/matthewmueller/pogo/internal/testutil"
	"github.com/tj/assert"
)

type test struct {
	dbs    string
	name   string
	before string
	after  string
	call   string
	expect string
	err    string
}

func formatName(t test) string {
	name := t.name
	if name == "" {
		name = testutil.Truncate(t.call, 20)
		if i := strings.Index(t.call, "("); i >= 0 {
			name = testutil.Truncate(t.call, i)
		}
	}
	return name
}

func filter(tt []test, db string) (tests []test) {
	for _, t := range tt {
		if t.dbs == "" {
			tests = append(tests, t)
			continue
		}
		dbs := strings.Split(t.dbs, " ")
		for _, d := range dbs {
			if d == db {
				tests = append(tests, t)
				break
			}
		}
	}
	return tests
}

func TestPG(t *testing.T) {
	url := os.Getenv("POSTGRES_URL")
	assert.NotEmpty(t, url)
	tests := filter(tests, "pg")
	for _, test := range tests {
		name := formatName(test)
		t.Run(name, func(t *testing.T) {
			pg, err := postgres.Open(url)
			assert.NoError(t, err)
			defer pg.Close()

			if test.after != "" {
				_, err = pg.Exec(test.after)
				assert.NoError(t, err)
			}
			if test.before != "" {
				_, err = pg.Exec(test.before)
				assert.NoError(t, err)
			}

			cwd, err := os.Getwd()
			assert.NoError(t, err)
			testpath := filepath.Join(cwd, "tmp", text.Snake(name))
			err = os.MkdirAll(testpath, 0755)
			assert.NoError(t, err)

			fs, err := pg.Generate([]string{"public"})
			assert.NoError(t, err)
			err = vfs.Write(fs, filepath.Join(testpath, "pogo"))
			assert.NoError(t, err)

			imp := testutil.GoImport(t, testpath)
			mainpath := filepath.Join(testpath, "main.go")
			stdout, stderr, remove := testutil.GoRun(t, mainpath, `
				package main

				import (
					"time"

					`+imp(`pogo`)+`
					`+imp(`pogo/enum`)+`
					`+imp(`pogo/team`)+`
					`+imp(`pogo/cron`)+`
					`+imp(`pogo/report`)+`
					`+imp(`pogo/standup`)+`
					`+imp(`pogo/question`)+`
					`+imp(`pogo/teammate`)+`
					`+imp(`pogo/standupteammate`)+`
				)

				func main() {
					now := time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)
					_ = now

					cfg, err := pgx.ParseConnectionString("`+url+`")
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						return
					}

					db, err := pgx.Connect(cfg)
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						return
					}
					defer db.Close()

					actual, err := `+test.call+`
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

			if stderr != "" {
				if test.err != "" {
					if test.err == stderr {
						return
					}
					fmt.Println("# Expected:")
					fmt.Println(test.err)
					fmt.Println()
					fmt.Println("# Actual:")
					fmt.Println(stderr)
					fmt.Println()
					t.Fatal(testutil.Diff(test.err, stderr))
				}
				t.Fatal(errors.New(stderr))
			}

			if test.expect != stdout {
				fmt.Println("# Expected:")
				fmt.Println(test.expect)
				fmt.Println()
				fmt.Println("# Actual:")
				fmt.Println(stdout)
				fmt.Println()
				t.Fatal(testutil.Diff(test.expect, stdout))
			}

			remove()
		})
	}
}

func TestSQLite(t *testing.T) {
	url := os.Getenv("SQLITE_URL")
	assert.NotEmpty(t, url)
	tests := filter(tests, "sqlite")
	for _, test := range tests {
		t.Run(formatName(test), func(t *testing.T) {
			sq, err := sqlite.Open(url)
			assert.NoError(t, err)
			defer sq.Close()
			fmt.Println(sq)

			if test.after != "" {
				_, err = sq.Exec(test.after)
				assert.NoError(t, err)
			}
			if test.before != "" {
				_, err = sq.Exec(test.before)
				assert.NoError(t, err)
			}

			vfs, err := sq.Generate([]string{"public"})
			assert.NoError(t, err)
			fmt.Println(vfs.ReadDir("/"))
		})
	}
}

var tests = []test{
	{
		dbs: `pg`,
		before: `
			create table if not exists teams (
				id serial primary key not null,
				token integer unique not null,
				team_name text not null,
				scope text[] not null default '{}',
				email citext,
				stripe_id text,
				active boolean not null default true,
				free_teammates integer not null default 4,
				cost_per_user integer not null default 1
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teams (token, team_name) values (22, 'b');
		`,
		after: `
			drop table if exists teams;
		`,
		call:   `team.FindByID(db, 2)`,
		expect: `{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
	{
		dbs: `sqlite`,
		before: `
			create table if not exists blogs (
				name text not null
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		after: `
			drop table if exists blogs;
		`,
		call:   `blog.FindByID(db, 2)`,
		expect: `{"id":2,"name":"b"}`,
	},
}
