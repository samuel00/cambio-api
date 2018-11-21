package main

type CambioServico struct {
	repositorio Repositorio
}

func (c CambioServico) getCambio(cambios *cambios, data string) error {
	return c.repositorio.getCambio(cambios, data)
}
