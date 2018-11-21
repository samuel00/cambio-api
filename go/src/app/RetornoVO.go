package main

type RetornoVo struct {
	StatusCode int    `json:"HTTPStatusCode,omitempty"`
	Mensagem   string `json:"mesage,omitempty"`
}
