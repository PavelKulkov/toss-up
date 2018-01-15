package services

import (
	"log"
	"toss-up/models"
)

func SaveTimetables(timetables []models.Timetable) ([]models.Timetable, error) {
	var err error
	for e := range timetables {
		err := db.QueryRow(`INSERT INTO timetables (match, group_id) VALUES($1,$2) RETURNING id`,
			timetables[e].Match, timetables[e].GroupId).Scan(&timetables[e].Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return timetables, err
}

func GenerateTimeTable(group models.Group) []models.Timetable {
	var timeTable []models.Timetable
	teams := group.Teams
	for i := 0; i < len(teams)-1; i++ {
		for j:=i;j< len(teams)-1; j++  {
			timeTable = append(timeTable, models.Timetable{GroupId: group.Id, Match: teams[i].Name + " - " + teams[j+1].Name})
		}
	}
	return timeTable
}