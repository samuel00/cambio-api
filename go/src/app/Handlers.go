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
		retorno := RetornoVo{Mensagem: "O parâmentro informado " + params["data"] + " não corresponde a uma data válida no forma YYYY-MM-DD", StatusCode: 400, QuantidadeCambio: 0}
		errorValidate := &ErrorValidate{Retorno: retorno}
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

	fmt.Println(t)
	var cambioVo CambioVO
	if len(listaDecambios.ListaCambios) != 0 {
		cambioVo.Cambios = listaDecambios.ListaCambios
		cambioVo.Retorno.QuantidadeCambio = len(listaDecambios.ListaCambios)
		cambioVo.Retorno.StatusCode = http.StatusOK
		cambioVo.Retorno.Mensagem = http.StatusText(200)
	} else {
		cambioVo.Cambios = listaDecambios.ListaCambios
		cambioVo.Retorno.QuantidadeCambio = len(listaDecambios.ListaCambios)
		cambioVo.Retorno.StatusCode = http.StatusNotFound
		cambioVo.Retorno.Mensagem = http.StatusText(404)
	}

	cambioJson, err := json.Marshal(cambioVo)
	if err != nil {
		panic(err)
	}

	setRequest(w, cambioJson, cambioVo.Retorno.StatusCode)
}

func setRequest(w http.ResponseWriter, cambioJson []uint8, status int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	w.Write(cambioJson)
}
