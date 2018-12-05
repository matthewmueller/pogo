package pogo_test

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matthewmueller/pogo"

	text "github.com/matthewmueller/go-text"
	"github.com/matthewmueller/pogo/internal/postgres"
	"github.com/matthewmueller/pogo/internal/sqlite"
	"github.com/matthewmueller/pogo/internal/testutil"
	"github.com/tj/assert"
)

type test struct {
	dbs    string
	schema string
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
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	tmpdir := filepath.Join(cwd, "tmp")
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

			testpath := filepath.Join(tmpdir, text.Snake(name))
			err = os.MkdirAll(testpath, 0755)
			assert.NoError(t, err)

			schema := "public"
			if test.schema != "" {
				schema = test.schema
			}

			pogopath := filepath.Join(testpath, "pogo")
			err = pogo.Generate(url, pogopath, schema)
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
					fmt.Println("# expect:")
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
				fmt.Println("# expect:")
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
	assert.NoError(t, os.RemoveAll(tmpdir))
}

func TestSQLite(t *testing.T) {
	uri := os.Getenv("SQLITE_URL")
	assert.NotEmpty(t, uri)
	tests := filter(tests, "sqlite")
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	tmpdir := filepath.Join(cwd, "tmp")

	u, err := url.Parse(uri)
	assert.NoError(t, err)
	path := filepath.Join(tmpdir, u.Path)
	err = os.MkdirAll(filepath.Dir(path), 0755)
	assert.NoError(t, err)
	dbpath := path + "?" + u.Query().Encode()

	for _, test := range tests {
		name := formatName(test)
		t.Run(name, func(t *testing.T) {
			sq, err := sqlite.Open(dbpath)
			assert.NoError(t, err)
			defer sq.Close()

			if test.after != "" {
				_, err = sq.Exec(test.after)
				assert.NoError(t, err)
			}
			if test.before != "" {
				_, err = sq.Exec(test.before)
				assert.NoError(t, err)
			}

			testpath := filepath.Join(tmpdir, text.Snake(name))
			err = os.MkdirAll(testpath, 0755)
			assert.NoError(t, err)

			schema := "public"
			if test.schema != "" {
				schema = test.schema
			}

			pogopath := filepath.Join(testpath, "pogo")
			err = pogo.Generate(dbpath, pogopath, schema)
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

					`+imp(`pogo`)+`
					`+imp(`pogo/blog`)+`
					`+imp(`pogo/post`)+`
				)

				func main() {
					now := time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)
					_ = now

					// open the database
					db, err := sql.Open("sqlite3", "`+dbpath+`")
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						return
					}

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
					fmt.Println("# expect:")
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
				fmt.Println("# expect:")
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
	assert.NoError(t, os.RemoveAll(tmpdir))
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
			drop table if exists teams cascade;
		`,
		call:   `team.FindByID(db, 2)`,
		expect: `{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
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
		`,
		after: `
			drop table if exists teams cascade;
		`,
		call:   `team.Insert(db, team.New().Token(11).TeamName("1"))`,
		expect: `{"id":1,"token":11,"team_name":"1","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
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
		call:   `team.UpdateByID(db, 2, team.New().StripeID("stripey"))`,
		expect: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
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
		call:   `team.UpdateByID(db, 2, team.New().StripeID("stripey").Active(false))`,
		expect: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","free_teammates":4,"cost_per_user":1}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
		`,
		call:   `teammate.FindBySlackID(db, "b")`,
		expect: `{"id":2,"team_id":1,"slack_id":"b","username":"b","timezone":"b"}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
		`,
		call:   `standup.FindByNameAndTeamID(db, "b", 1)`,
		expect: `{"id":2,"team_id":1,"name":"b","channel":"b","time":"a","timezone":"a"}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table if not exists posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table if not exists reports (
				id serial primary key not null,
				teammate_id integer not null references teammates(id) on delete cascade,
				standup_id integer not null references standups(id) on delete cascade,
				post_id integer references posts(id) on delete cascade on update cascade,
				unique("teammate_id", post_id),
				"status" report_status not null default 'ASKED',
				"timestamp" serial not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into reports (teammate_id, standup_id) values (1, 1);
			insert into reports (teammate_id, standup_id) values (1, 1);
			insert into reports (teammate_id, standup_id, status) values (1, 1, 'COMPLETE');
			insert into reports (teammate_id, standup_id, status) values (2, 1, 'COMPLETE');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
		`,
		call:   `report.Find(db, report.NewFilter().TeammateID(1), report.NewOrder().Timestamp(report.DESC))`,
		expect: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table if not exists posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table if not exists reports (
				id serial primary key not null,
				teammate_id integer not null references teammates(id) on delete cascade,
				standup_id integer not null references standups(id) on delete cascade,
				post_id integer references posts(id) on delete cascade on update cascade,
				unique("teammate_id", post_id),
				"status" report_status not null default 'ASKED',
				"timestamp" serial not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into reports (teammate_id, standup_id) values (1, 1);
			insert into reports (teammate_id, standup_id) values (1, 1);
			insert into reports (teammate_id, standup_id, status) values (1, 1, 'COMPLETE');
			insert into reports (teammate_id, standup_id, status) values (2, 1, 'COMPLETE');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
		`,
		call:   `report.FindMany(db, report.NewFilter().Status(enum.ReportStatusAsked))`,
		expect: `[{"id":1,"teammate_id":1,"standup_id":1,"status":"ASKED","timestamp":1},{"id":2,"teammate_id":1,"standup_id":1,"status":"ASKED","timestamp":2}]`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table if not exists standups_teammates (
				id serial not null primary key,
				standup_id integer not null references standups(id),
				teammate_id integer not null references teammates(id),
				unique(standup_id, teammate_id),
				"status" standup_teammate_status not null,
				"time" time without time zone not null,
				owner bool not null default 'false'
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
		`,
		call:   `standupteammate.Insert(db, standupteammate.New().StandupID(1).TeammateID(2).Status(enum.StandupTeammateStatusActive).Time("1:00").Owner(true))`,
		expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table if not exists standups_teammates (
				id serial not null primary key,
				standup_id integer not null references standups(id),
				teammate_id integer not null references teammates(id),
				unique(standup_id, teammate_id),
				"status" standup_teammate_status not null,
				"time" time without time zone not null,
				owner bool not null default 'false'
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
		`,
		call:   `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().StandupID(1).TeammateID(2).Time("1:00").Status(enum.StandupTeammateStatusActive).Owner(true))`,
		expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table if not exists standups_teammates (
				id serial not null primary key,
				standup_id integer not null references standups(id),
				teammate_id integer not null references teammates(id),
				unique(standup_id, teammate_id),
				"status" standup_teammate_status not null,
				"time" time without time zone not null,
				owner bool not null default 'false'
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into standups_teammates (standup_id, teammate_id, "time", "status", owner) values (1, 2, '12:00', 'ACTIVE', false);
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
		`,
		call:   `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Status(enum.StandupTeammateStatusInvited).Owner(true))`,
		expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"INVITED","time":"01:00:00","owner":true}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table if not exists standups_teammates (
				id serial not null primary key,
				standup_id integer not null references standups(id),
				teammate_id integer not null references teammates(id),
				unique(standup_id, teammate_id),
				"status" standup_teammate_status not null,
				"time" time without time zone not null,
				owner bool not null default 'false'
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into standups_teammates (standup_id, teammate_id, "time", "status", owner) values (1, 2, '12:00', 'ACTIVE', false);
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
		`,
		call:   `standupteammate.UpdateByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Owner(true))`,
		expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j2', '* * * * 1-5');
		`,
		after: `
			drop table if exists crons;
		`,
		call:   `cron.DeleteByJob(db, "j2")`,
		expect: `{"id":2,"job":"j2","frequency":"* * * * 1-5"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j2', '* * * * 1-5');
		`,
		after: `
			drop table if exists crons;
		`,
		call:   `cron.DeleteByID(db, 2)`,
		expect: `{"id":2,"job":"j2","frequency":"* * * * 1-5"}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teams (token, team_name) values (22, 'b');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
		`,
		call:   `teammate.UpsertBySlackID(db, "b", teammate.New().TeamID(2).Username("b").Timezone("b"))`,
		expect: `{"id":2,"team_id":2,"slack_id":"b","username":"b","timezone":"b"}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teams (token, team_name) values (22, 'b');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
		`,
		call:   `teammate.UpsertByID(db, 2, teammate.New().TeamID(2).SlackID("b").Username("b").Timezone("b"))`,
		expect: `{"id":2,"team_id":2,"slack_id":"b","username":"b","timezone":"b"}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teams (token, team_name) values (22, 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (2, 'a', 'a', 'a', 'a');
			insert into standups (team_id, "name", channel, "time", timezone) values (2, 'b', 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'c', 'c', 'c', 'c');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
		`,
		call:   `standup.FindMany(db, standup.NewFilter().TeamID(2), standup.NewOrder().Channel(standup.DESC))`,
		expect: `[{"id":2,"team_id":2,"name":"b","channel":"b","time":"b","timezone":"b"},{"id":1,"team_id":2,"name":"a","channel":"a","time":"a","timezone":"a"}]`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table if not exists standups_teammates (
				id serial not null primary key,
				standup_id integer not null references standups(id),
				teammate_id integer not null references teammates(id),
				unique(standup_id, teammate_id),
				"status" standup_teammate_status not null,
				"time" time without time zone not null,
				owner bool not null default 'false'
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'c', 'c', 'c');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
			insert into standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 1, 'ACTIVE', '12:00', false);
			insert into standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 3, 'ACTIVE', '1:00', true);
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
		`,
		call:   `standupteammate.FindMany(db, standupteammate.NewFilter().StandupID(1))`,
		expect: `[{"id":1,"standup_id":1,"teammate_id":1,"status":"ACTIVE","time":"12:00:00"},{"id":2,"standup_id":1,"teammate_id":3,"status":"ACTIVE","time":"01:00:00","owner":true}]`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		after: `
			drop table if exists crons;
		`,
		call:   `cron.Delete(db, cron.NewFilter().JobContains("j1"))`,
		expect: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		after: `
			drop table if exists crons;
		`,
		call:   `cron.Delete(db, cron.NewFilter().JobStartsWith("j1"))`,
		expect: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		after: `
			drop table if exists crons;
		`,
		call:   `cron.DeleteMany(db, cron.NewFilter().JobContains("2"))`,
		expect: `[{"id":2,"job":"j20","frequency":"* * * * 1-5"},{"id":3,"job":"j21","frequency":"* * * * 1-5"}]`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
		`,
		call:   `teammate.Find(db, teammate.NewFilter().SlackIDIn("b", "c"))`,
		expect: `{"id":2,"team_id":1,"slack_id":"b","username":"b","timezone":"b"}`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table if not exists posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table if not exists reports (
				id serial primary key not null,
				teammate_id integer not null references teammates(id) on delete cascade,
				standup_id integer not null references standups(id) on delete cascade,
				post_id integer references posts(id) on delete cascade on update cascade,
				unique("teammate_id", post_id),
				"status" report_status not null default 'ASKED',
				"timestamp" serial not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into reports (teammate_id, standup_id) values (1, 1);
			insert into reports (teammate_id, standup_id) values (1, 1);
			insert into reports (teammate_id, standup_id, status) values (1, 1, 'COMPLETE');
			insert into reports (teammate_id, standup_id, status) values (2, 1, 'COMPLETE');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
		`,
		call:   `report.Find(db, report.NewFilter().TeammateID(2).StandupID(1).TimestampGt(2), report.NewOrder().Timestamp(report.DESC))`,
		expect: `{"id":4,"teammate_id":2,"standup_id":1,"status":"COMPLETE","timestamp":4}`,
	},
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
			drop table if exists teams cascade;
		`,
		call:   `team.FindMany(db)`,
		expect: `[{"id":1,"token":11,"team_name":"a","active":true,"free_teammates":4,"cost_per_user":1},{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}]`,
	},
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table if not exists standups_teammates (
				id serial not null primary key,
				standup_id integer not null references standups(id),
				teammate_id integer not null references teammates(id),
				unique(standup_id, teammate_id),
				"status" standup_teammate_status not null,
				"time" time without time zone not null,
				owner bool not null default 'false'
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'c', 'c', 'c');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
			insert into standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 1, 'ACTIVE', '12:00', false);
			insert into standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 3, 'ACTIVE', '1:00', true);
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
		`,
		call:   `standupteammate.Find(db, standupteammate.NewFilter().Owner(true).StandupIDIn(1, 3))`,
		expect: `{"id":2,"standup_id":1,"teammate_id":3,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
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
			insert into teams (token, team_name) values (33, 'c');
		`,
		after: `
			drop table if exists teams cascade;
		`,
		call:   `team.Update(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 44))`,
		expect: `{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
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
			insert into teams (token, team_name) values (33, 'c');
		`,
		after: `
			drop table if exists teams cascade;
		`,
		call:   `team.UpdateMany(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 22))`,
		expect: `[{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1},{"id":2,"token":22,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}]`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		after: `
			drop table if exists crons;
		`,
		call: `cron.UpdateByID(db, 1, cron.New())`,
		err:  `cron.UpdateByID: no input provided`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		after: `
			drop table if exists crons;
		`,
		call:   `cron.UpdateByID(db, 1, cron.New().NullableFrequency(nil))`,
		expect: `{"id":1,"job":"j1"}`,
	},
	{
		dbs:    "pg",
		schema: "jack",
		before: `
			create schema "jack";
			create table jack.convos (
				"user" text primary key,
				"intent" text,
				"slot" text,
				"state" jsonb not null default '{}'::jsonb,
				"ttl" int
			);
			insert into jack.convos ("user", "intent") values ('U0QS7USPJ', 'standup_join');
			insert into jack.convos ("user", "intent") values ('U0QS890N5', 'standup_join');
		`,
		after: `
			drop table if exists jack.convos cascade;
			drop schema "jack" cascade;
		`,
		call:   `convo.FindMany(db, convo.NewFilter().UserIn("U0QS7USPJ", "U0QS890N5"))`,
		expect: `[{"user":"U0QS7USPJ","intent":"standup_join","state":{}},{"user":"U0QS890N5","intent":"standup_join","state":{}}]`,
	},
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
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table if not exists questions (
				id serial not null primary key,
				standup_id integer not null references standups(id) on delete cascade,
				"order" smallint not null,
				question text not null,
				unique(standup_id, "order")
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists standups cascade;
			drop table if exists questions cascade;
		`,
		call: `question.InsertMany(
			db,
			question.New().StandupID(1).Order(2).Question("what's my name?"),
			question.New().StandupID(1).Order(1).Question("what's my age?"),
		)`,
		expect: `[{"id":1,"standup_id":1,"order":2,"question":"what's my name?"},{"id":2,"standup_id":1,"order":1,"question":"what's my age?"}]`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.Find(db, event.NewFilter().TimeLte(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
		expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.Find(db, event.NewFilter().TimeLt(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
		expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.Find(db, event.NewFilter().TimeGte(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))`,
		expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.Find(db, event.NewFilter().TimeGt(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))`,
		expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.Find(db, event.NewFilter().Time(time.Date(2018, 9, 4, 0, 0, 0, 0, time.UTC)))`,
		expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.Find(db, event.NewFilter().TimeNot(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
		expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.FindMany(db, event.NewFilter().NullableTime(nil))`,
		expect: `[{"id":1},{"id":2},{"id":3}]`,
	},
	{
		dbs: `pg`,
		before: `
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
		`,
		after: `
			drop table if exists events cascade;
		`,
		call:   `event.FindMany(db, event.NewFilter().NullableTime(&now))`,
		expect: `[]`,
	},
	{
		dbs:  `pg`,
		name: "empty_in_finds_nothing",
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
			drop table if exists teams cascade;
		`,
		call:   `team.FindMany(db, team.NewFilter().IDIn().TokenIn(11, 22))`,
		expect: `[]`,
	},
	{
		dbs:  `pg`,
		name: "nullable_fk",
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
			create table if not exists teammates (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				slack_id text unique not null,
				username text not null,
				first_name text,
				last_name text,
				email text,
				avatar text,
				timezone text not null
			);
			create table if not exists standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table if not exists posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table if not exists reports (
				id serial primary key not null,
				teammate_id integer not null references teammates(id) on delete cascade,
				standup_id integer not null references standups(id) on delete cascade,
				post_id integer references posts(id) on delete cascade on update cascade,
				unique("teammate_id", post_id),
				"status" report_status not null default 'ASKED',
				"timestamp" serial not null
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
			insert into teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			insert into posts (standup_id) values (1);
			insert into posts (standup_id) values (1);
			insert into reports (teammate_id, standup_id, post_id) values (1, 1, 1);
			insert into reports (teammate_id, standup_id, post_id) values (1, 1, 2);
			insert into reports (teammate_id, standup_id, status) values (1, 1, 'COMPLETE');
			insert into reports (teammate_id, standup_id, status) values (2, 1, 'COMPLETE');
		`,
		after: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
		`,
		call:   `report.Find(db, report.NewFilter().NullablePostID(nil).TeammateID(1))`,
		expect: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
	},
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
			insert into teams (token, email, team_name) values (11, 'mattmuelle@gmail.com', 'a');
		`,
		after: `
			drop table if exists teams cascade;
		`,
		call:   `team.Find(db, team.NewFilter().Email("maTTMuelle@gmail.com"))`,
		expect: `{"id":1,"token":11,"team_name":"a","email":"mattmuelle@gmail.com","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
	{
		dbs: `sqlite`,
		before: `
			create table if not exists blogs (
				name text not null
			);
			create table if not exists posts (
				title text not null,
				is_draft integer not null default true
			);
			create table if not exists posts (
				post_id integer references blogs (rowid) on delete cascade on update cascade,
				is_draft integer not null default true,
				slug text not null,
				title text not null,
				body text not null,
				created_at text not null default (now()),
				updated_at text not null default (now()),
				unique(slug)
			);
			insert into blogs (name) values ('a');
			insert into blogs (name) values ('b');
		`,
		after: `
			drop table if exists blogs;
			drop table if exists posts;
		`,
		call:   `blog.FindByID(db, 2)`,
		expect: `{"id":2,"name":"b"}`,
	},
}
