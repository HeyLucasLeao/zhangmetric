package main

import (
	config "zhang/configuration"
	pipe "zhang/pipeline"
)

func main() {
	//columns := config.SparseColumns()
	f := config.NewReadNpz()
	shape, indptr, indices, data := config.NewReadNpy(f)
	csc_matrix := pipe.NewCSCMatrix(shape[0], shape[1], indptr, indices, data)
	pipe.NewColumnValues(csc_matrix)
}
