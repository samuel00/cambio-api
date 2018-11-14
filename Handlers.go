package main

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

// Display all from the people var
func GetCambio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	listaCambios := cambios{}
	err := RepoGetCambio(&listaCambios, params["data"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cambioJson, err := json.Marshal(listaCambios)
	if err != nil {
		panic(err)
	}

	fmt.Println(listaCambios)
	setRequest(w, cambioJson)
}

func setRequest(w http.ResponseWriter, cambioJson []uint8) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(cambioJson)
}
