package main

import (
	"fmt"
	config "zhang/configuration"
	pipe "zhang/pipeline"
)

func main() {
	f := config.NewReadNpz()
	shape, indptr, indices, data := config.NewReadNpy(f)
	csr_matrix := pipe.NewCSRMatrix(shape[0], shape[1], indptr, indices, data)
	csc_matrix := csr_matrix.ToCSC()
	products := config.ReadProducts()
	idx := pipe.NewProductIndex(products, "3")
	mat := pipe.NewProductMatrix(csc_matrix, idx)
	fmt.Println(mat)
}
