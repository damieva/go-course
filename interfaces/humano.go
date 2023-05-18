package interfaces

/*
	las interfaces sirve solamente para añadir la definicion de los objectos
*/

type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
	EstaVivo() bool
}
