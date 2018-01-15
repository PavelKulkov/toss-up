package controllers

import (
	"net/http"
	"toss-up/services"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"database/sql"
)

func GenerateTimetable(w http.ResponseWriter, r* http.Request) {

	stage, err := services.FindGroupStageByIsFinishedFalse()

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if stage.Name == "" {
			http.Error(w, "unfinished groupstage not found", http.StatusNotFound)
		} else {
			matchResults, err := services.FindResultsByGroupStageId(stage.Id)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				if matchResults == nil {
					for e := range stage.Groups {
						table := services.GenerateTimeTable(stage.Groups[e])
						services.SaveTimetables(table)
						stage.Groups[e].TimeTable = table
					}
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(stage.Groups)
				} else {
					http.Error(w, "Games already started", http.StatusBadRequest)
				}
			}
		}
	}
}

func GetTimetable(w http.ResponseWriter, r* http.Request) {
	timetables, err := services.FindTimetableCurrentGroupStage()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if timetables == nil {
			http.Error(w, "No current group stage or timetable not yet generated", http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(timetables)
		}
	}
}

func UpdateResultMatch(w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	timetableId, err := strconv.Atoi(vars["timetableId"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	type Result struct {
		Result string `json:"result"`
	}

	decoder := json.NewDecoder(r.Body)
	var result Result
	err = decoder.Decode(&result)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	timetable, err := services.UpdateResultMatch(timetableId, sql.NullString{String:result.Result, Valid:true})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if timetable.Match == "" {
			http.Error(w, "Timetable not found", http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(timetable)
		}
	}
}
