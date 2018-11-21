package main



type CambioVO struct{
  StatusCode int `json:"HTTPStatusCode,omitempty"`
  Mensagem string `json:"mesage,omitempty"`
  QuantidadeCambio int `json:"length,omitempty"`
  Cambios []Cambio `json:"list,omitempty"`
}
