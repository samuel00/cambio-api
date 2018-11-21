package main

type RetornoVo struct {
	StatusCode       int    `json:"HTTPStatusCode,omitempty"`
	Mensagem         string `json:"message,omitempty"`
	QuantidadeCambio int    `json:"length,omitempty"`
}
