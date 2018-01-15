package models

type Group struct {
	Id           int
	Teams        []Team
	TimeTable    []Timetable
	GroupStageId int
}
