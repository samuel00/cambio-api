package main

var currentId int

var cambio []Cambio

// Give us some seed data
func init() {

	/*RepoCreatePerson(Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	RepoCreatePerson(Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})*/
}

type Repositorio interface {
	getCambio(cambios *cambios, data string) error
}

type CambioRepositorio struct{}

func (c CambioRepositorio) getCambio(cambios *cambios, data string) error {
	initDb()
	rows, err := db.Table("tab_cambio").Select("tc_id, tc_data, tc_valor_origem, tc_valor_destino, ttc_moeda_origem, ttc_moeda_destino").Joins("inner join tab_tipo_cambio on tab_tipo_cambio.ttc_id = tc_id where tc_data = '" + data + "'").Rows()
	for rows.Next() {
		c := Cambio{}
		err = rows.Scan(
			&c.Id,
			&c.Data,
			&c.ValorOrigem,
			&c.ValorDestino,
			&c.TipoCambio.MoedaOrigem,
			&c.TipoCambio.MoedaDestino,
		)
		if err != nil {
			return err
		}
		cambios.ListaCambios = append(cambios.ListaCambios, c)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
