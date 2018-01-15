package models

import (
	"time"
)

type GroupStage struct {
	Id         int
	DateStart  time.Time
	DateEnd    time.Time
	Name       string
	IsFinished bool
	Groups     []Group
}
