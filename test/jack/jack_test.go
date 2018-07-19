package jack_test

import (
	"bytes"
	"errors"
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matthewmueller/pogo/testutil"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var url = os.Getenv("JACK_POSTGRES_URL")

var up = `
	-- Schema

	create schema jack;

	-- TEAMS

	-- add all the other columns
	create table if not exists jack.teams (
		id serial primary key not null,

		token integer unique not null,
		team_name text not null,
		scope text[] not null default '{}',
		email text,
		stripe_id text,
		active boolean not null default true,
		free_teammates integer not null default 4,
		cost_per_user integer not null default 1
	);

	-- add all the other columns
	create table if not exists jack.teammates (
		id serial primary key not null,
		team_id integer not null references jack.teams(id) on delete cascade,

		slack_id text unique not null,
		username text not null,
		first_name text,
		last_name text,
		email text,
		avatar text,
		timezone text not null
	);

	-- STANDUPS

	create table if not exists jack.standups (
		id serial primary key not null,
		team_id integer not null references jack.teams(id) on delete cascade,

		"name" text not null,
		channel text unique not null,
		"time" text not null,
		timezone text not null
	);

	create table if not exists jack.questions (
		id serial not null primary key,
		"order" smallint not null,
		standup_id integer not null references jack.standups(id) on delete cascade,

		question text not null
	);

	-- REPORTS

	create type jack.report_status as enum ('ASKED', 'SKIP', 'COMPLETE');

	-- add all the other columns
	create table if not exists jack.reports (
		id serial primary key not null,
		teammate_id integer not null references jack.teammates(id) on delete cascade,
		standup_id integer not null references jack.standups(id) on delete cascade,

		"status" jack.report_status not null default 'ASKED',
		"timestamp" serial not null
	);

	create table if not exists jack.answers (
		id serial not null primary key,
		question_id integer not null references jack.questions(id) on delete cascade,
		teammate_id integer not null references jack.teammates(id) on delete cascade,

		answer text not null
	);

	-- STANDUPS_TEAMMATES
	create table if not exists jack.standups_teammates (
		id serial not null primary key,
		standup_id integer not null references jack.standups(id),
		teammate_id integer not null references jack.teammates(id),
		unique(standup_id, teammate_id),

		"time" text not null,
		owner bool not null default 'false'
	);

	-- CRONS
	create table if not exists jack.crons (
		id serial not null primary key,
		"job" text unique not null,
		frequency text
	);

	-- jack add
	create function jack.add(a int, b int) returns int as $$
	begin
		return a + b;
	end;
	$$ language plpgsql immutable strict;
`

var down = `
	drop table if exists jack.crons cascade;

	drop table if exists jack.standups_teammates cascade;

	drop table if exists jack.questions cascade;

	drop table if exists jack.answers cascade;

	drop table if exists jack.reports cascade;
	drop type if exists jack.report_status;

	drop table if exists jack.standups cascade;

	drop table if exists jack.teammates cascade;

	drop table if exists jack.teams cascade;

	drop extension if exists pgcrypto cascade;

	drop function if exists jack.add(int, int);

	drop schema if exists jack cascade;
`

func TestPogo(t *testing.T) {
	conn, teardown := testutil.Connect(t, url)
	defer teardown()

	testutil.Exec(t, conn, down)
	testutil.Exec(t, conn, up)

	var tests = []struct {
		Setup    string
		Query    string
		Function string
		Expected string
		Error    string
	}{
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Query:    `SELECT * FROM teams WHERE id = $1`,
			Function: `team.FindByID(db, 2)`,
			Expected: `{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Query:    `INSERT INTO jack.teams (token, team_name) VALUES ($1, $2) RETURNING id`,
			Function: `team.Insert(db, team.New().Token(11).TeamName("1"))`,
			Expected: `{"id":1,"token":11,"team_name":"1","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Query:    `UPDATE jack.teams SET stripe_id = $1 WHERE id = $2 RETURNING *`,
			Function: `team.UpdateByID(db, 2, team.New().StripeID("stripey"))`,
			Expected: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Query:    `UPDATE jack.teams SET stripe_id = $1, active = false WHERE id = $2 RETURNING *`,
			Function: `team.UpdateByID(db, 2, team.New().StripeID("stripey").Active(false))`,
			Expected: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			`,
			Query:    `SELECT id FROM teammates WHERE slack_id = $1`,
			Function: `teammate.FindBySlackID(db, "b")`,
			Expected: `{"id":2,"team_id":1,"slack_id":"b","username":"b","timezone":"b"}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.reports (teammate_id, standup_id) values (1, 1);
				insert into jack.reports (teammate_id, standup_id) values (1, 1);
				insert into jack.reports (teammate_id, standup_id, status) values (1, 1, 'COMPLETE');
				insert into jack.reports (teammate_id, standup_id, status) values (2, 1, 'COMPLETE');
			`,
			Query:    `SELECT * FROM reports WHERE teammate_id = $1 ORDER BY "timestamp" DESC LIMIT 1`,
			Function: `report.Find(db, report.NewFilter().TeammateID(1), report.NewOrder().Timestamp(report.DESC))`,
			Expected: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			`,
			Query:    `INSERT INTO standups_teammates (standup_id, teammate_id, time) VALUES ($1, $2, $3) RETURNING *`,
			Function: `standupteammate.Insert(db, standupteammate.New().StandupID(1).TeammateID(2).Time("1:00").Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"time":"1:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			`,
			Query:    `INSERT INTO standups_teammates (standup_id, teammate_id) VALUES ($1, $2) ON CONFLICT (standup_id, teammate_id) DO UPDATE SET status = 'ACTIVE' RETURNING *`,
			Function: `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"time":"1:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "time", owner) values (1, 2, '12:00', false);
			`,
			Query:    `INSERT INTO standups_teammates (standup_id, teammate_id) VALUES ($1, $2) ON CONFLICT (standup_id, teammate_id) DO UPDATE SET status = 'ACTIVE' RETURNING *`,
			Function: `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"time":"1:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "time", owner) values (1, 2, '12:00', false);
			`,
			Query:    `UPDATE jack.standups_teammates SET "time" = '1:00', "owner" = true WHERE teammate_id = $1 AND standup_id = $2 RETURNING *`,
			Function: `standupteammate.UpdateByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"time":"1:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j2', '* * * * 1-5');
			`,
			Query:    `DELETE FROM crons WHERE job = $1 RETURNING *`,
			Function: `cron.DeleteByJob(db, "j2")`,
			Expected: `{"id":2,"job":"j2","frequency":"* * * * 1-5"}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j2', '* * * * 1-5');
			`,
			Query:    `DELETE FROM crons WHERE id = $1 RETURNING *`,
			Function: `cron.DeleteByID(db, 2)`,
			Expected: `{"id":2,"job":"j2","frequency":"* * * * 1-5"}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			`,
			Query:    `INSERT INTO teammates (slack_id, team_id) VALUES ($1, $2) ON CONFLICT (slack_id) DO UPDATE SET id = teammates.id RETURNING *`,
			Function: `teammate.UpsertBySlackID(db, "b", teammate.New().TeamID(2).Username("b").Timezone("b"))`,
			Expected: `{"id":2,"team_id":2,"slack_id":"b","username":"b","timezone":"b"}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			`,
			Query:    `INSERT INTO teammates (slack_id, team_id) VALUES ($1, $2) ON CONFLICT (slack_id) DO UPDATE SET id = teammates.id RETURNING *`,
			Function: `teammate.UpsertByID(db, 2, teammate.New().TeamID(2).SlackID("b").Username("b").Timezone("b"))`,
			Expected: `{"id":2,"team_id":2,"slack_id":"b","username":"b","timezone":"b"}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (2, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (2, 'b', 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'c', 'c', 'c', 'c');
			`,
			Query:    `SELECT * FROM standups WHERE team_id = $1`,
			Function: `standup.FindMany(db, standup.NewFilter().TeamID(2), standup.NewOrder().Channel(standup.DESC))`,
			Expected: `[{"id":2,"team_id":2,"name":"b","channel":"b","time":"b","timezone":"b"},{"id":1,"team_id":2,"name":"a","channel":"a","time":"a","timezone":"a"}]`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'c', 'c', 'c');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "time", owner) values (1, 1, '12:00', false);
				insert into jack.standups_teammates (standup_id, teammate_id, "time", owner) values (1, 3, '1:00', true);
			`,
			Query:    `SELECT * FROM standups_teammates WHERE standup_id = $1`,
			Function: `standupteammate.FindMany(db, standupteammate.NewFilter().StandupID(1))`,
			Expected: `[{"id":1,"standup_id":1,"teammate_id":1,"time":"12:00"},{"id":2,"standup_id":1,"teammate_id":3,"time":"1:00","owner":true}]`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Query:    `DELETE FROM crons WHERE job SIMILAR TO $1`,
			Function: `cron.Delete(db, cron.NewFilter().JobContains("j1"))`,
			Expected: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Query:    `DELETE FROM crons WHERE job SIMILAR TO $1`,
			Function: `cron.Delete(db, cron.NewFilter().JobStartsWith("j1"))`,
			Expected: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Query:    `DELETE FROM crons WHERE job SIMILAR TO $1`,
			Function: `cron.DeleteMany(db, cron.NewFilter().JobContains("2"))`,
			Expected: `[{"id":2,"job":"j20","frequency":"* * * * 1-5"},{"id":3,"job":"j21","frequency":"* * * * 1-5"}]`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			`,
			Query:    `SELECT * FROM teammates WHERE id = IN ($1)`,
			Function: `teammate.Find(db, teammate.NewFilter().SlackIDIn("b", "c"))`,
			Expected: `{"id":2,"team_id":1,"slack_id":"b","username":"b","timezone":"b"}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.reports (teammate_id, standup_id) values (1, 1);
				insert into jack.reports (teammate_id, standup_id) values (1, 1);
				insert into jack.reports (teammate_id, standup_id, status) values (1, 1, 'COMPLETE');
				insert into jack.reports (teammate_id, standup_id, status) values (2, 1, 'COMPLETE');
			`,
			Query:    `SELECT * FROM reports WHERE teammate_id = $1 AND standup_id = $2 AND "timestamp" > (timestamp '1d' - INTERVAL '1hr') ORDER BY "timestamp" DESC LIMIT 1`,
			Function: `report.Find(db, report.NewFilter().TeammateID(2).StandupID(1).TimestampGt(2), report.NewOrder().Timestamp(report.DESC))`,
			Expected: `{"id":4,"teammate_id":2,"standup_id":1,"status":"COMPLETE","timestamp":4}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Query:    `select * from teams`,
			Function: `team.FindMany(db)`,
			Expected: `[{"id":1,"token":11,"team_name":"a","active":true,"free_teammates":4,"cost_per_user":1},{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}]`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'c', 'c', 'c');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "time", owner) values (1, 1, '12:00', false);
				insert into jack.standups_teammates (standup_id, teammate_id, "time", owner) values (1, 3, '1:00', true);
			`,
			Query:    `SELECT * FROM standups_teammates WHERE owner = true AND standup_id = ANY($1)`,
			Function: `standupteammate.Find(db, standupteammate.NewFilter().Owner(true).StandupIDIn(1, 3))`,
			Expected: `{"id":2,"standup_id":1,"teammate_id":3,"time":"1:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
				insert into jack.teams (token, team_name) values (33, 'c');
			`,
			Query:    `UPDATE jack.teams SET team_name = 'cool' team_name WHERE token IN (11, 44) RETURNING *`,
			Function: `team.Update(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 44))`,
			Expected: `{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
				insert into jack.teams (token, team_name) values (33, 'c');
			`,
			Query:    `UPDATE jack.teams SET team_name = 'cool' team_name WHERE token IN (11, 44) RETURNING *`,
			Function: `team.UpdateMany(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 22))`,
			Expected: `[{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1},{"id":2,"token":22,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}]`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Query:    `UPDATE crons SET frequency = NULL WHERE id = $1`,
			Function: `cron.UpdateByID(db, 1, cron.New())`,
			Error:    `cron.UpdateByID: no input provided`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Query:    `UPDATE crons SET frequency = NULL WHERE id = $1`,
			Function: `cron.UpdateByID(db, 1, cron.New().NullableFrequency(nil))`,
			Expected: `{"id":1,"job":"j1"}`,
		},
		// {
		// 	Query:    `INSERT INTO crons (job, frequency, tz) VALUES ($1, $2, $3) ON CONFLICT (job) DO UPDATE SET frequency = concat($4::text, ' ', substring(crons.frequency from '[\\d\\-\\,\\*]+$')), tz = $3 RETURNING *`,
		// 	Expected: ``,
		// },
	}

	gopath := build.Default.GOPATH

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	importBase, err := filepath.Rel(filepath.Join(gopath, "src"), cwd)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll("./_tmp"); err != nil {
		t.Fatal(err)
	}

	if err := os.RemoveAll("./pogo"); err != nil {
		t.Fatal(err)
	}

	cleanup := testutil.Build(t, url, "jack", "./pogo")
	_ = cleanup
	// defer cleanup()

	for _, test := range tests {
		name := truncate(test.Function, 20)
		if i := strings.Index(test.Function, "("); i >= 0 {
			name = truncate(test.Function, i)
		}

		t.Run(name, func(t *testing.T) {
			testutil.Exec(t, conn, down)
			testutil.Exec(t, conn, up)

			if test.Setup != "" {
				testutil.Exec(t, conn, test.Setup)
			}

			stdout, stderr, remove := testutil.Run(t, name, `
				package main
	
				import (
					"`+importBase+`/pogo/enum"
					pogo "`+importBase+`/pogo"
					team "`+importBase+`/pogo/team"
					cron "`+importBase+`/pogo/cron"
					report "`+importBase+`/pogo/report"
					standup "`+importBase+`/pogo/standup"
					teammate "`+importBase+`/pogo/teammate"
					standupteammate "`+importBase+`/pogo/standup-teammate"
				)
	
				func main() {
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
	
					actual, err := `+test.Function+`
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
						return
					}
	
					buf, err := json.Marshal(actual)
					if err != nil {
						fmt.Fprintln(os.Stderr, err.Error())
						return
					}
					
					fmt.Fprintf(os.Stdout, "%s", string(buf))
				}
			`)

			if stderr != "" {
				if test.Error != "" {
					if test.Error == stderr {
						return
					}
					t.Fatal(diff(test.Error, stderr))
				}
				t.Fatal(errors.New(stderr))
			}

			if test.Expected != stdout {
				t.Fatal(diff(test.Expected, stdout))
			}

			remove()
		})
	}

	// cleanup()
}

func diff(expected, actual string) string {

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(expected, actual, false)

	var buf bytes.Buffer
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			buf.WriteString("\x1b[102m\x1b[30m")
			buf.WriteString(diff.Text)
			buf.WriteString("\x1b[0m")
		case diffmatchpatch.DiffDelete:
			buf.WriteString("\x1b[101m\x1b[30m")
			buf.WriteString(diff.Text)
			buf.WriteString("\x1b[0m")
		case diffmatchpatch.DiffEqual:
			buf.WriteString(diff.Text)
		}
	}

	result := buf.String()
	result = strings.Replace(result, "\\n", "\n", -1)
	result = strings.Replace(result, "\\t", "\t", -1)
	return result
}

func truncate(str string, num int) string {
	s := str
	if len(str) > num {
		s = str[0:num]
	}
	return s
}