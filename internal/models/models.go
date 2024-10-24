package models

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

var persona = Persona{
	Nombre:   "Juan",
	Apellido: "Perez",
	Edad:     30,
}
