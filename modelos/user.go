package modelos

/*
Una estructura es:
	una aproximacion del POO en go.
	un tipo de datos que puede contener uno o muchos tipos de datos distintos.
*/

import (
	"time"
)

// Este paquete es un estándar en el que solo vamos a colocar definiciones de datos

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// el * sirve para indicar que una funcion pertenece a una estructura
func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAt
	usuario.Status = status
}
