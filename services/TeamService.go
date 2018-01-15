package services

import (
	"database/sql"
	"log"
	"toss-up/models"
	"errors"
)

func getNewTeams() []models.Team {
	var teams []models.Team
	rows, err := db.Query(`SELECT * FROM teams WHERE group_id ISNULL`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var description string
			var groupId sql.NullInt64
			err = rows.Scan(&id, &name, &description, &groupId)
			if err == nil {
				team := models.Team{Id: id, Name: name, Description: description}
				teams = append(teams, team)
			} else {
				log.Fatal(err)
				return nil
			}
		}
	} else {
		log.Fatal(err)
		return nil
	}
	return teams
}

func CreateTeam(team models.Team) (models.Team, error) {
	err := db.QueryRow(`INSERT INTO Teams(name, description) VALUES ($1, $2) RETURNING id`,
		team.Name, team.Description).Scan(&team.Id)
	if err != nil {
		log.Println(err)
		return team, err
	}
	return team, err
}

func DeleteTeam(teamId int) error {
	team, err := FindTeamById(teamId)
	if err != nil {
		log.Println(err)
		return err
	}
	if team.Name == "" {
		return errors.New("team not found")
	}
	_, err = db.Exec(`DELETE FROM Teams where id = $1`, teamId)
	return err
}

func UpdateTeam(teamId int, team models.Team) (models.Team, error) {
	dbTeam, err := FindTeamById(teamId)
	if err != nil {
		return dbTeam, err
	}
	if dbTeam.Name == "" {
		return dbTeam, errors.New("Team not found")
	}
	_, err = db.Exec(`UPDATE Teams SET name = $1, description = $2, group_id = $3 WHERE id = $4`,
		team.Name, team.Description, team.GroupId, teamId)
	if err != nil {
		return dbTeam, err
	}
	dbTeam.Name = team.Name
	dbTeam.Description = team.Description
	return dbTeam, err
}

func FindTeamById(teamId int) (models.Team, error) {
	var dbTeam models.Team
	err := db.QueryRow(`SELECT id, name, description FROM Teams WHERE id = $1`, teamId).Scan(&dbTeam.Id, &dbTeam.Name,
		&dbTeam.Description)
	if err != nil {
		return dbTeam, err
	}
	return dbTeam, err
}
