package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-test-project/models"
	"net/http"
	"strconv"
)

type EventApi struct {
	DB *sql.DB
}

func (eventApi *EventApi) GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid event ID")
		return
	}

	l := models.Event{ID: id}
	if err := l.GetEvent(eventApi.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Event not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

}

func (eventApi *EventApi) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&event); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	if err := event.CreateEvent(eventApi.DB); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, event)
}

func (eventApi *EventApi) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid event ID")
		return
	}

	event := models.Event{ID: id}

	if err := event.DeleteEvent(eventApi.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (eventApi *EventApi) PublishEvent(w http.ResponseWriter, r *http.Request) {

}
