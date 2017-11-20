begin;
drop table if exists jack.standups_teammates cascade;
drop trigger if exists created_at on jack.standups_teammates;
drop trigger if exists updated_at on jack.standups_teammates;

drop table if exists jack.reports cascade;
drop type if exists jack.report_status;
drop trigger if exists created_at on jack.reports;
drop trigger if exists updated_at on jack.reports;

drop table if exists jack.standups cascade;
drop trigger if exists created_at on jack.standups;
drop trigger if exists updated_at on jack.standups;

drop table if exists jack.teammates cascade;
drop trigger if exists created_at on jack.teammates;
drop trigger if exists updated_at on jack.teammates;

drop table if exists jack.teams cascade;
drop trigger if exists created_at on jack.teams;
drop trigger if exists updated_at on jack.teams;

drop table if exists jack."team-invites" cascade;
drop trigger if exists created_at on jack.teams;
drop trigger if exists updated_at on jack.teams;

drop function if exists jack.set_created_at();
drop function if exists jack.set_updated_at();
drop extension if exists pgcrypto cascade;

drop schema if exists jack cascade;
commit;