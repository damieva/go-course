package interfaces

/*
	Podemos definir una interfaz superior a las que ya tenemos simplemente con que las otras interfaces definan sus metodos ya lo tendríamos.
	Podriamos defeinir un perro, gato o cualquier planta como un ser vivo y preguntar por las propiedades y en base a ello dar un tratamiento u otro a ese objecto.
	Esto cumpliría con los patrones de herencia y polimorfismo de los POO.
*/

type SerVivo interface {
	EstaVivo() bool
}
