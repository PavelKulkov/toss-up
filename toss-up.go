package main

import (
	_ "github.com/lib/pq"
	"log"
	"github.com/gorilla/mux"
	"toss-up/controllers"
	"net/http"
)



func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/teams", controllers.CreateTeam).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodPost)
	router.HandleFunc("/teams/{teamId}", controllers.DeleteTeam).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodDelete)
	router.HandleFunc("/teams/{teamId}", controllers.UpdateTeam).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodPut)
	router.HandleFunc("/group_stage/generate", controllers.CreateGroupStage).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodGet)
	router.HandleFunc("/group_stage/{groupStageId}", controllers.GetGroupStage).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodGet)
	router.HandleFunc("/timetable/generate", controllers.GenerateTimetable).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}










//var Teams []Team
//Teams = append(Teams, Team{Name:"CSKA", Description:"11"})
//Teams = append(Teams, Team{Name:"SPARTAK", Description:"11"})
//Teams = append(Teams, Team{Name:"ZENIT", Description:"11"})
//Teams = append(Teams, Team{Name:"BARCELONA", Description:"11"})
//Teams = append(Teams, Team{Name:"MU", Description:"11"})
//Teams = append(Teams, Team{Name:"MAN.CITY", Description:"11"})
////Teams = append(Teams, Team{Name: "testTeam", Description:"11"})
////Teams = append(Teams, Team{Name: "8", Description:"11"})
////Teams = append(Teams, Team{Name: "9", Description:"11"})
////Teams = append(Teams, Team{Name: "10", Description:"11"})
////Teams = append(Teams, Team{Name: "11", Description:"11"})
//ints, _ := generateGroups(len(Teams))
//
//distributeTeams := distributeTeams(ints, Teams)
//
//for e := range distributeTeams {
//	fmt.Println(generateTimeTable(distributeTeams[e]))
//}

