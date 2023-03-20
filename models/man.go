package models

type Man struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	vivo       bool
}

// Definiendo estas funciones estamos implementando la interfaz Human
func (m *Man) Respirar()    { m.respirando = true }
func (m *Man) Comer()       { m.comiendo = true }
func (m *Man) Pensar()      { m.pensando = true }
func (m *Man) Sexo() string { return "male" }
