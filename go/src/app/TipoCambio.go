package main

import "github.com/jinzhu/gorm"

type TipoCambio struct {
	gorm.Model
	Id           int    `gorm:"id"; json:"-"`
	MoedaOrigem  string `json:"moedaOrigem,omitempty"`
	MoedaDestino string `json:"moedaDestino,omitempty"`
}
