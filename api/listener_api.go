package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"testProject/models"
)

type ListenerApi struct {
	DB *sql.DB
}

func (a *ListenerApi) GetListener(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid listener ID")
		return
	}

	l := models.Listener{ID: id}
	if err := l.GetListener(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Listener not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, l)
}

func (a *ListenerApi) DeleteListener(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid listener ID")
		return
	}

	l := models.Listener{ID: id}

	if err := l.DeleteListener(a.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *ListenerApi) CreateListener(w http.ResponseWriter, r *http.Request)  {
	listener := models.Listener{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&listener); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := listener.CreateListener(a.DB); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, listener)
}