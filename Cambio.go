package main

import ("time")

type Cambio struct {
    id        int   `json:"-"`
    data time.Time
    valorOrigem float64
    valorDestino float64
    tipoCambio TipoCambio
}
