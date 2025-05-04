package monte_carlo_model

type MonteCarloRequest struct {
	Iterations int
}

type MonteCarloResponse struct {
	Estimate float64
	Percent  float64
}
