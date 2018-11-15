package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Cambio struct {
	gorm.Model
	TipoCambio      TipoCambio `json:"tipoCambio,omitempty";gorm:"foreignkey:TipoCambioRefer"`
	TipoCambioRefer uint
	Id              int       `json:"-"`
	Data            time.Time `json:"data,omitempty"`
	ValorOrigem     float64   `json:"valorOrigem,omitempty"`
	ValorDestino    float64   `json:"valorDestino,omitempty"`
}

func (Cambio) TableName() string {
	return "tab_cambio"
}

type cambios struct {
	ListaCambios []Cambio `json:"listaCambios,omitempty"`
}
