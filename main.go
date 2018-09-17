package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testProject/api"
)

var (
	listener api.ListenerApi
	event    api.EventApi
)

func main() {
	var db *sql.DB
	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", "root", "password", "test")

	// open DB connection
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// inject DB connection in API
	listener.DB = db
	event.DB = db

	router := router()

	server := http.ListenAndServe(":8000", router)

	log.Fatal(server)
}

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/listener/{id}", listener.GetListener).Methods("GET")
	router.HandleFunc("/listener/", listener.CreateListener).Methods("POST")
	router.HandleFunc("/listener/{id}", listener.DeleteListener).Methods("DELETE")
	router.HandleFunc("/listener/{id}", listener.UpdateListener).Methods("PUT")

	router.HandleFunc("/event/{id}", event.GetEvent).Methods("GET")
	router.HandleFunc("/event/", event.CreateEvent).Methods("POST")
	router.HandleFunc("/event/publish/{id}", event.PublishEvent).Methods("POST")
	router.HandleFunc("/event/{id}", event.DeleteEvent).Methods("DELETE")

	router.Use(loggingMiddleware)

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
