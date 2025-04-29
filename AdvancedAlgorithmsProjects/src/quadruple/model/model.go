package model

type QuadrupletRequest struct {
	Start int
	End   int
}

type QuadrupletResponse struct {
	PrimeCount       int
	PairsCount       int
	QuadrupletsCount int
	Quadruplets      [][4]int
}
