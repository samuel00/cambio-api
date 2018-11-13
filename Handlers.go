package main

import (
	"encoding/json"
	"net/http"
)

// Display all from the people var
func GetCambio(w http.ResponseWriter, r *http.Request) {
	cambioJson, err := json.Marshal(cambios)
	if err != nil {
		panic(err)
	}
	setRequest(w, cambioJson)
}

func setRequest(w http.ResponseWriter, peopleJson []uint8) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(peopleJson)
}
