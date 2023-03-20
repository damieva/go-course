package main

import (
	"time"

	"github.com/damieva/go-course/models"
)

func main() {
	u := new(models.User) //Creamos un objetos
	u.AddUser(10, "Denis", time.Now(), true)
	fmt.Println(u)
}
