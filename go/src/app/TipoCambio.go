package main

type TipoCambio struct {
	Id           int    `json:"-"`
	MoedaOrigem  string `json:"moedaOrigem,omitempty"`
	MoedaDestino string `json:"moedaDestino,omitempty"`
}
