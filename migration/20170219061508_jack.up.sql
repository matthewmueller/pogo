-- Schema

create schema jack;

-- EXTENSIONS

create extension if not exists pgcrypto;

-- TRIGGERS

create function jack.set_created_at() returns trigger as $$
begin
  -- We will let the inserter manually set a `created_at` time if they desire.
  if (new.created_at is null) then
    new.created_at := current_timestamp;
  end if;
  return new;
end;
$$ language plpgsql;

create function jack.set_updated_at() returns trigger as $$
begin
  new.updated_at := current_timestamp;
  return new;
end;
$$ language plpgsql;

-- TEAMS

-- add all the other columns
create table if not exists jack.teams (
  id uuid primary key not null default gen_random_uuid(),
  slack_team_id text unique not null,
  slack_team_access_token text unique not null,
  slack_bot_access_token text unique not null,
  slack_bot_id text not null,
  team_name text not null,
  scope text[] not null default '{}',
  email text,
  stripe_id text,
  active boolean not null default true,
  free_teammates integer not null default 4,
  cost_per_user integer not null default 1,
  trial_ends timestamp with time zone not null default (now() at time zone 'utc' + interval '14 days'),
  created_at timestamp with time zone default (now() at time zone 'utc'),
  updated_at timestamp with time zone default (now() at time zone 'utc')
);

create trigger created_at before insert on jack.teams for each row execute procedure jack.set_created_at();
create trigger updated_at before update on jack.teams for each row execute procedure jack.set_updated_at();

-- add all the other columns
create table if not exists jack.teammates (
  id uuid primary key not null default gen_random_uuid(),
  slack_id text unique not null,
  username text not null,
  first_name text,
  last_name text,
  email text,
  avatar text,
  timezone text not null default 'America/Los_Angeles',
  created_at timestamp with time zone default (now() at time zone 'utc'),
  updated_at timestamp with time zone default (now() at time zone 'utc')
);

create trigger created_at before insert on jack.teammates for each row execute procedure jack.set_created_at();
create trigger updated_at before update on jack.teammates for each row execute procedure jack.set_updated_at();

-- STANDUPS

create table if not exists jack.standups (
  id uuid primary key not null default gen_random_uuid(),
  name text not null,
  slack_channel_id text unique not null,
  "time" text not null default '11:00:00',
  timezone text not null default 'America/Los_Angeles',
  questions json not null default '["What did you get done yesterday?","What are you working on right now?","When do you think you''ll be done with that?","Is there anything that''s blocking you?"]',
  team_id uuid not null references jack.teams(id) on delete cascade,
  created_at timestamp with time zone default (now() at time zone 'utc'),
  updated_at timestamp with time zone default (now() at time zone 'utc')
);

create trigger created_at before insert on jack.standups for each row execute procedure jack.set_created_at();
create trigger updated_at before update on jack.standups for each row execute procedure jack.set_updated_at();

-- REPORTS

create type jack.report_status as enum ('ASKED', 'SKIP', 'COMPLETE');

-- add all the other columns
create table if not exists jack.reports (
  id uuid primary key not null default gen_random_uuid(),
  user_id uuid not null references jack.teammates(id) on delete cascade,
  "timestamp" timestamp with time zone not null,
  questions json not null default '{}',
  standup_id uuid not null references jack.standups(id) on delete cascade,
  status jack.report_status not null default 'ASKED',
  created_at timestamp with time zone default (now() at time zone 'utc'),
  updated_at timestamp with time zone default (now() at time zone 'utc')
);

--
-- created_at/update_at Triggers
--

create trigger created_at before insert on jack.reports for each row execute procedure jack.set_created_at();
create trigger updated_at before update on jack.reports for each row execute procedure jack.set_updated_at();

-- STANDUPS_TEAMMATES

-- add all the other columns
create table if not exists jack.standups_teammates (
  standup_id uuid not null references jack.standups(id),
  teammate_id uuid not null references jack.teammates(id),
  team_owner bool not null default 'false',
  created_at timestamp with time zone default (now() at time zone 'utc'),
  updated_at timestamp with time zone default (now() at time zone 'utc'),
  unique(teammate_id, standup_id)
);

--
-- created_at/update_at Triggers
--

create trigger created_at before insert on jack.standups_teammates for each row execute procedure jack.set_created_at();
create trigger updated_at before update on jack.standups_teammates for each row execute procedure jack.set_updated_at();

create table if not exists jack."team-invites" (
  id uuid primary key not null default gen_random_uuid(),
  emails text[] not null default '{}'
)