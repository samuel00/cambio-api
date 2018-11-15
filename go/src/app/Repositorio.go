package main

var currentId int

var cambio []Cambio

// Give us some seed data
func init() {

	/*RepoCreatePerson(Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	RepoCreatePerson(Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})*/
}

func RepoGetCambio(cambios *[]TipoCambio, data string) error {
	initDb()
	//rows, err := db.Query("SELECT tc_id,tc_data,tc_valor_origem,tc_valor_destino,ttc.ttc_id,ttc_moeda_origem,ttc_moeda_destino FROM tab_cambio tc INNER JOIN tab_tipo_cambio ttc ON ttc.ttc_id = tc.tc_tipo_cambio where tc_data = '" + data + "'")
	//db.Where("data = ?", data).Find(&cambios)
	var tipoCambio TipoCambio
	var cambio Cambio
	cambio.Id = 1
	err := db.Model(&cambio).Related(&tipoCambio).Find(&cambios).Error
	if err != nil {
		return err
	}
	return nil
}
