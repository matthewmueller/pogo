package test

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/matthewmueller/pogo/testutil"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestPostgres(t *testing.T) {
	url := os.Getenv("JACK_POSTGRES_URL")

	up := `
	-- Schema
	create extension if not exists citext;

	create schema jack;

	-- TEAMS

	-- add all the other columns
	create table if not exists jack.teams (
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
		timezone text not null,

		unique(team_id, "name")
	);

	create table if not exists jack.questions (
		id serial not null primary key,
		standup_id integer not null references jack.standups(id) on delete cascade,

		"order" smallint not null,
		question text not null,

		unique(standup_id, "order")
	);

	-- posts table
	create table if not exists jack.posts (
		id serial primary key,
		standup_id integer not null references jack.standups(id) on delete cascade on update cascade,
		created_at timestamp with time zone default (now() at time zone 'utc')
	);

	-- REPORTS

	create type jack.report_status as enum ('ASKED', 'SKIP', 'COMPLETE');

	-- add all the other columns
	create table if not exists jack.reports (
		id serial primary key not null,
		teammate_id integer not null references jack.teammates(id) on delete cascade,
		standup_id integer not null references jack.standups(id) on delete cascade,

		-- post groups our reports
		post_id integer references jack.posts(id) on delete cascade on update cascade,
		unique("teammate_id", post_id),

		"status" jack.report_status not null default 'ASKED',
		"timestamp" serial not null
	);

	create table if not exists jack.answers (
		id serial not null primary key,
		question_id integer not null references jack.questions(id) on delete cascade,
		teammate_id integer not null references jack.teammates(id) on delete cascade,

		answer text not null
	);

	create type jack.standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');

	-- STANDUPS_TEAMMATES
	create table if not exists jack.standups_teammates (
		id serial not null primary key,
		standup_id integer not null references jack.standups(id),
		teammate_id integer not null references jack.teammates(id),
		unique(standup_id, teammate_id),

		"status" jack.standup_teammate_status not null,
		"time" time without time zone not null,
		owner bool not null default 'false'
	);

	-- CRONS
	create table if not exists jack.crons (
		id serial not null primary key,
		"job" text unique not null,
		frequency text
	);

	-- cron event
	create table if not exists jack.events (
		id serial not null primary key,
		"time" timestamp with time zone
	);

	-- convos
	create table jack.convos (
		"user" text primary key,
		"intent" text,
		"slot" text,
		"state" jsonb not null default '{}'::jsonb,
		"ttl" int
	);

	-- jack add
	create function jack.add(a int, b int) returns int as $$
	begin
		return a + b;
	end;
	$$ language plpgsql immutable strict;
`

	down := `
	drop table if exists jack.convos cascade;

	drop table if exists jack.crons cascade;

	drop table if exists jack.standups_teammates cascade;
	drop type if exists jack.standup_teammate_status cascade;

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

	drop extension if exists citext;
`

	conn, teardown := testutil.Connect(t, url)
	defer teardown()

	testutil.Exec(t, conn, down)
	testutil.Exec(t, conn, up)

	var tests = []struct {
		Name     string
		Setup    string
		Function string
		Expected string
		Error    string
	}{
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Function: `team.FindByID(db, 2)`,
			Expected: `{"id":2,"token":22,"team_name":"b","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Function: `team.Insert(db, team.New().Token(11).TeamName("1"))`,
			Expected: `{"id":1,"token":11,"team_name":"1","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Function: `team.UpdateByID(db, 2, team.New().StripeID("stripey"))`,
			Expected: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Function: `team.UpdateByID(db, 2, team.New().StripeID("stripey").Active(false))`,
			Expected: `{"id":2,"token":22,"team_name":"b","stripe_id":"stripey","free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			`,
			Function: `teammate.FindBySlackID(db, "b")`,
			Expected: `{"id":2,"team_id":1,"slack_id":"b","username":"b","timezone":"b"}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
			`,
			Function: `standup.FindByNameAndTeamID(db, "b", 1)`,
			Expected: `{"id":2,"team_id":1,"name":"b","channel":"b","time":"a","timezone":"a"}`,
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
			Function: `report.Find(db, report.NewFilter().TeammateID(1), report.NewOrder().Timestamp(report.DESC))`,
			Expected: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
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
			Function: `report.FindMany(db, report.NewFilter().Status(enum.ReportStatusAsked))`,
			Expected: `[{"id":1,"teammate_id":1,"standup_id":1,"status":"ASKED","timestamp":1},{"id":2,"teammate_id":1,"standup_id":1,"status":"ASKED","timestamp":2}]`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			`,
			Function: `standupteammate.Insert(db, standupteammate.New().StandupID(1).TeammateID(2).Status(enum.StandupTeammateStatusActive).Time("1:00").Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			`,
			Function: `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().StandupID(1).TeammateID(2).Time("1:00").Status(enum.StandupTeammateStatusActive).Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "time", "status", owner) values (1, 2, '12:00', 'ACTIVE', false);
			`,
			Function: `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Status(enum.StandupTeammateStatusInvited).Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"status":"INVITED","time":"01:00:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 2, 'ACTIVE', '12:00', false);
			`,
			Function: `standupteammate.UpdateByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Owner(true))`,
			Expected: `{"id":1,"standup_id":1,"teammate_id":2,"status":"ACTIVE","time":"01:00:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j2', '* * * * 1-5');
			`,
			Function: `cron.DeleteByJob(db, "j2")`,
			Expected: `{"id":2,"job":"j2","frequency":"* * * * 1-5"}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j2', '* * * * 1-5');
			`,
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
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 1, 'ACTIVE', '12:00', false);
				insert into jack.standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 3, 'ACTIVE', '1:00', true);
			`,
			Function: `standupteammate.FindMany(db, standupteammate.NewFilter().StandupID(1))`,
			Expected: `[{"id":1,"standup_id":1,"teammate_id":1,"status":"ACTIVE","time":"12:00:00"},{"id":2,"standup_id":1,"teammate_id":3,"status":"ACTIVE","time":"01:00:00","owner":true}]`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Function: `cron.Delete(db, cron.NewFilter().JobContains("j1"))`,
			Expected: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Function: `cron.Delete(db, cron.NewFilter().JobStartsWith("j1"))`,
			Expected: `{"id":1,"job":"j1","frequency":"* * * * *"}`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Function: `cron.DeleteMany(db, cron.NewFilter().JobContains("2"))`,
			Expected: `[{"id":2,"job":"j20","frequency":"* * * * 1-5"},{"id":3,"job":"j21","frequency":"* * * * 1-5"}]`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
			`,
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
			Function: `report.Find(db, report.NewFilter().TeammateID(2).StandupID(1).TimestampGt(2), report.NewOrder().Timestamp(report.DESC))`,
			Expected: `{"id":4,"teammate_id":2,"standup_id":1,"status":"COMPLETE","timestamp":4}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
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
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'b', 'b', 'a', 'a');
				insert into jack.standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 1, 'ACTIVE', '12:00', false);
				insert into jack.standups_teammates (standup_id, teammate_id, "status", "time", owner) values (1, 3, 'ACTIVE', '1:00', true);
			`,
			Function: `standupteammate.Find(db, standupteammate.NewFilter().Owner(true).StandupIDIn(1, 3))`,
			Expected: `{"id":2,"standup_id":1,"teammate_id":3,"status":"ACTIVE","time":"01:00:00","owner":true}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
				insert into jack.teams (token, team_name) values (33, 'c');
			`,
			Function: `team.Update(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 44))`,
			Expected: `{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
				insert into jack.teams (token, team_name) values (33, 'c');
			`,
			Function: `team.UpdateMany(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 22))`,
			Expected: `[{"id":1,"token":11,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1},{"id":2,"token":22,"team_name":"cool","active":true,"free_teammates":4,"cost_per_user":1}]`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Function: `cron.UpdateByID(db, 1, cron.New())`,
			Error:    `cron.UpdateByID: no input provided`,
		},
		{
			Setup: `
				insert into jack.crons ("job", "frequency") values ('j1', '* * * * *');
				insert into jack.crons ("job", "frequency") values ('j20', '* * * * 1-5');
				insert into jack.crons ("job", "frequency") values ('j21', '* * * * 1-5');
			`,
			Function: `cron.UpdateByID(db, 1, cron.New().NullableFrequency(nil))`,
			Expected: `{"id":1,"job":"j1"}`,
		},
		{
			Setup: `
				insert into jack.convos ("user", "intent") values ('U0QS7USPJ', 'standup_join');
				insert into jack.convos ("user", "intent") values ('U0QS890N5', 'standup_join');
			`,
			Function: `convo.FindMany(db, convo.NewFilter().UserIn("U0QS7USPJ", "U0QS890N5"))`,
			Expected: `[{"user":"U0QS7USPJ","intent":"standup_join","state":{}},{"user":"U0QS890N5","intent":"standup_join","state":{}}]`,
		},
		// TODO: move elsewhere
		// {
		// 	Function: `convo.New().User("U0QS7USPJ").Intent("standup_join").Insert(db)`,
		// 	Expected: `{"user":"U0QS7USPJ","intent":"standup_join","state":{}}`,
		// },
		// {
		// 	Setup: `
		// 		insert into jack.teams (token, team_name) values (11, 'a');
		// 		insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
		// 		insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
		// 	`,
		// 	Function: `standupteammate.New().StandupID(1).TeammateID(1).Status(enum.StandupTeammateStatusActive).Time("1:00").Owner(true).Insert(db)`,
		// 	Expected: `{"id":1,"standup_id":1,"teammate_id":1,"status":"ACTIVE","time":"01:00:00","owner":true}`,
		// },
		{
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
			`,
			Function: `question.InsertMany(
				db,
				question.New().StandupID(1).Order(2).Question("what's my name?"),
				question.New().StandupID(1).Order(1).Question("what's my age?"),
			)`,
			Expected: `[{"id":1,"standup_id":1,"order":2,"question":"what's my name?"},{"id":2,"standup_id":1,"order":1,"question":"what's my age?"}]`,
		},
		{
			Setup: `
				insert into jack.events ("time") values ('2018-09-04 00:00:00+00');
			`,
			Function: `event.Find(db, event.NewFilter().TimeLte(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
			Expected: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
		},
		{
			Setup: `
				insert into jack.events ("time") values ('2018-09-04 00:00:00+00');
			`,
			Function: `event.Find(db, event.NewFilter().TimeLt(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
			Expected: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
		},
		{
			Setup: `
				insert into jack.events ("time") values ('2018-09-04 00:00:00+00');
			`,
			Function: `event.Find(db, event.NewFilter().TimeGte(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))`,
			Expected: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
		},
		{
			Setup: `
				insert into jack.events ("time") values ('2018-09-04 00:00:00+00');
			`,
			Function: `event.Find(db, event.NewFilter().TimeGt(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))`,
			Expected: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
		},
		{
			Setup: `
				insert into jack.events ("time") values ('2018-09-04 00:00:00+00');
			`,
			Function: `event.Find(db, event.NewFilter().Time(time.Date(2018, 9, 4, 0, 0, 0, 0, time.UTC)))`,
			Expected: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
		},
		{
			Setup: `
				insert into jack.events ("time") values ('2018-09-04 00:00:00+00');
			`,
			Function: `event.Find(db, event.NewFilter().TimeNot(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))`,
			Expected: `{"id":1,"time":"2018-09-04T07:00:00+07:00"}`,
		},
		{
			Setup: `
				insert into jack.events ("time") values (NULL);
				insert into jack.events ("time") values (NULL);
				insert into jack.events ("time") values (NULL);
			`,
			Function: `event.FindMany(db, event.NewFilter().NullableTime(nil))`,
			Expected: `[{"id":1},{"id":2},{"id":3}]`,
		},
		{
			Setup: `
				insert into jack.events ("time") values (NULL);
				insert into jack.events ("time") values (NULL);
				insert into jack.events ("time") values (NULL);
			`,
			Function: `event.FindMany(db, event.NewFilter().NullableTime(&now))`,
			Expected: `[]`,
		},
		{
			Name: "empty_in_finds_nothing",
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teams (token, team_name) values (22, 'b');
			`,
			Function: `team.FindMany(db, team.NewFilter().IDIn().TokenIn(11, 22))`,
			Expected: `[]`,
		},
		{
			Name: "nullable_fk",
			Setup: `
				insert into jack.teams (token, team_name) values (11, 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'a', 'a', 'a');
				insert into jack.teammates (team_id, slack_id, username, timezone) values (1, 'b', 'b', 'b');
				insert into jack.standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
				insert into jack.posts (standup_id) values (1);
				insert into jack.posts (standup_id) values (1);
				insert into jack.reports (teammate_id, standup_id, post_id) values (1, 1, 1);
				insert into jack.reports (teammate_id, standup_id, post_id) values (1, 1, 2);
				insert into jack.reports (teammate_id, standup_id, status) values (1, 1, 'COMPLETE');
				insert into jack.reports (teammate_id, standup_id, status) values (2, 1, 'COMPLETE');
			`,
			Function: `report.Find(db, report.NewFilter().NullablePostID(nil).TeammateID(1))`,
			Expected: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
		},
		{
			Setup: `
				insert into jack.teams (token, email, team_name) values (11, 'mattmuelle@gmail.com', 'a');
			`,
			Function: `team.Find(db, team.NewFilter().Email("maTTMuelle@gmail.com"))`,
			Expected: `{"id":1,"token":11,"team_name":"a","email":"mattmuelle@gmail.com","active":true,"free_teammates":4,"cost_per_user":1}`,
		},
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
	defer func() {
		if !t.Failed() {
			cleanup()
		}
	}()

	for _, test := range tests {
		name := test.Name
		if name == "" {
			name = truncate(test.Function, 20)
			if i := strings.Index(test.Function, "("); i >= 0 {
				name = truncate(test.Function, i)
			}
		}

		// run the functions
		t.Run(name+"/function", func(t *testing.T) {
			testutil.Exec(t, conn, down)
			testutil.Exec(t, conn, up)

			if test.Setup != "" {
				testutil.Exec(t, conn, test.Setup)
			}

			stdout, stderr, remove := testutil.Run(t, name, `
				package main

				import (
					"time"

					"`+importBase+`/pogo/enum"
					pogo "`+importBase+`/pogo"
					team "`+importBase+`/pogo/team"
					cron "`+importBase+`/pogo/cron"
					report "`+importBase+`/pogo/report"
					standup "`+importBase+`/pogo/standup"
					question "`+importBase+`/pogo/question"
					teammate "`+importBase+`/pogo/teammate"
					standupteammate "`+importBase+`/pogo/standup-teammate"
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

					actual, err := `+test.Function+`
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
				if test.Error != "" {
					if test.Error == stderr {
						return
					}
					fmt.Println("# Expected:")
					fmt.Println(test.Error)
					fmt.Println()
					fmt.Println("# Actual:")
					fmt.Println(stderr)
					fmt.Println()
					t.Fatal(diff(test.Error, stderr))
				}
				t.Fatal(errors.New(stderr))
			}

			if test.Expected != stdout {
				fmt.Println("# Expected:")
				fmt.Println(test.Expected)
				fmt.Println()
				fmt.Println("# Actual:")
				fmt.Println(stdout)
				fmt.Println()
				t.Fatal(diff(test.Expected, stdout))
			}

			remove()
		})

		// run the models
		parts := strings.SplitN(test.Function, ".", 2)
		model := parts[0]
		call := parts[1]

		// run the functions
		t.Run(name+"/model", func(t *testing.T) {
			testutil.Exec(t, conn, down)
			testutil.Exec(t, conn, up)

			if test.Setup != "" {
				testutil.Exec(t, conn, test.Setup)
			}

			stdout, stderr, remove := testutil.Run(t, name, `
				package main

				import (
					"time"

					"`+importBase+`/pogo/enum"
					pogo "`+importBase+`/pogo"
					team "`+importBase+`/pogo/team"
					cron "`+importBase+`/pogo/cron"
					report "`+importBase+`/pogo/report"
					standup "`+importBase+`/pogo/standup"
					question "`+importBase+`/pogo/question"
					teammate "`+importBase+`/pogo/teammate"
					standupteammate "`+importBase+`/pogo/standup-teammate"
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

					var model `+model+`.Model
					actual, err := model.`+call+`
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
				if test.Error != "" {
					if test.Error == stderr {
						return
					}
					fmt.Println("# Expected:")
					fmt.Println(test.Error)
					fmt.Println()
					fmt.Println("# Actual:")
					fmt.Println(stderr)
					fmt.Println()
					t.Fatal(diff(test.Error, stderr))
				}
				t.Fatal(errors.New(stderr))
			}

			if test.Expected != stdout {
				fmt.Println("# Expected:")
				fmt.Println(test.Expected)
				fmt.Println()
				fmt.Println("# Actual:")
				fmt.Println(stdout)
				fmt.Println()
				t.Fatal(diff(test.Expected, stdout))
			}

			remove()
		})
	}
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
