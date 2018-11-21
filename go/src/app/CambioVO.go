package main



type CambioVO struct{
  Retorno RetornoVo
  Cambios []Cambio `json:"list,omitempty"`
}
