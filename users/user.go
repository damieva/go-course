package users

import (
	"fmt"
	"time"

	"github.com/damieva/go-course/modelos"
)

func AltaUsuario() {
	u := new(modelos.User) //En esta linea instanciamos un User
	u.AddUser(10, "Denis", time.Now(), true)
	fmt.Println(u)
}
