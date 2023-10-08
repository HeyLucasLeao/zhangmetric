package pipe

import (
	"github.com/james-bowman/sparse"
)

func NewCSCMatrix(r int, c int, indptr []int, indices []int, data []float64) *sparse.CSC {
	return sparse.NewCSC(r, c, indptr, indices, data)
}

func NewColumnValues(m *sparse.CSC) []float64 {
	arr := make([]float64, m.RawMatrix().J)
	for i := 0; i < m.RawMatrix().I; i++ {
		row := m.At(i, 0)
		if row > 0 {
			arr[i] = row
		}
	}
	return arr
}
