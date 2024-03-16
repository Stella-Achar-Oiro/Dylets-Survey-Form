package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FormData struct {
	Name           string
	Email          string
	Age            string
	Role           string
	Recommendation string
	Improvements   []string
	Comments       string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", surveyFormHandler).Methods("GET")
	r.HandleFunc("/submit", formSubmitHandler).Methods("POST")

    r.PathPrefix("/client/").Handler(http.StripPrefix("/client/", http.FileServer(http.Dir("/workspaces/Dylets-Survey-Form/client"))))

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func surveyFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("/workspaces/Dylets-Survey-Form/client/index.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func formSubmitHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	formData := FormData{
		Name:           r.FormValue("name"),
		Email:          r.FormValue("email"),
		Age:            r.FormValue("age"),
		Role:           r.FormValue("role"),
		Recommendation: r.FormValue("recommendation"),
		Improvements:   r.Form["improvements"],
		Comments:       r.FormValue("comments"),
	}

	// Convert the formData struct to JSON
	jsonData, err := json.Marshal(formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data as the response
	w.Write(jsonData)
}
