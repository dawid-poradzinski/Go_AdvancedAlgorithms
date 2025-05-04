package monte_carlo_logic

import (
	"math"
	"math/rand"
)

var maxY float64 = 4.0

var maxX float64 = 2

func CalculateInegralValue(interations int) (float64, float64) {

	integralValue := MonteCarlo(interations)

	return integralValue, integralValue / (2 * math.Pi * 4) * 100
}

func MonteCarlo(interations int) float64 {

	insideCircle := 0
	outsideCirlce := 0

	for _ = range interations {

		randomX, randomY := PickRandomPoint()

		calculateY := CalculateValueInPoint(randomX)

		if randomY <= calculateY {
			insideCircle++
		} else {
			outsideCirlce++
		}

	}

	return float64(outsideCirlce) / float64(insideCircle+outsideCirlce) * 8 * math.Pi
}

func PickRandomPoint() (float64, float64) {
	x := rand.Float64() * maxX * math.Pi
	y := rand.Float64() * maxY

	return x, y
}

func CalculateValueInPoint(x float64) float64 {

	result := 0.0

	result += math.Abs(math.Sin(x * math.Pi))
	result += math.Abs(math.Sin(2 * x * math.Pi))
	result += math.Abs(math.Sin(4 * x * math.Pi))
	result += math.Abs(math.Sin(8 * x * math.Pi))

	return result
}
