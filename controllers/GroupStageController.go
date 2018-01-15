package controllers

import (
	"toss-up/services"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func CreateGroupStage(w http.ResponseWriter, r *http.Request) {
	stage, err := services.CreateGroupStage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stage)
	}
}

func GetGroupStage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var groupStageId int
	groupStageId, err := strconv.Atoi(vars["groupStageId"])
	if &groupStageId == nil || err != nil {
		http.Error(w, "group stage id not found", http.StatusBadRequest)
	}
	stage, err := services.FindGroupStageById(groupStageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if &stage == nil {
		http.Error(w, "group stage not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stage)
}
