package main

import (
	"fmt"
	"time"

	"github.com/damieva/go-course/models"
)

var u = new(models.User)

func main() {
	//u := new(models.User) //Creamos un objetos
	u.AddUser(10, "Denis", time.Now(), true)
	fmt.Printf("u: %v\n", u)
}
