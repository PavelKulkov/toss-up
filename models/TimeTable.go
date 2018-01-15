package models

import "database/sql"

type Timetable struct {
	Id      int
	Match   string
	Result sql.NullString
	GroupId sql.NullInt64
}
