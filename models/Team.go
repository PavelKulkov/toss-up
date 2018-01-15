package models

import "database/sql"

type Team struct {
	Id          int
	Name        string
	Description string
	GroupId     sql.NullInt64
}
