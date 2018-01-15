package services

import (
	"database/sql"
	"log"
	"toss-up/models"
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

func SaveTeams(teams []models.Team) ([]models.Team, error) {
	var err error
	for e := range teams {
		err := db.QueryRow(`INSERT INTO Teams(name, description) VALUES ($1, $2) RETURNING id`,
			teams[e].Name, teams[e].Description).Scan(&teams[e].Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return teams, err
}

func DeleteTeam(teamId int) (bool, error) {
	//TODO Хотелось бы понимать когда команда просто не найдена в БД, чтобы возвращать 404
	_, err := db.Exec(`DELETE FROM Teams where id = $1`, teamId)

	if err != nil {
		return false, err
	}

	return true, err
}

func UpdateTeam(teamId int, team models.Team) (models.Team, error) {
	dbTeam, err := FindById(teamId)
	//TODO Переделать
	if &dbTeam == nil || err != nil {
		return dbTeam, err
	}
	_, err = db.Exec(`UPDATE Teams SET name = $1, description = $2, group_id = $3 WHERE id = $4`,
		team.Name, team.Description, team.GroupId, teamId)

	dbTeam.Name = team.Name
	dbTeam.Description = team.Description
	//if err != nil {
	//	return Team{}, err
	//}
	return dbTeam, err
}

func FindById(teamId int) (models.Team, error) {
	var dbTeam models.Team
	var id int
	var name string
	var description string
	err := db.QueryRow(`SELECT id, name, description FROM Teams WHERE id = $1`, teamId).Scan(&id, &name, &description)
	if err != nil {
		return dbTeam, err
	}
	dbTeam = models.Team{Id: id, Name: name, Description: description}
	return dbTeam, err
}
