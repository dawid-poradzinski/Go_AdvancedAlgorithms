package bisection_logic

import (
	bisection_model "AdvancedAlgorithmsProjects/src/bisection/model"
	"errors"
	"math"
)

func ValidateRange(delta float64, min float64, max float64) error {

	if min > max {
		return errors.New("max must be bigger than min")
	}

	if delta > math.Abs(min-max) {
		return errors.New("delta must be lower than range")
	}

	return nil

}

func CalculateValueInPoint(numbers []float64, x float64) float64 {

	power := 1.0

	result := 0.0

	for i := 0; i < len(numbers); i++ {
		if power == 0 {
			break
		}
		result += numbers[i] * power
		power *= x
	}

	return result

}

func Bisection(model bisection_model.BisectionRequest) (float64, error) {

	if err := ValidateRange(model.Delta, model.Min, model.Max); err != nil {
		return -1, err
	}

	minY := CalculateValueInPoint(model.Numbers, model.Min)
	maxY := CalculateValueInPoint(model.Numbers, model.Max)

	if minY*maxY > 0 {
		return -1, errors.New("Range starts and end with the same + or - sign")
	}

	return CheckRange(model.Min, model.Max, model.Numbers, minY, maxY, model.Delta), nil
}

func CheckRange(minX float64, maxX float64, numbers []float64, minY float64, maxY float64, epsilon float64) float64 {

	middleX := 0.0

	for math.Abs(minX-maxX) > epsilon {

		middleX = (minX + maxX) / 2

		middleY := CalculateValueInPoint(numbers, middleX)

		if middleY == 0 {
			return middleX
		} else if (minY > 0) == (middleY > 0) {
			minX = middleX
			minY = middleY
		} else {
			maxX = middleX
			maxY = middleY
		}
	}

	return middleX

}
