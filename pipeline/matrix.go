package pipe

import (
	"github.com/james-bowman/sparse"
)

func NewCSRMatrix(r int, c int, indptr []int, indices []int, data []float64) *sparse.CSR {
	return sparse.NewCSR(r, c, indptr, indices, data)
}

func (s ProductScorer) NewProductMatrix(idx int) *[]float64 {

	arr := make([]float64, s.Matrix.RawMatrix().J)

	if idx == -1 {
		return &arr
	}

	sl := s.Matrix.RawMatrix().Ind[s.Matrix.RawMatrix().Indptr[idx]:s.Matrix.RawMatrix().Indptr[idx+1]]

	for _, j := range sl {
		arr[j] = 1
	}

	return &arr
}

func (s ProductScorer) NewProductIndex(product_name string) int {
	// Check if a key exists
	idx, ok := (*s.Products)[product_name]

	if !ok {
		return -1
	}

	return idx
}
