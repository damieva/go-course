package models

import (
	"fmt"
	"time"
)

// Objeto user
type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

// Metodos para user
func (user *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.CreatedAt = createdAt
	user.Status = status
}

func (user *User) GetUser() {
	fmt.Println(user)
}
