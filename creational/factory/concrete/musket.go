// concrete product

package concrete

import (
	"patterns/creational/factory/abstract"
)

type Musket struct {
	abstract.Gun
}

func NewMusket() abstract.IGun {
	return &Musket{
		abstract.Gun{
			Name:  "Musket Gun",
			Power: 4,
		},
	}
}
