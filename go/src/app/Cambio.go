package main

import (
	"time"
)

type Cambio struct {
	TipoCambio   TipoCambio `json:"tipoCambio,omitempty" gorm:"foreignkey:TipoCambId;association_foreignkey:TipoCambioId"`
	TipoCambId   int        `gorm:"column:tc_tipo_cambio"`
	Id           int        `json:"-" gorm:"primary_key;column:tc_id"`
	Data         time.Time  `json:"data,omitempty" gorm:"column:tc_data"`
	ValorOrigem  float64    `json:"valorOrigem,omitempty"  gorm:"column:tc_valor_origem"`
	ValorDestino float64    `json:"valorDestino,omitempty"  gorm:"column:tc_valor_destino"`
}

func (cambio *Cambio) TableName() string {
	return "tab_cambio"
}

type cambios struct {
	ListaCambios []Cambio `json:"listaCambios,omitempty"`
}
