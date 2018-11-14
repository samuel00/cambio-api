package main

var currentId int

var cambio []Cambio

// Give us some seed data
func init() {

	/*RepoCreatePerson(Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	RepoCreatePerson(Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})*/
}

func RepoGetCambio(cambios *cambios, data string) error {
	initDb()
	rows, err := db.Query("SELECT tc_id,tc_data,tc_valor_origem,tc_valor_destino,ttc.id,ttc_moeda_origem,ttc_moeda_destino FROM tab_cambio tc INNER JOIN tab_tipo_cambio ttc ON ttc.id = tc.tc_tipo_cambio where tc_data = '" + data + "'")
	if err != nil {
		return err
	}
	defer rows.Close()
	defer db.Close()
	for rows.Next() {
		c := Cambio{}
		err = rows.Scan(
			&c.Id,
			&c.Data,
			&c.ValorOrigem,
			&c.ValorDestino,
			&c.TipoCambio.Id,
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
