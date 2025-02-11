package builders

import "patterns/creational/builder/entites"

type IglooBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.WindowType = "Snow Window"
}

func (b *IglooBuilder) SetDoorType() {
	b.DoorType = "Snow Door"
}

func (b *IglooBuilder) SetNumFloor() {
	b.Floor = 1
}

func (b *IglooBuilder) GetHouse() entites.House {
	return entites.House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
