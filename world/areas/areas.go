package areas

import ()

type Area struct {
	Name string
	Uuid string
}

func NewLimbo() *Area {
	return &Area{
		Name: "Limbo",
		Uuid: "limbo",
	}
}
