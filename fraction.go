package recipe

import (
	"fmt"
)

type Fraction struct {
	Numerator   int `json:"numerator"`
	Denominator int `json:"denominator"`
}

var ErrDivideByZero error = fmt.Errorf("fraction: can't divide by 0")


func MakeFraction(numerator int, denominator int) (Fraction, error) {
	if denominator == 0 {
		return Fraction{}, ErrDivideByZero
	}

	return Fraction{Numerator: numerator, Denominator: denominator}, nil
}

func MakeCopyFraction(prev Fraction) Fraction {
	return Fraction{Numerator: prev.Numerator, Denominator: prev.Denominator}
}

func (this *Fraction) Multiply(num int) {
	this.Numerator *= num
}

func (this *Fraction) Divide(num int) error {
	if num == 0 {
		return ErrDivideByZero
	}

	this.Denominator *= num

	return nil
}

func (this Fraction) Value() float32 {
	return float32(this.Numerator) / float32(this.Denominator)
}

func (this *Fraction) Reduce() {
	var minValue int = min(this.Numerator, this.Denominator)

	for i := minValue; i >= 2; i-- {
		if this.Denominator%i == 0 && this.Numerator%i == 0 {
			for this.Denominator%i == 0 && this.Numerator%i == 0 {
				this.Numerator /= i
				this.Denominator /= i
			}

			break
		}
	}

}

func (this Fraction) String() string {
	this.Reduce()
	var count int = this.Numerator / this.Denominator

	if this.Denominator == 1 {
		return fmt.Sprintf("%v", this.Numerator)
	} else if count > 0 && this.Numerator > this.Denominator {
		return fmt.Sprintf("%v %v/%v", count, this.Numerator-count*this.Denominator, this.Denominator)
	} else {
		return fmt.Sprintf("%v/%v", this.Numerator, this.Denominator)
	}
}
