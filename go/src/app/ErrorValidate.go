package main

type ErrorValidate struct {
	Mensagem   string `json:"message,omitempty"`
	HTTPStatus int    `json:"statusCode,omitempty"`
}
