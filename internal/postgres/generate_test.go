package postgres_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	text "github.com/matthewmueller/go-text"
	"github.com/matthewmueller/pogo"
	"github.com/matthewmueller/pogo/internal/postgres"
	"github.com/matthewmueller/pogo/internal/testutil"
	"github.com/pkg/errors"
	"github.com/tj/assert"
)

func TestGo(t *testing.T) {
	url := os.Getenv("POSTGRES_URL")
	assert.NotEmpty(t, url)
	// tests := filter(tests, "pg")
	cwd, err := os.Getwd()
	assert.NoError(t, err)
	tmpdir := filepath.Join(cwd, "tmp")
	assert.NoError(t, os.RemoveAll(tmpdir))
	defer func() {
		if !t.Failed() {
			assert.NoError(t, os.RemoveAll(tmpdir))
		}
	}()

	for _, test := range tests {
		name := testutil.Name(test)
		t.Run(name, func(t *testing.T) {
			pg, err := postgres.Open(url)
			assert.NoError(t, err)
			defer pg.Close()

			if test.After != "" {
				_, err = pg.Exec(test.After)
				assert.NoError(t, err)
			}
			if test.Before != "" {
				_, err = pg.Exec(test.Before)
				assert.NoError(t, err)
			}

			testpath := filepath.Join(tmpdir, text.Snake(name))
			err = os.MkdirAll(testpath, 0755)
			assert.NoError(t, err)
			pogopath := filepath.Join(testpath, "pogo")
			err = pogo.Generate(url, pogopath, test.Schema)
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
					`+imp(`pogo/event`)+`
					`+imp(`pogo/exercise`)+`
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

					actual, err := `+test.Func+`
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
					fmt.Println("# Expect:")
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
				fmt.Println("# Expect:")
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
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `team.FindByID(db, 2)`,
		Expect: `{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `team.Insert(db, team.New().Token(11).TeamName("1"))`,
		Expect: `{"id":1,"token":11,"team_name":"1","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams;
			drop extension if exists citext cascade;
		`,
		Func:   `team.UpdateByID(db, 2, team.New().StripeID("stripey"))`,
		Expect: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams;
			drop extension if exists citext cascade;
		`,
		Func:   `team.UpdateByID(db, 2, team.New().StripeID("stripey").Active(false))`,
		Expect: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","free_teammates":4,"cost_per_user":1}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `teammate.FindBySlackID(db, "b")`,
		Expect: `{"id":2,"team_id":1,"slack_id":"b","username":"b","timezone":"b"}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standup.FindByNameAndTeamID(db, "b", 1)`,
		Expect: `{"id":2,"team_id":1,"name":"b","channel":"b","time":"a","timezone":"a"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists blogs (
				name text not null
			);
			`,
		After: `
				drop table if exists blogs;
				drop extension if exists citext cascade;
			`,
		Func:  `blog.Find(db)`,
		Error: `blog not found`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `report.Find(db, report.NewFilter().TeammateID(1), report.NewOrder().Timestamp(report.DESC))`,
		Expect: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `report.FindMany(db, report.NewFilter().Status(enum.ReportStatusAsked))`,
		Expect: `[{"id":1,"teammate_id":1,"standup_id":1,"status":"ASKED","timestamp":1},{"id":2,"teammate_id":1,"standup_id":1,"status":"ASKED","timestamp":2}]`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standupteammate.Insert(db, standupteammate.New().StandupID(1).TeammateID(2).Status(enum.StandupTeammateStatusActive).Time("1:00").Owner(true))`,
		Expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().StandupID(1).TeammateID(2).Time("1:00").Status(enum.StandupTeammateStatusActive).Owner(true))`,
		Expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Status(enum.StandupTeammateStatusInvited).Owner(true))`,
		Expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"INVITED","time":"01:00:00","owner":true}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standupteammate.UpdateByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Owner(true))`,
		Expect: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j2', '* * * * 1-5');
		`,
		After: `
			drop table if exists crons;
			drop extension if exists citext cascade;
		`,
		Func:   `cron.DeleteByJob(db, "j2")`,
		Expect: `{"id":2,"job":"j2","frequency":"* * * * 1-5"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j2', '* * * * 1-5');
		`,
		After: `
			drop table if exists crons;
			drop extension if exists citext cascade;
		`,
		Func:   `cron.DeleteByID(db, 2)`,
		Expect: `{"id":2,"job":"j2","frequency":"* * * * 1-5"}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `teammate.UpsertBySlackID(db, "b", teammate.New().TeamID(2).Username("b").Timezone("b"))`,
		Expect: `{"id":2,"team_id":2,"slack_id":"b","username":"b","timezone":"b"}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `teammate.Upsert(db, teammate.New().ID(2).TeamID(2).SlackID("b").Username("b").Timezone("b"))`,
		Expect: `{"id":2,"team_id":2,"slack_id":"b","username":"b","timezone":"b"}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `teammate.Upsert(db, teammate.New().TeamID(2).SlackID("c").Username("c").Timezone("c"))`,
		Expect: `{"id":3,"team_id":2,"slack_id":"c","username":"c","timezone":"c"}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standup.FindMany(db, standup.NewFilter().TeamID(2), standup.NewOrder().Channel(standup.DESC))`,
		Expect: `[{"id":2,"team_id":2,"name":"b","channel":"b","time":"b","timezone":"b"},{"id":1,"team_id":2,"name":"a","channel":"a","time":"a","timezone":"a"}]`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standupteammate.FindMany(db, standupteammate.NewFilter().StandupID(1))`,
		Expect: `[{"id":1,"standup_id":1,"teammate_id":1,"status":"ACTIVE","time":"12:00:00"},{"id":2,"standup_id":1,"teammate_id":3,"status":"ACTIVE","time":"01:00:00","owner":true}]`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		After: `
			drop table if exists crons;
			drop extension if exists citext cascade;
		`,
		Func:   `cron.Delete(db, cron.NewFilter().JobContains("j1"))`,
		Expect: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		After: `
			drop table if exists crons;
			drop extension if exists citext cascade;
		`,
		Func:   `cron.Delete(db, cron.NewFilter().JobStartsWith("j1"))`,
		Expect: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		After: `
			drop table if exists crons;
			drop extension if exists citext cascade;
		`,
		Func:   `cron.DeleteMany(db, cron.NewFilter().JobContains("2"))`,
		Expect: `[{"id":2,"job":"j20","frequency":"* * * * 1-5"},{"id":3,"job":"j21","frequency":"* * * * 1-5"}]`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `teammate.Find(db, teammate.NewFilter().SlackIDIn("b", "c"))`,
		Expect: `{"id":2,"team_id":1,"slack_id":"b","username":"b","timezone":"b"}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `report.Find(db, report.NewFilter().TeammateID(2).StandupID(1).TimestampGt(2), report.NewOrder().Timestamp(report.DESC))`,
		Expect: `{"id":4,"teammate_id":2,"standup_id":1,"status":"COMPLETE","timestamp":4}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `team.FindMany(db)`,
		Expect: `[{"id":1,"token":11,"team_name":"a","active":true,"free_teammates":4,"cost_per_user":1},{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}]`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `standupteammate.Find(db, standupteammate.NewFilter().Owner(true).StandupIDIn(1, 3))`,
		Expect: `{"id":2,"standup_id":1,"teammate_id":3,"status":"ACTIVE","time":"01:00:00","owner":true}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `team.Update(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 44))`,
		Expect: `{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists teams (
				id serial primary key not null,
				token integer unique not null
			);
			insert into teams (token) values (11);
		`,
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Func:  `team.Update(db, team.New().Token(11), team.NewFilter().Token(10))`,
		Error: `team not found`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `team.UpdateMany(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 22))`,
		Expect: `[{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1},{"id":2,"token":22,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}]`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		After: `
			drop table if exists crons;
			drop extension if exists citext cascade;
		`,
		Func:  `cron.UpdateByID(db, 1, cron.New())`,
		Error: `cron.UpdateByID: no input provided`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		After: `
			drop table if exists crons;
			drop extension if exists citext cascade;
		`,
		Func:   `cron.UpdateByID(db, 1, cron.New().NullableFrequency(nil))`,
		Expect: `{"id":1,"job":"j1"}`,
	},
	{
		Schema: "jack",
		Before: `
			create extension if not exists citext;
			create schema if not exists jack;
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
		After: `
			drop table if exists jack.convos cascade;
			drop schema if exists jack cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `convo.FindMany(db, convo.NewFilter().UserIn("U0QS7USPJ", "U0QS890N5"))`,
		Expect: `[{"user":"U0QS7USPJ","intent":"standup_join","state":{}},{"user":"U0QS890N5","intent":"standup_join","state":{}}]`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists standups cascade;
			drop table if exists questions cascade;
			drop extension if exists citext cascade;
		`,
		Func: `question.InsertMany(
			db,
			question.New().StandupID(1).Order(2).Question("what's my name?"),
			question.New().StandupID(1).Order(1).Question("what's my age?"),
		)`,
		Expect: `[{"id":1,"standup_id":1,"order":2,"question":"what's my name?"},{"id":2,"standup_id":1,"order":1,"question":"what's my age?"}]`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.Find(db, event.NewFilter().TimeLte(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
		Expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.Find(db, event.NewFilter().TimeLt(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
		Expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.Find(db, event.NewFilter().TimeGte(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))`,
		Expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.Find(db, event.NewFilter().TimeGt(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))`,
		Expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.Find(db, event.NewFilter().Time(time.Date(2018, 9, 4, 0, 0, 0, 0, time.UTC)))`,
		Expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.Find(db, event.NewFilter().TimeNot(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
		Expect: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
	},
	{
		Name: "nullable_time_nil",
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.FindMany(db, event.NewFilter().NullableTime(nil))`,
		Expect: `[{"id":1},{"id":2},{"id":3}]`,
	},
	{
		Name: "nullable_time_ok",
		Before: `
			create extension if not exists citext;
			create table if not exists events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
		`,
		After: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `event.FindMany(db, event.NewFilter().NullableTime(&now))`,
		Expect: `[]`,
	},
	{
		Name: "empty_in_finds_nothing",
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `team.FindMany(db, team.NewFilter().IDIn().TokenIn(11, 22))`,
		Expect: `[]`,
	},
	{
		Name: "nullable_fk",
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `report.Find(db, report.NewFilter().NullablePostID(nil).TeammateID(1))`,
		Expect: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
	},
	{
		Before: `
			create extension if not exists citext;
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
		After: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Func:   `team.Find(db, team.NewFilter().Email("maTTMuelle@gmail.com"))`,
		Expect: `{"id":1,"token":11,"team_name":"a","email":"mattmuelle@gmail.com","active":true,"free_teammates":4,"cost_per_user":1}`,
	},
	{
		Before: `
			create table if not exists exercises (
				id serial primary key not null,
				distance decimal(5, 3) not null
			);
			insert into exercises (distance) values (12.213);
		`,
		After: `
			drop table if exists exercises cascade;
		`,
		Func:   `exercise.Find(db, exercise.NewFilter().Distance(12.213))`,
		Expect: `{"id":1,"distance":12.213}`,
	},
}
