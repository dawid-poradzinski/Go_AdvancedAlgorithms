package service

import (
	"AdvancedAlgorithmsProjects/src/quadruple/model"
)

func AnalyzeQuadruplets(start int, end int) model.QuadrupletResponse {
	primes := MakePrimes(start, end)
	return FindPairsAndQuadruplets(primes)
}

func MakePrimes(start int, end int) []int {

	primes := []int{}
	primesInRange := []int{}

	for i := 2; i <= end; i += 2 {
		isPrime := true
		for _, num := range primes {
			if num*num > i {
				break
			}
			if i%num == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			if i >= start && i <= end {
				primesInRange = append(primesInRange, i)
			}
			if i == 2 {
				i--
			}
		}
	}

	return primesInRange

}

func FindPairsAndQuadruplets(primes []int) model.QuadrupletResponse {

	var res model.QuadrupletResponse
	var quadruplets = [][4]int{}

	for i := 0; i < len(primes)-1; i++ {
		if CheckIfPair(primes[i], primes[i+1]) {
			res.PairsCount++
			if i < len(primes)-3 && primes[i+2]-primes[i+1] == 4 && CheckIfPair(primes[i+2], primes[i+3]) {
				quadruplets = append(quadruplets, [4]int{primes[i], primes[i+1], primes[i+2], primes[i+3]})
			}
		}
	}

	res.PrimeCount = len(primes)
	res.Quadruplets = quadruplets
	res.QuadrupletsCount = len(res.Quadruplets)
	return res
}

func CheckIfPair(first int, second int) bool {
	return second-first == 2
}
