package models

type Group struct {
	Id           int
	Teams        []Team
	TimeTable    []string
	GroupStageId int
}
