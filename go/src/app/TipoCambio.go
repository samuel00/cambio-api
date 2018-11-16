package main

type TipoCambio struct {
	TipoCambioId int    `gorm:"column:ttc_id" json:"-"`
	MoedaOrigem  string `gorm:"column:ttc_moeda_origem" json:"moedaOrigem,omitempty"`
	MoedaDestino string `gorm:"column:ttc_moeda_destino" json:"moedaDestino,omitempty"`
}

func (tipoCambio *TipoCambio) TableName() string {
	return "tab_tipo_cambio"
}
