package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Display all from the people var
func GetCambio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cambios := cambios{}
	//tipoCambio := []TipoCambio{}
	err := RepoGetCambio(&cambios, params["data"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		fmt.Println("Erro:----------------------->")
		fmt.Println(err.Error())
		return
	}

	cambioJson, err := json.Marshal(cambios)
	if err != nil {
		panic(err)
	}

	fmt.Println(cambios)
	setRequest(w, cambioJson)
}

func setRequest(w http.ResponseWriter, cambioJson []uint8) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(cambioJson)
}
