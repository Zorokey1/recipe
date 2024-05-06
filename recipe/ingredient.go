package recipe

import (
	"fmt"
)

type Ingredient struct {
	Measurement
	Name string
}

func (this Ingredient) String() string {
	return fmt.Sprintf("%v %v", this.Measurement.String(), this.Name)
}

func MakeIngredient(numerator int, denominator int, unit string, name string) (Ingredient, error) {
	measurement, err := MakeMeasurement(numerator, denominator, unit)

	if err != nil {
		return Ingredient{}, fmt.Errorf("MakeIngredient: failed to make ingredient: %w", err)
	}

	return Ingredient{Measurement: measurement, Name: name}, err
}
