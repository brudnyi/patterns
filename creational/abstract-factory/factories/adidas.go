package factories

import (
	"patterns/creational/abstract-factory/abstract"
	"patterns/creational/abstract-factory/concrete"
)

type Adidas struct{}

func (n *Adidas) MakeShoe() abstract.IShoe {
	return &concrete.AdidasShoe{
		Shoe: abstract.Shoe{
			Logo: "adidas",
			Size: 14,
		},
	}
}

func (n *Adidas) MakeShort() abstract.IShort {
	return &concrete.AddidasShort{
		Short: abstract.Short{
			Logo: "adidas",
			Size: 14,
		},
	}
}
