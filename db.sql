create table teams
(
	id serial not null
		constraint teams_pkey
			primary key,
	name varchar(255) not null,
	description varchar(255) not null,
	group_id integer
)
;

create table groups
(
	id serial not null
		constraint groups_pkey
			primary key,
	group_stage_id integer
)
;

alter table teams
	add constraint teams_groups_id_fk
		foreign key (group_id) references groups
;

create table group_stages
(
	id serial not null
		constraint group_stages_pkey
			primary key,
	date_start timestamp not null,
	date_end timestamp not null,
	name varchar(255) not null,
	is_finished boolean default false not null
)
;

alter table groups
	add constraint groups_group_stages_id_fk
		foreign key (group_stage_id) references group_stages
;

create table timetables
(
	id serial not null
		constraint timetable_pkey
			primary key,
	match varchar(255) not null,
	group_id integer
		constraint timetable_groups_id_fk
			references groups,
	result varchar(255)
)
;

