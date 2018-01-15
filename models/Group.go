package models

import "database/sql"

type Group struct {
	Id           sql.NullInt64
	Teams        []Team
	TimeTable    []Timetable
	GroupStageId int
}
