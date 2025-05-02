package bisection_model

type BisectionRequest struct {
	Numbers []float64
	Delta   float64
	Min     float64
	Max     float64
}

type BisectionResponse struct {
	Number []float64
}
