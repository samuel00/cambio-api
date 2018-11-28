package main

import (
	"time"
)

type Cambio struct {
	Id           int       `json:"-" gorm:"primary_key;column:tc_id"`
	Data         time.Time `json:"data,omitempty" gorm:"column:tc_data"`
	ValorOrigem  float64   `json:"valorOrigem,omitempty"  gorm:"column:tc_valor_origem"`
	ValorDestino float64   `json:"valorDestino,omitempty"  gorm:"column:tc_valor_destino"`
	MoedaOrigem  string    `gorm:"column:tc_moeda_origem" json:"moedaOrigem,omitempty"`
	MoedaDestino string    `gorm:"column:tc_moeda_destino" json:"moedaDestino,omitempty"`
}

func (cambio *Cambio) TableName() string {
	return "tab_cambio"
}

type cambios struct {
	ListaCambios []Cambio `json:"listaCambios,omitempty"`
}
