// app.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// App - Application stuct
type App struct {
	Router *mux.Router
}

// Initialise - Initialise the application
func (a *App) Initialise() {
	a.Router = mux.NewRouter()
	a.initialiseRoutes()
}

// Run - Run the application
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initialiseRoutes() {
	a.Router.HandleFunc("/", a.getEmployees).Methods("GET")
	a.Router.HandleFunc("/{id:[0-9]{8}}", a.getEmployee).Methods("GET")
}

func (a *App) getEmployees(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var start uint64
	var err error
	if s, ok := query["s"]; ok {
		start, err = strconv.ParseUint(s[0], 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid s parameter")
			return
		}
	}
	var number uint64
	if n, ok := query["n"]; ok {
		number, err = strconv.ParseUint(n[0], 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid n parameter")
			return
		}
	}
	employees, err := getEmployees(start, number)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, employees)
}

func (a *App) getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid employee ID")
		return
	}

	employee, err := getEmployee(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, employee)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
