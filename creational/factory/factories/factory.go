package factories

import (
	"fmt"
	"patterns/creational/factory/abstract"
	"patterns/creational/factory/concrete"
)

func GetGun(gunType string) (abstract.IGun, error) {
	if gunType == "ak47" {

		return concrete.NewAk47(), nil
	}
	if gunType == "musket" {
		return concrete.NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
