package controllers

import (
	"github.com/gorilla/mux"
	"strconv"
	"net/http"
	"encoding/json"
	"toss-up/models"
	"toss-up/services"
)

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var team models.Team
	err := decoder.Decode(&team)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	_, err = services.SaveTeams([]models.Team{team})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}

func DeleteTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var teamId int
	teamId, _ = strconv.Atoi(vars["teamId"])
	if &teamId == nil {
		http.Error(w, "team id not found", http.StatusBadRequest)
	}
	result, _ := services.DeleteTeam(teamId)
	if result {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func UpdateTeam(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	var teamId int
	teamId, _ = strconv.Atoi(vars["teamId"])
	if &teamId == nil {
		http.Error(w, "team id not found", http.StatusBadRequest)
	}
	decoder := json.NewDecoder(r.Body)
	var team models.Team
	err := decoder.Decode(&team)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	updatedTeam, err := services.UpdateTeam(teamId, team)
	if err != nil || &updatedTeam == nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedTeam)
	}
}