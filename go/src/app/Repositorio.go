package main

import "fmt"

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
	db = db.Where(map[string]interface{}{"tc_data": data}).Find(&cambios.ListaCambios)
	if db.Error != nil {
		fmt.Println(db.Error)
		return db.Error
	}
	return nil
}
