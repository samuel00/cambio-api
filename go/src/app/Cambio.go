package main

import (
	"time"
)

type Cambio struct {
	TipoCambId   int       `json:"tipoCambio,omitempty" gorm:"column:tc_tipo_cambio;"`
	Id           int       `json:"-" gorm:"primary_key;column:tc_id"`
	Data         time.Time `json:"data,omitempty" gorm:"column:tc_data"`
	ValorOrigem  float64   `json:"valorOrigem,omitempty"  gorm:"column:tc_valor_origem"`
	ValorDestino float64   `json:"valorDestino,omitempty"  gorm:"column:tc_valor_destino"`
}

func (cambio *Cambio) TableName() string {
	return "tab_cambio"
}
