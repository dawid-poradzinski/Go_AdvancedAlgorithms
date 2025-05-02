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

func Bisection(model bisection_model.BisectionRequest) ([]float64, error) {

	err := ValidateRange(model.Delta, model.Min, model.Max)
	if err == nil {
		return nil, err
	}

	result := []float64{}

	return result, nil
}

func CheckRange(numbers []float64, min float64, max float64, delta float64, dive int, lastY float64) []float64 {

	result := []float64{}

	if dive < 3 {

		tenPercent := (math.Abs(min-max) * 0.1)

		if tenPercent < delta {
			return result
		}

		last := min

		for x := last + tenPercent; x <= max; x += tenPercent {

			y := CalculateValueInPoint(numbers, x)

			if y == 0 {
				result = append(result, x)
			} else if (lastY > 0 && y > 0) || (lastY < 0 && y < 0) {
				result = append(result, CheckRange(numbers, last, x, delta, dive+1, y)...)
			} else {
				result = append(result, CheckRange(numbers, last, x, delta, 0, y)...)
			}

			lastY = y
		}

	}

	return result
}
