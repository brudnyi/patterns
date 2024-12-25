package factories

import (
	"fmt"
	"patterns/creational/abstract-factory/abstract"
)

type ISportsFactory interface {
	MakeShoe() abstract.IShoe
	MakeShort() abstract.IShort
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	if brand == "nike" {
		return &Nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
