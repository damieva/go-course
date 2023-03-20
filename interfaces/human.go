package interfaces

// una interfaz no lleva vbles, solo la definicion de las funciones (métodos) de un objeto
type Human interface {
	Respirar()
	Comer()
	Pensar()
	Sexo() string
}
