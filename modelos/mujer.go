package modelos

type Mujer struct {
	Hombre //la mujer esta HEREDANDO todas las propiedades del hombre
}

func (m *Mujer) Respirar()    { m.respirando = true }
func (m *Mujer) Comer()       { m.comiendo = true }
func (m *Mujer) Pensar()      { m.pensando = true }
func (m *Mujer) Sexo() string { return "Mujer" }

func (m *Mujer) EstaVivo() bool {
	if m.respirando {
		return true
	}
	return false
}
