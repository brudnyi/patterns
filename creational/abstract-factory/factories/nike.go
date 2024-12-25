package factories

import (
	"patterns/creational/abstract-factory/abstract"
	"patterns/creational/abstract-factory/concrete"
)

type Nike struct{}

func (n *Nike) MakeShoe() abstract.IShoe {
	return &concrete.NikeShoe{
		Shoe: abstract.Shoe{
			Logo: "nike",
			Size: 12,
		},
	}
}

func (n *Nike) MakeShort() abstract.IShort {
	return &concrete.NikeShort{
		Short: abstract.Short{
			Logo: "nike",
			Size: 12,
		},
	}
}
