package services

import (
	"log"
	"toss-up/models"
	"database/sql"
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
		for j := i; j < len(teams)-1; j++ {
			timeTable = append(timeTable, models.Timetable{GroupId: group.Id, Match: teams[i].Name + " - " + teams[j+1].Name})
		}
	}
	return timeTable
}

func FindResultsByGroupStageId(groupStageId int) ([]string, error) {
	var results []string
	rows, err := db.Query(`SELECT result FROM groups 
				JOIN timetables t ON groups.id = t.group_id
				WHERE group_stage_id = $1`, groupStageId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result string
		err := rows.Scan(&result)
		if err == nil {
			results = append(results, result)
		} else {
			return nil, err
		}
	}

	if len(results) == 0 {
		return nil, err
	} else {
		return results, err
	}
}

func FindTimetableCurrentGroupStage() ([]models.Timetable, error) {
	stage, err := FindGroupStageByIsFinishedFalse()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if stage.Name == "" {
		return nil, err
	}
	rows, err := db.Query(`SELECT t.id,t.group_id, t.match, t.result FROM groups
	  					JOIN timetables t ON groups.id = t.group_id
	  					WHERE group_stage_id = $1`, stage.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var timetables []models.Timetable
	for rows.Next() {
		var timetable models.Timetable
		err := rows.Scan(&timetable.Id, &timetable.GroupId, &timetable.Match, &timetable.Result)
		if err == nil {
			timetables = append(timetables, timetable)
		} else {
			return nil, err
		}
	}

	if len(timetables) == 0 {
		return nil, err
	} else {
		return timetables, err
	}
}

func UpdateResultMatch(timetableId int, result sql.NullString) (models.Timetable, error) {
	timetable, err := FindTimetableById(timetableId)
	if timetable.Match == "" {
		return timetable, err
	}
	_, err = db.Exec(`UPDATE timetables SET result = $1 WHERE id = $2`, result, timetableId)
	if err != nil {
		log.Println(err)
		return timetable, err
	}

	timetable.Result = result
	return timetable, err
}

func FindTimetableById(timetableId int) (models.Timetable, error) {
	var dbTimetable models.Timetable
	err := db.QueryRow(`SELECT * FROM timetables WHERE id = $1`, timetableId).Scan(&dbTimetable.Id, &dbTimetable.Match,
		&dbTimetable.GroupId, &dbTimetable.Result)
	if err != nil {
		log.Println(err)
		return dbTimetable, err
	}
	return dbTimetable, err
}
