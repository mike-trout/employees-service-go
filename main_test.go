// main_test.go

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllEmployees(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?s=1&n=1", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp.Code)

	body := strings.TrimSpace(resp.Body.String())
	jsonBytes, _ := json.Marshal(Employees)
	json := strings.TrimSpace(string(jsonBytes))
	if body != json {
		t.Errorf("Expected %s. Got %s", json, body)
	}
}

func TestGetSingleEmployee(t *testing.T) {
	req, _ := http.NewRequest("GET", "/11100102", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusOK, resp.Code)

	body := strings.TrimSpace(resp.Body.String())
	jsonBytes, _ := json.Marshal(Employees[0])
	json := strings.TrimSpace(string(jsonBytes))
	if body != json {
		t.Fatalf("Expected %s. Got %s", json, body)
	}
}

func TestGetNonExistentEmployee(t *testing.T) {
	req, _ := http.NewRequest("GET", "/99999999", nil)
	resp := executeRequest(req)

	checkResponseCode(t, http.StatusInternalServerError, resp.Code)

	json := `{"error":"Employee not found"}`
	if body := resp.Body.String(); body != json {
		t.Fatalf("Expected %s. Got %s", json, body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	a := App{}
	a.Initialise()
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
