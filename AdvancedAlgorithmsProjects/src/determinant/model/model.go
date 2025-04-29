package model

type DeterminantRequest struct {
	Size   int         `binding:"required"`
	Matrix [][]float64 `binding:"required"`
}

type DeterminantResponse struct {
	Determinant float64
}
