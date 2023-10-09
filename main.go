package main

import (
	"fmt"
	config "zhang/configuration"
	pipe "zhang/pipeline"
)

func main() {
	//columns := config.SparseColumns()
	f := config.NewReadNpz()
	shape, indptr, indices, data := config.NewReadNpy(f)
	csr_matrix := pipe.NewCSRMatrix(shape[0], shape[1], indptr, indices, data)
	csc_matrix := csr_matrix.ToCSC()
	arr := pipe.NewColumnValues(csc_matrix, 0)
	fmt.Println(arr)
}
