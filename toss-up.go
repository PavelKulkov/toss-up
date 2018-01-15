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
	router.Headers("Content-Type", "application/json")
	router.HandleFunc("/teams", controllers.CreateTeam).Methods(http.MethodPost)
	router.HandleFunc("/teams/{teamId}", controllers.DeleteTeam).Methods(http.MethodDelete)
	router.HandleFunc("/teams/{teamId}", controllers.UpdateTeam).Methods(http.MethodPut)
	router.HandleFunc("/group_stages/generate", controllers.CreateGroupStage).Methods(http.MethodGet)
	router.HandleFunc("/group_stages/{groupStageId}", controllers.GetGroupStage).Methods(http.MethodGet)
	router.HandleFunc("/timetables/generate", controllers.GenerateTimetable).Methods(http.MethodGet)
	router.HandleFunc("/timetables",controllers.GetTimetable ).Methods(http.MethodGet)
	router.HandleFunc("/timetables/{timetableId}", controllers.UpdateResultMatch).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8080", router))
}
