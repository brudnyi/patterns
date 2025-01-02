package directors

import (
	"patterns/creational/builder/builders"
	"patterns/creational/builder/entites"
)

type Director struct {
	Builder builders.IBuilder
}

func NewDirector(b builders.IBuilder) Director {
	return Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b builders.IBuilder) {
	d.Builder = b
}

func (d Director) BuildHouse() entites.House {
	d.Builder.SetNumFloor()
	d.Builder.SetWindowType()
	d.Builder.SetDoorType()
	return d.Builder.GetHouse()
}
