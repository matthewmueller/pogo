package tests

// Postgres tests
var Postgres = []*Test{
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		Schema: `

		`,
		QueryGo: `team.FindByID(db, 2)`,
		Expect:  `{"active":true,"cost_per_user":1,"free_teammates":4,"id":2,"team_name":"b","token":22}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.Insert(db, team.New().Token(11).TeamName("1"))`,
		Expect:  `{"active":true,"cost_per_user":1,"free_teammates":4,"id":1,"team_name":"1","token":11}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.UpdateByID(db, 2, team.New().StripeID("stripey"))`,
		Expect:  `{"active":true,"cost_per_user":1,"free_teammates":4,"id":2,"stripe_id":"stripey","team_name":"b","token":22}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.UpdateByID(db, 2, team.New().StripeID("stripey").Active(false))`,
		Expect:  `{"cost_per_user":1,"free_teammates":4,"id":2,"stripe_id":"stripey","team_name":"b","token":22}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
		Down: `
			drop table if exists teammates cascade;
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `teammate.FindBySlackID(db, "b")`,
		Expect:  `{"id":2,"slack_id":"b","team_id":1,"timezone":"b","username":"b"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standup.FindByNameAndTeamID(db, "b", 1)`,
		Expect:  `{"channel":"b","id":2,"name":"b","team_id":1,"time":"a","timezone":"a"}`,
	},
	{
		Up: `
			create extension citext;
			create table blogs (
				name text not null
			);
			`,
		Down: `
				drop table if exists blogs cascade;
				drop extension if exists citext cascade;
			`,
		QueryGo: `blog.Find(db)`,
		Error:   `blog not found`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table reports (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		Schema:  ``,
		QueryGo: `report.Find(db, report.NewFilter().TeammateID(1), report.NewOrder().Timestamp(report.DESC))`,
		Expect:  `{"id":3,"standup_id":1,"status":"COMPLETE","teammate_id":1,"timestamp":3}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table reports (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `report.FindMany(db, report.NewFilter().Status(enum.ReportStatusAsked))`,
		Expect:  `[{"id":1,"standup_id":1,"status":"ASKED","teammate_id":1,"timestamp":1},{"id":2,"standup_id":1,"status":"ASKED","teammate_id":1,"timestamp":2}]`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table standups_teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standupteammate.Insert(db, standupteammate.New().StandupID(1).TeammateID(2).Status(enum.StandupTeammateStatusActive).Time("1:00").Owner(true))`,
		Expect:  `{"id":1,"owner":true,"standup_id":1,"status":"ACTIVE","teammate_id":2,"time":"01:00:00"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table standups_teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().StandupID(1).TeammateID(2).Time("1:00").Status(enum.StandupTeammateStatusActive).Owner(true))`,
		Expect:  `{"id":1,"owner":true,"standup_id":1,"status":"ACTIVE","teammate_id":2,"time":"01:00:00"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table standups_teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standupteammate.UpsertByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Status(enum.StandupTeammateStatusInvited).Owner(true))`,
		Expect:  `{"id":1,"owner":true,"standup_id":1,"status":"INVITED","teammate_id":2,"time":"01:00:00"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table standups_teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standupteammate.UpdateByStandupIDAndTeammateID(db, 1, 2, standupteammate.New().Time("1:00").Owner(true))`,
		Expect:  `{"id":1,"owner":true,"standup_id":1,"status":"ACTIVE","teammate_id":2,"time":"01:00:00"}`,
	},
	{
		Up: `
			create extension citext;
			create table crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j2', '* * * * 1-5');
		`,
		Down: `
			drop table if exists crons cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `cron.DeleteByJob(db, "j2")`,
		Expect:  `{"frequency":"* * * * 1-5","id":2,"job":"j2"}`,
	},
	{
		Up: `
			create extension citext;
			create table crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j2', '* * * * 1-5');
		`,
		Down: `
			drop table if exists crons cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `cron.DeleteByID(db, 2)`,
		Expect:  `{"frequency":"* * * * 1-5","id":2,"job":"j2"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `teammate.UpsertBySlackID(db, "b", teammate.New().TeamID(2).Username("b").Timezone("b"))`,
		Expect:  `{"id":2,"slack_id":"b","team_id":2,"timezone":"b","username":"b"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `teammate.Upsert(db, teammate.New().ID(2).TeamID(2).SlackID("b").Username("b").Timezone("b"))`,
		Expect:  `{"id":2,"slack_id":"b","team_id":2,"timezone":"b","username":"b"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `teammate.Upsert(db, teammate.New().TeamID(2).SlackID("c").Username("c").Timezone("c"))`,
		Expect:  `{"id":3,"slack_id":"c","team_id":2,"timezone":"c","username":"c"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standup.FindMany(db, standup.NewFilter().TeamID(2), standup.NewOrder().Channel(standup.DESC))`,
		Expect:  `[{"channel":"b","id":2,"name":"b","team_id":2,"time":"b","timezone":"b"},{"channel":"a","id":1,"name":"a","team_id":2,"time":"a","timezone":"a"}]`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table standups_teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standupteammate.FindMany(db, standupteammate.NewFilter().StandupID(1))`,
		Expect:  `[{"id":1,"standup_id":1,"status":"ACTIVE","teammate_id":1,"time":"12:00:00"},{"id":2,"owner":true,"standup_id":1,"status":"ACTIVE","teammate_id":3,"time":"01:00:00"}]`,
	},
	{
		Up: `
			create extension citext;
			create table crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		Down: `
			drop table if exists crons cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `cron.Delete(db, cron.NewFilter().JobContains("j1"))`,
		Expect:  `{"frequency":"* * * * *","id":1,"job":"j1"}`,
	},
	{
		Up: `
			create extension citext;
			create table crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		Down: `
			drop table if exists crons cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `cron.Delete(db, cron.NewFilter().JobStartsWith("j1"))`,
		Expect:  `{"frequency":"* * * * *","id":1,"job":"j1"}`,
	},
	{
		Up: `
			create extension citext;
			create table crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		Down: `
			drop table if exists crons cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `cron.DeleteMany(db, cron.NewFilter().JobContains("2"))`,
		Expect:  `[{"frequency":"* * * * 1-5","id":2,"job":"j20"},{"frequency":"* * * * 1-5","id":3,"job":"j21"}]`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `teammate.Find(db, teammate.NewFilter().SlackIDIn("b", "c"))`,
		Expect:  `{"id":2,"slack_id":"b","team_id":1,"timezone":"b","username":"b"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table reports (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `report.Find(db, report.NewFilter().TeammateID(2).StandupID(1).TimestampGt(2), report.NewOrder().Timestamp(report.DESC))`,
		Expect:  `{"id":4,"standup_id":1,"status":"COMPLETE","teammate_id":2,"timestamp":4}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.FindMany(db)`,
		Expect:  `[{"active":true,"cost_per_user":1,"free_teammates":4,"id":1,"team_name":"a","token":11},{"active":true,"cost_per_user":1,"free_teammates":4,"id":2,"team_name":"b","token":22}]`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create type standup_teammate_status as enum ('INVITED','ACTIVE','INACTIVE');
			create table standups_teammates (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists standup_teammate_status cascade;
			drop table if exists standups_teammates cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `standupteammate.Find(db, standupteammate.NewFilter().Owner(true).StandupIDIn(1, 3))`,
		Expect:  `{"id":2,"owner":true,"standup_id":1,"status":"ACTIVE","teammate_id":3,"time":"01:00:00"}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.Update(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 44))`,
		Expect:  `{"active":true,"cost_per_user":1,"free_teammates":4,"id":1,"team_name":"cool","token":11}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
				id serial primary key not null,
				token integer unique not null
			);
			insert into teams (token) values (11);
		`,
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.Update(db, team.New().Token(11), team.NewFilter().Token(10))`,
		Error:   `team not found`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.UpdateMany(db, team.New().TeamName("cool"), team.NewFilter().TokenIn(11, 22))`,
		Expect:  `[{"active":true,"cost_per_user":1,"free_teammates":4,"id":1,"team_name":"cool","token":11},{"active":true,"cost_per_user":1,"free_teammates":4,"id":2,"team_name":"cool","token":22}]`,
	},
	{
		Up: `
			create extension citext;
			create table crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		Down: `
			drop table if exists crons cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `cron.UpdateByID(db, 1, cron.New())`,
		Error:   `cron.UpdateByID: no input provided`,
	},
	{
		Up: `
			create extension citext;
			create table crons (
				id serial not null primary key,
				"job" text unique not null,
				frequency text
			);
			insert into crons ("job", "frequency") values ('j1', '* * * * *');
			insert into crons ("job", "frequency") values ('j20', '* * * * 1-5');
			insert into crons ("job", "frequency") values ('j21', '* * * * 1-5');
		`,
		Down: `
			drop table if exists crons cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `cron.UpdateByID(db, 1, cron.New().NullableFrequency(nil))`,
		Expect:  `{"id":1,"job":"j1"}`,
	},
	{
		Up: `
			create extension citext;
			create schema jack;
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
		Down: `
			drop table if exists jack.convos cascade;
			drop schema if exists jack;
			drop extension if exists citext cascade;
		`,
		QueryGo: `convo.FindMany(db, convo.NewFilter().UserIn("U0QS7USPJ", "U0QS890N5"))`,
		Expect:  `[{"intent":"standup_join","state":{},"user":"U0QS7USPJ"},{"intent":"standup_join","state":{},"user":"U0QS890N5"}]`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table questions (
				id serial not null primary key,
				standup_id integer not null references standups(id) on delete cascade,
				"order" smallint not null,
				question text not null,
				unique(standup_id, "order")
			);
			insert into teams (token, team_name) values (11, 'a');
			insert into standups (team_id, "name", channel, "time", timezone) values (1, 'a', 'a', 'a', 'a');
		`,
		Down: `
			drop table if exists teams cascade;
			drop table if exists standups cascade;
			drop table if exists questions cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `question.InsertMany(
			db,
			question.New().StandupID(1).Order(2).Question("what's my name?"),
			question.New().StandupID(1).Order(1).Question("what's my age?"),
		)`,
		Expect: `[{"id":1,"order":2,"question":"what's my name?","standup_id":1},{"id":2,"order":1,"question":"what's my age?","standup_id":1}]`,
	},
	{
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `
			func() (interface{}, error) {
				evt, err := event.Find(db, event.NewFilter().TimeLte(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))
				if err != nil {
					return evt, err
				}
				*evt.Time = evt.Time.In(time.UTC)
				return evt, nil
			}()
		`,
		Expect: `{"id":1,"time":"2018-09-04T00:00:00Z"}`,
	},
	{
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `
			func() (interface{}, error) {
				evt, err := event.Find(db, event.NewFilter().TimeLt(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))
				if err != nil {
					return evt, err
				}
				*evt.Time = evt.Time.In(time.UTC)
				return evt, nil
			}()
		`,
		Expect: `{"id":1,"time":"2018-09-04T00:00:00Z"}`,
	},
	{
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00+00');
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `
			func() (interface{}, error) {
				evt, err := event.Find(db, event.NewFilter().TimeGte(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))
				if err != nil {
					return evt, err
				}
				*evt.Time = evt.Time.In(time.UTC)
				return evt, nil
			}()
		`,
		Expect: `{"id":1,"time":"2018-09-04T00:00:00Z"}`,
	},
	{
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00Z');
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `
			func() (interface{}, error) {
				evt, err := event.Find(db, event.NewFilter().TimeGt(time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)))
				if err != nil {
					return evt, err
				}
				*evt.Time = evt.Time.In(time.UTC)
				return evt, nil
			}()
		`,
		Expect: `{"id":1,"time":"2018-09-04T00:00:00Z"}`,
	},
	{
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00Z');
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `
			func() (interface{}, error) {
				evt, err := event.Find(db, event.NewFilter().Time(time.Date(2018, 9, 4, 0, 0, 0, 0, time.UTC)))
				if err != nil {
					return evt, err
				}
				*evt.Time = evt.Time.In(time.UTC)
				return evt, nil
			}()
		`,
		Expect: `{"id":1,"time":"2018-09-04T00:00:00Z"}`,
	},
	{
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values ('2018-09-04 00:00:00Z');
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `
		func() (interface{}, error) {
			evt, err := event.Find(db, event.NewFilter().TimeNot(time.Date(2018, 9, 5, 0, 0, 0, 0, time.UTC)))
			if err != nil {
				return evt, err
			}
			*evt.Time = evt.Time.In(time.UTC)
			return evt, nil
		}()
	`,
		Expect: `{"id":1,"time":"2018-09-04T00:00:00Z"}`,
	},
	{
		Title: "nullable_time_nil",
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `event.FindMany(db, event.NewFilter().NullableTime(nil))`,
		Expect:  `[{"id":1},{"id":2},{"id":3}]`,
	},
	{
		Title: "nullable_time_ok",
		Up: `
			create extension citext;
			create table events (
				id serial not null primary key,
				"time" timestamp with time zone
			);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
			insert into events ("time") values (NULL);
		`,
		Down: `
			drop table if exists events cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `event.FindMany(db, event.NewFilter().NullableTime(&now))`,
		Expect:  `[]`,
	},
	{
		Title: "empty_in_finds_nothing",
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.FindMany(db, team.NewFilter().IDIn().TokenIn(11, 22))`,
		Expect:  `[]`,
	},
	{
		Title: "nullable_fk",
		Up: `
			create extension citext;
			create table teams (
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
			create table teammates (
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
			create table standups (
				id serial primary key not null,
				team_id integer not null references teams(id) on delete cascade,
				"name" text not null,
				channel text unique not null,
				"time" text not null,
				timezone text not null,
				unique(team_id, "name")
			);
			create table posts (
				id serial primary key,
				standup_id integer not null references standups(id) on delete cascade on update cascade,
				created_at timestamp with time zone default (now() at time zone 'utc')
			);
			create type report_status as enum ('ASKED', 'SKIP', 'COMPLETE');
			create table reports (
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
		Down: `
			drop table if exists teams cascade;
			drop table if exists teammates cascade;
			drop table if exists standups cascade;
			drop type if exists report_status cascade;
			drop table if exists reports cascade;
			drop table if exists posts cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `report.Find(db, report.NewFilter().NullablePostID(nil).TeammateID(1))`,
		Expect:  `{"id":3,"standup_id":1,"status":"COMPLETE","teammate_id":1,"timestamp":3}`,
	},
	{
		Up: `
			create extension citext;
			create table teams (
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
		Down: `
			drop table if exists teams cascade;
			drop extension if exists citext cascade;
		`,
		QueryGo: `team.Find(db, team.NewFilter().Email("maTTMuelle@gmail.com"))`,
		Expect:  `{"active":true,"cost_per_user":1,"email":"mattmuelle@gmail.com","free_teammates":4,"id":1,"team_name":"a","token":11}`,
	},
	{
		Up: `
			create table exercises (
				id serial primary key not null,
				distance decimal(5, 3) not null
			);
			insert into exercises (distance) values (12.213);
		`,
		Down: `
			drop table if exists exercises cascade;
		`,
		QueryGo: `exercise.Find(db, exercise.NewFilter().Distance(12.213))`,
		Expect:  `{"distance":12.213,"id":1}`,
	},
	{
		Up: `
			create table exercises (
				id serial primary key not null,
				distance decimal(5, 3) not null
			);
		`,
		Down: `
			drop table if exists exercises cascade;
		`,
		QueryGo: `exercise.Insert(db, exercise.New().Distance(12.213))`,
		Expect:  `{"distance":12.213,"id":1}`,
	},
	{
		Up: `
			create table migrate (
				version bigint not null primary key
			);
		`,
		Down: `
			drop table if exists migrate cascade;
		`,
		QueryGo: `migrate.Insert(db, migrate.New().Version(1))`,
		Expect:  `{"version":1}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.FindByKeyAndName(db, "b", "a")`,
		Expect:  `{"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.FindByEmailAndValue(db, "d", "c")`,
		Expect:  `{"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.UpdateByKeyAndName(db, "b", "a", variable.New().Email("e"))`,
		Expect:  `{"email":"e","key":"b","name":"a","value":"c"}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.UpdateByEmailAndValue(db, "d", "c", variable.New().Name("e"))`,
		Expect:  `{"email":"d","key":"b","name":"e","value":"c"}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.DeleteByKeyAndName(db, "b", "a")`,
		Expect:  `{"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key),
				unique(value, email)
			);
			insert into variables (name, key, value, email) values ('a', 'b', 'c', 'd');
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.DeleteByEmailAndValue(db, "d", "c")`,
		Expect:  `{"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.Upsert(db, variable.New().Key("a").Name("b").Value("c").Email("d"))`,
		Expect:  `{"email":"d","key":"a","name":"b","value":"c"}`,
	},
	{
		Up: `
			create table variables (
				name text not null,
				key text not null,
				value text not null,
				email text not null,
				primary key(name, key)
			);
		`,
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.Upsert(db, variable.New().Key("a").Name("e").Value("c").Email("d"))`,
		Expect:  `{"email":"d","key":"a","name":"e","value":"c"}`,
	},
	{
		Up: `
			create table variables (
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
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.UpsertByEmailAndValue(db, "d", "c", variable.New().Name("a").Key("b").Count(1))`,
		Expect:  `{"count":1,"email":"d","key":"b","name":"a","value":"c"}`,
	},
	{
		Up: `
			create table variables (
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
		Down: `
			drop table if exists variables cascade;
		`,
		QueryGo: `variable.UpsertByEmailAndValue(db, "h", "c", variable.New().Name("d").Key("b").Count(1))`,
		Expect:  `{"count":1,"email":"h","key":"b","name":"d","value":"c"}`,
	},
	{
		Up: `
			create table big (
				id serial primary key not null,
				large bigint not null
			);
		`,
		Down: `
			drop table if exists big cascade;
		`,
		QueryGo: `big.Insert(db, big.New().Large(int64(123)))`,
		Expect:  `{"id":1,"large":123}`,
	},
	{
		Up: `
			create table big (
				id serial primary key not null,
				large bigint not null
			);
			insert into big (large) values (123);
		`,
		Down: `
			drop table if exists big cascade;
		`,
		QueryGo: `big.FindByID(db, 1)`,
		Expect:  `{"id":1,"large":123}`,
	},
	{
		Up: `
			create table intabbrev (
				id serial primary key not null,
				n int not null
			);
		`,
		Down: `
			drop table if exists intabbrev cascade;
		`,
		QueryGo: `intabbrev.Insert(db, intabbrev.New().N(123))`,
		Expect:  `{"id":1,"n":123}`,
	},
	{
		Up: `
			create table intabbrev (
				id serial primary key not null,
				n int not null
			);
			insert into intabbrev (n) values (123);
		`,
		Down: `
			drop table if exists intabbrev cascade;
		`,
		QueryGo: `intabbrev.FindByID(db, 1)`,
		Expect:  `{"id":1,"n":123}`,
	},
	{
		Up: `
			create table orders (
				id serial primary key not null,
				n int not null
			);
			insert into orders (n) values (123);
		`,
		Down: `
			drop table if exists orders cascade;
		`,
		QueryGo: `order.FindByID(db, 1)`,
		Expect:  `{"id":1,"n":123}`,
	},
	{
		Up: `
		create table "public"."a" (
			one integer not null,
			two integer not null,
			constraint "a_pkey" primary key ("one", "two")
		);

		create table "public"."b" (
			one integer not null,
			two integer not null,
			constraint "b_one_fkey" foreign key ("one", "two") references "public"."a" ("one", "two")
		);
		`,
		Down: `
			drop table if exists a cascade;
			drop table if exists b cascade;
		`,
		QueryGo: `b.FindByID(db, 1)`,
		Expect:  ``,
	},
}
