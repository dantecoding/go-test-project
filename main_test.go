package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const URL = "localhost:8000"

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := router()
	router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Needed code %d. Got %d\n", expected, actual)
	}
}

func TestCreateEvent(t *testing.T) {
	payload := []byte(`{"name":"test event","published":false}`)

	req, _ := http.NewRequest("POST", URL+"/event/", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test event" {
		t.Errorf("Expected event name to be 'test event'. Got '%v'", m["name"])
	}

	if m["published"] != false {
		t.Errorf("Expected published to be 'false'. Got '%v'", m["published"])
	}
}

func TestGetEvent(t *testing.T) {
	req, _ := http.NewRequest("GET", URL+"/event/2", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
