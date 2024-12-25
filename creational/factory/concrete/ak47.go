// concrete product

package concrete

import (
	"patterns/creational/factory/abstract"
)

type Ak47 struct {
	abstract.Gun
}

func NewAk47() abstract.IGun {
	return &Ak47{
		Gun: abstract.Gun{
			Name:  "AK47 gun",
			Power: 7,
		},
	}
}
