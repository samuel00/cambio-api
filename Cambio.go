package main

import ("time")

type Cambio struct {
    Id                int `json:"-"`
    Data              time.Time `json:"data,omitempty"`
    ValorOrigem       float64 `json:"valorOrigem,omitempty"`
    ValorDestino      float64 `json:"valorDestino,omitempty"`
    //tipoCambio TipoCambio `json:"tipoCambio,omitempty"`
}

type cambios struct {
    ListaCambios      []Cambio `json:"listaCambios,omitempty"`;
}
