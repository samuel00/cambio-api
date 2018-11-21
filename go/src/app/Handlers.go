package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Display all from the people var
func GetCambio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	//tipoCambio := []TipoCambio{}
	t, error := time.Parse("2006-01-02", params["data"])
	if error != nil {
		errorValidate := &ErrorValidate{Mensagem: "O parâmentro informado não corresponde a uma data válida no forma YYYY-MM-DD", HTTPStatus: 400}
		cambioJson, err := json.Marshal(errorValidate)
		if err != nil {
			panic(err)
		}
		setRequest(w, cambioJson, http.StatusBadRequest)
		return
	}

	listaDecambios := cambios{}
	cambioRepositorio := CambioRepositorio{}
	cambioServico := CambioServico{cambioRepositorio}
	err := cambioServico.getCambio(&listaDecambios, params["data"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cambioJson, err := json.Marshal(listaDecambios)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
	setRequest(w, cambioJson, http.StatusOK)
}

func setRequest(w http.ResponseWriter, cambioJson []uint8, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	w.Write(cambioJson)
}
