package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"testProject/models"
)

type ListenerApi struct {
	DB *sql.DB
}

func (listenerApi *ListenerApi) GetListener(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid listener ID")
		return
	}

	listener := models.Listener{ID: id}
	if err := listener.GetListener(listenerApi.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			RespondWithError(w, http.StatusNotFound, "Listener not found")
		default:
			RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RespondWithJSON(w, http.StatusOK, listener)
}

func (listenerApi *ListenerApi) DeleteListener(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid listener ID")
		return
	}

	listener := models.Listener{ID: id}

	if err := listener.DeleteListener(listenerApi.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (listenerApi *ListenerApi) CreateListener(w http.ResponseWriter, r *http.Request) {
	listener := models.Listener{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&listener); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	if err := listener.CreateListener(listenerApi.DB); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, listener)
}

func (listenerApi *ListenerApi) UpdateListener(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid listener ID")
		return
	}

	var listener models.Listener
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&listener); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	listener.ID = id

	if err := listener.UpdateListener(listenerApi.DB); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, listener)
}
