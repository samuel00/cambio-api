package main

var currentId int

var cambio []Cambio

// Give us some seed data
func init() {

	/*RepoCreatePerson(Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	RepoCreatePerson(Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})*/
}

func RepoGetCambio(cambios *cambios) error {
	initDb()
	rows, err := db.Query("SELECT tc_id, tc_data, tc_valor_origem, tc_valor_destino FROM tab_cambio")
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
