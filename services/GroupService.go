package services

import (
	"errors"
	"strconv"
	"toss-up/models"
	"log"
)

const (
	STANDART_COUNT_TEAMS_IN_GROUP = 4
	MIN_COUNT_TEAMS_IN_GROUP      = 3
)

func SaveGroups(groups []models.Group) ([]models.Group, error) {
	var err error
	for e := range groups {
		err := db.QueryRow(`INSERT INTO groups (group_stage_id) VALUES($1) RETURNING id`,
			groups[e].GroupStageId).Scan(&groups[e].Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return groups, err
}

func GenerateGroups(countOfTeams int) ([]int, error) {
	if countOfTeams < MIN_COUNT_TEAMS_IN_GROUP {
		return nil, errors.New("Для генерации групп кол-во команд не может быть <" + strconv.Itoa(MIN_COUNT_TEAMS_IN_GROUP))
	}
	i := countOfTeams % STANDART_COUNT_TEAMS_IN_GROUP
	if i == 0 {
		return getGroups(countOfTeams, STANDART_COUNT_TEAMS_IN_GROUP), nil
	}
	if countOfTeams%3 == 0 {
		return getGroups(countOfTeams, MIN_COUNT_TEAMS_IN_GROUP), nil
	}
	if i >= 3 {
		return append(getGroups(countOfTeams, STANDART_COUNT_TEAMS_IN_GROUP), i), nil
	} else {
		groups := getGroups(countOfTeams, STANDART_COUNT_TEAMS_IN_GROUP)
		for j := 0; j < i; j++ {
			groups[j] += 1
		}
		return groups, nil
	}
}

func DistributeTeams(countOfGroups []int, teams []models.Team, groupStageId int) []models.Group {
	var groups []models.Group
	tmp := 0
	limit := 0
	for _, countOfTeamsInGroup := range countOfGroups {
		group := models.Group{GroupStageId: groupStageId}
		limit += countOfTeamsInGroup
		for ; tmp < limit; tmp++ {
			group.Teams = append(group.Teams, teams[tmp])
		}
		groups = append(groups, group)
	}
	return groups
}

func getGroups(countOfAllTeams, countOfTeamsInGroup int) []int {
	var result []int
	for j := 0; j < (countOfAllTeams / countOfTeamsInGroup); j++ {
		result = append(result, countOfTeamsInGroup)
	}
	return result
}
