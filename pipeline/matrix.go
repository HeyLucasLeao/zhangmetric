package pipe

import (
	"github.com/james-bowman/sparse"
)

func NewCSRMatrix(r int, c int, indptr []int, indices []int, data []float64) *sparse.CSR {
	return sparse.NewCSR(r, c, indptr, indices, data)
}

func NewColumnValues(mat *sparse.CSC, idx int) *[]float64 {
	arr := make([]float64, mat.RawMatrix().J)
	sl := mat.RawMatrix().Ind[mat.RawMatrix().Indptr[idx]:mat.RawMatrix().Indptr[idx+1]]

	for _, j := range sl {
		arr[j] = 1
	}
	return &arr
}
