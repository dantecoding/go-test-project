package api

import (
	"database/sql"
	"net/http"
)

type EventApi struct {
	DB *sql.DB
}

func (e *EventApi) GetEvent(w http.ResponseWriter, r *http.Request)  {

}

func (e *EventApi) CreateEvent(w http.ResponseWriter, r *http.Request)  {

}

func (e *EventApi) DeleteEvent(w http.ResponseWriter, r *http.Request)  {

}

func (e *EventApi) PublishEvent(w http.ResponseWriter, r *http.Request)  {

}