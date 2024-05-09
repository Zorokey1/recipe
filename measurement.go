package recipe

import (
	"fmt"
)

type Measurement struct {
	Fraction `json:"fraction"`
	Unit     string `json:"unit"`
}

func (this Measurement) String() string {
	unit := this.Unit
	if this.Fraction.Value() > 1 {
		unit += "s"
	}

	return fmt.Sprintf("%s %v", this.Fraction.String(), unit)
}

func MakeMeasurement(numerator int, denominator int, unit string) (Measurement, error) {
	frac, err := MakeFraction(numerator, denominator)

	if err != nil {
		return Measurement{}, fmt.Errorf("MakeMeasurement: failed to create new measurement: %w", err)
	}

	return Measurement{Fraction: frac, Unit: unit}, err
}
