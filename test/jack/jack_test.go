package jack_test

import (
	"bytes"
	"errors"
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"testing"

	text "github.com/matthewmueller/go-text"
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

		team_owner bool not null default 'false'
	);

	-- jack add
	create function jack.add(a int, b int) returns int as $$
	begin
		return a + b;
	end;
	$$ language plpgsql immutable strict;
`

var down = `
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
			Function: `report.Find(db, report.NewFilter().TeammateID(1), report.NewOrderBy().Timestamp(report.Desc))`,
			Expected: `{"id":3,"teammate_id":1,"standup_id":1,"status":"COMPLETE","timestamp":3}`,
		},
		// {
		// 	Query:    `UPDATE reports SET questions = $1 WHERE id = $2 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO reports (user_id, timestamp, standup_id, status) VALUES ($1, $2, $3, $4) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO teammate_standups (standup_id, teammate_id, time) VALUES ($1, $2, $3) ON CONFLICT (standup_id, teammate_id) DO UPDATE SET time = $3, status = 'ACTIVE' RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO teammate_standups (standup_id, teammate_id) VALUES ($1, $2) ON CONFLICT (standup_id, teammate_id) DO UPDATE SET status = 'ACTIVE' RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE teammate_standups SET status = 'OPTOUT' WHERE teammate_id = $1 AND standup_id = $2 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `DELETE FROM crons WHERE job = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM standups WHERE id = $1 AND team_id = $2`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT id, team_name, bot_access_token, team_access_token, bot_user_id FROM teams WHERE team_id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teammates WHERE id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teammates WHERE slack_id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE standups SET name = $1, channel_id = $2, time = $3, tz = $4, team_id = $5, owner = $6 WHERE id = $7 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO teammates (slack_id, team_id) VALUES ($1, $2) ON CONFLICT (slack_id) DO UPDATE SET id = teammates.id RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE teams SET active = TRUE WHERE team_id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM standups WHERE team_id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO crons (job, frequency, tz) VALUES ($1, $2, $3) ON CONFLICT (job) DO NOTHING RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teammate_standups WHERE standup_id = $1 AND status = $2`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teams where id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO crons (job, frequency, tz) VALUES ($1, $2, $3) ON CONFLICT (job) DO NOTHING`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE teams SET active = FALSE WHERE team_id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teammate_standups WHERE standup_id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `DELETE FROM crons WHERE job SIMILAR TO $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO teams (team_id, team_access_token, bot_access_token, bot_user_id, team_name, scope, email, trial_ends, active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (team_id) DO UPDATE SET team_access_token = $2, bot_access_token = $3, bot_user_id = $4, team_name = $5, scope = $6, email = $7, active = $9 WHERE teams.id = teams.id RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM standups WHERE team_id = $1 AND id = $2 LIMIT 1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM crons WHERE job LIKE $2`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM standups WHERE id = $1 LIMIT 1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT teammate_id FROM teammate_standups WHERE standup_id = $1 AND status = $2`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teammates WHERE id = ANY ($1)`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT bot_access_token,team_access_token FROM teams WHERE team_id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT bot_access_token FROM teams WHERE team_id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM conversations WHERE user_id = $1 ORDER BY updated DESC LIMIT 1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM reports WHERE user_id = $1 AND standup_id = $2 AND "timestamp" > (timestamp '1d' - INTERVAL '1hr') ORDER BY "timestamp" DESC LIMIT 1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teams`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM teammate_standups WHERE status = 'ACTIVE' AND standup_id = ANY($1)`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE teams SET stripe_id = $1 WHERE id = $2 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE teams SET stripe_id = $1 WHERE id = $2 RETURNING team_id`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO invitations ("from", "to", standup, "for") VALUES ($1, $2, $3, $4) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO reports (user_id, timestamp, standup_id) VALUES ($1, $2, $3) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO reports (user_id, timestamp, standup_id) VALUES ($1, $3, $2) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE teams SET team_access_token = $1, bot_access_token = $2, bot_user_id = $3, team_name = $4, scope = $5, email = $6, active = TRUE WHERE team_id = $7 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO teams (team_id, team_access_token, bot_access_token, bot_user_id, team_name, scope, email, trial_ends) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO conversations (user_id, topic, context, state) VALUES ($1, $2, $2, $3) RETURNING id`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE conversations SET context = $1, state = $2 WHERE id = $3`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE conversations SET context = NULL WHERE id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT questions FROM reports WHERE id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE reports SET status = $2, questions = $3 WHERE id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE reports SET questions = $2 WHERE id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO standups (name, channel_id, time, tz, team_id, owner) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO teammate_standups (time, teammate_id, standup_id) VALUES ($1, $2, $3) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO teammates (slack_id, team_id) VALUES ($1, $2) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE teams SET active = FALSE WHERE id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `DELETE FROM teammate_standups WHERE teammate_id = $1 AND standup_id = $2`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `DELETE FROM teammate_standups WHERE teammate_id = $1 AND standup_id = $2 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM conversations WHERE user_id = $1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE conversations SET state = $2 WHERE id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE invitations SET status = $2 WHERE id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE standups SET name = $2, channel_id = $3, time = $4, tz = $5, team_id = $6 WHERE id = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM invitations WHERE "to" = $1 AND status = $2 AND "for" = $3 LIMIT 1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `SELECT * FROM jack.invitations WHERE "to" = $1 AND status = $2 AND "for" = $3 LIMIT 1`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO jack.teammates (slack_id, team_id) VALUES ($1, $2) RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `UPDATE jack.conversations SET state = $2 WHERE id = $1 RETURNING id`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `DELETE FROM jack.crons WHERE job = $1 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO crons (job, frequency, tz) VALUES ($1, $2, $3) ON CONFLICT (job) DO UPDATE SET frequency = concat($4::text, ' ', substring(crons.frequency from '[\\d\\-\\,\\*]+$')), tz = $3 RETURNING *`,
		// 	Expected: ``,
		// },
		// {
		// 	Query:    `INSERT INTO crons (job, frequency, tz) VALUES ($1, $2, $3) ON CONFLICT (job) DO UPDATE SET frequency = concat(' ', $4::text, substring(crons.frequency from '[\\d\\-\\,\\*]+$')), tz = $3 RETURNING *`,
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

	cleanup := testutil.Build(t, url, "jack", "./pogo")
	_ = cleanup
	// defer remove()

	for _, test := range tests {
		t.Run(truncate(text.Snake(test.Function), 20), func(t *testing.T) {

			testutil.Exec(t, conn, down)
			testutil.Exec(t, conn, up)

			if test.Setup != "" {
				testutil.Exec(t, conn, test.Setup)
			}

			stdout, stderr, remove := testutil.Run(t, text.Snake(test.Function), `
				package main
	
				import (
					"`+importBase+`/pogo/enum"
					pogo "`+importBase+`/pogo"
					team "`+importBase+`/pogo/team"
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
					t.Fatal(diff(test.Error, stderr))
					return
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
		if num > 3 {
			num -= 3
		}
		s = str[0:num]
	}
	return s
}
