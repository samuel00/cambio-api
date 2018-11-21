package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// smsServiceMock
type repositorioMock struct {
	mock.Mock
}

func (r *repositorioMock) getCambio(cambios *cambios, data string) error {
	//args := r.Called(cambios, data)
	args := r.Called(cambios, data)
	return args.Error(0)
}

func (r *repositorioMock) DummyFunc() {
	fmt.Println("Dummy")
}

func TestGetCambio(t *testing.T) {
	repositorioMock := new(repositorioMock)

	// we then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return
	// true as it was successful in sending a notification
	var error error
	listaDecambios := cambios{}
	data := "2018-11-14"
	repositorioMock.On("getCambio", &listaDecambios, data).Return(error)

	cambioServico := CambioServico{repositorioMock}
	// and call said method
	err := cambioServico.getCambio(&listaDecambios, data)
	assert.Equal(t, err, nil)
}
