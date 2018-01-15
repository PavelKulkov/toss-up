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
	router.HandleFunc("/group_stages/generate", controllers.CreateGroupStage).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodGet)
	router.HandleFunc("/group_stages/{groupStageId}", controllers.GetGroupStage).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodGet)
	router.HandleFunc("/timetables/generate", controllers.GenerateTimetable).Headers("Content-Type", "application/json; charset=utf-8").Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
