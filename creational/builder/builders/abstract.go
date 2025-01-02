package builders

import "patterns/creational/builder/entites"

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() entites.House
}
