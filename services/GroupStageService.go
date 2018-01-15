package services

import (
	"log"
	"toss-up/models"
	"time"
	"errors"
)

func CreateGroupStage() (models.GroupStage, error) {
	teams := getNewTeams()
	if len(teams) < 3 {
		return models.GroupStage{}, errors.New("Недостаточно команд для генерации групповой стадии")
	}
	stage, err := saveGroupStage(models.GroupStage{DateStart: time.Now(),
		Name: "Group Stage 1", IsFinished: false})
	countOfGroups, err := GenerateGroups(len(teams))
	distributeTeams := DistributeTeams(countOfGroups, teams, stage.Id)
	groups, err := SaveGroups(distributeTeams)
	stage.Groups = groups
	for e := range groups {
		for e1 := range groups[e].Teams {
			groups[e].Teams[e1].GroupId = groups[e].Id
			UpdateTeam(groups[e].Teams[e1].Id, groups[e].Teams[e1])
		}
	}
	return stage, err
}

//TODO Со Scan какая-то дичь, нужно зарефакторить
func FindGroupStageById(groupStageId int) (models.GroupStage, error) {
	var groupStage models.GroupStage
	rows, err := db.Query(`SELECT gs.id, gs.date_start, gs.date_end, gs.name,gs.is_finished, g.id as group_id, t.name, t.description  
		FROM group_stages gs
  		JOIN groups g ON gs.id = g.group_stage_id
  		JOIN teams t ON g.id = t.group_id WHERE gs.id = $1`, groupStageId)
	if err != nil {
		return groupStage, err
	}
	groups := groupStage.Groups
	for rows.Next() {
		var tmpGroup models.Group
		var team models.Team
		err = rows.Scan(&groupStage.Id, &groupStage.DateStart, &groupStage.DateEnd, &groupStage.Name, &groupStage.IsFinished,
			&tmpGroup.Id, &team.Name, &team.Description)
		if err == nil {
			//Пока не понимаю как тут смапить результат запроса в объект, поэтому пришлось такое сделать=(
			if len(groups) == 0 {
				tmpGroup.Teams = append(tmpGroup.Teams, team)
				groups = append(groups, tmpGroup)
			} else {
				flag := true
				for i := 0; i < len(groups); i++ {
					if groups[i].Id == tmpGroup.Id {
						groups[i].Teams = append(groups[i].Teams, team)
						flag = false
						break
					}
				}
				if flag {
					tmpGroup.Teams = append(tmpGroup.Teams, team)
					groups = append(groups, tmpGroup)
				}
			}
		} else {
			return groupStage, err
		}
	}

	if err != nil {
		return groupStage, err
	}
	groupStage.Groups = groups
	return groupStage, nil
}

func FindGroupStageByIsFinishedFalse() (models.GroupStage, error) {
	var groupStage models.GroupStage
	rows, err := db.Query(`SELECT gs.id, gs.date_start, gs.date_end, gs.name,gs.is_finished, g.id as group_id, t.name, t.description  
		FROM group_stages gs
  		JOIN groups g ON gs.id = g.group_stage_id
  		JOIN teams t ON g.id = t.group_id WHERE gs.is_finished = FALSE`)
	if err != nil {
		return groupStage, err
	}

	for rows.Next() {
		var tmpGroup models.Group
		var team models.Team
		err = rows.Scan(&groupStage.Id, &groupStage.DateStart, &groupStage.DateEnd, &groupStage.Name, &groupStage.IsFinished,
			&tmpGroup.Id, &team.Name, &team.Description)
		if err == nil {
			//Пока не понимаю как тут смапить результат запроса в объект, поэтому пришлось такое сделать=(
			if len(groupStage.Groups) == 0 {
				tmpGroup.Teams = append(tmpGroup.Teams, team)
				groupStage.Groups = append(groupStage.Groups, tmpGroup)
			} else {
				flag := true
				for i := 0; i < len(groupStage.Groups); i++ {
					if groupStage.Groups[i].Id == tmpGroup.Id {
						groupStage.Groups[i].Teams = append(groupStage.Groups[i].Teams, team)
						flag = false
						break
					}
				}
				if flag {
					tmpGroup.Teams = append(tmpGroup.Teams, team)
					groupStage.Groups = append(groupStage.Groups, tmpGroup)
				}
			}
		} else {
			return groupStage, err
		}
	}

	if err != nil {
		return groupStage, err
	}

	return groupStage, nil
}

func saveGroupStage(groupStage models.GroupStage) (models.GroupStage, error) {
	err := db.QueryRow(`INSERT INTO group_stages (date_start, date_end, name, is_finished) VALUES ($1, $2, $3, $4) RETURNING id`, groupStage.DateStart,
		groupStage.DateEnd, groupStage.Name, groupStage.IsFinished).Scan(&groupStage.Id)
	if err != nil {
		log.Println(err)
		return groupStage, err
	}
	return groupStage, err
}
