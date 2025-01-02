package builders

import "patterns/creational/builder/entites"

type NormalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (n *NormalBuilder) NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) SetWindowType() {
	n.WindowType = "Wooden Window"
}

func (n *NormalBuilder) SetDoorType() {
	n.DoorType = "Wooden Door"

}
func (n *NormalBuilder) SetNumFloor() {
	n.Floor = 2
}

func (n *NormalBuilder) GetHouse() entites.House {
	return entites.House{
		WindowType: n.WindowType,
		DoorType:   n.DoorType,
		Floor:      n.Floor,
	}
}
