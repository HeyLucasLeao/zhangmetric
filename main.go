package main

import (
	"fmt"
	"os"
	config "zhang/configuration"
	pipe "zhang/pipeline"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading .env file")
		return
	}

	f := config.NewReadNpz()
	shape, indptr, indices, data := config.NewReadNpy(f)
	products := config.NewReadPickle()
	csr_matrix := pipe.NewCSRMatrix(shape[0], shape[1], indptr, indices, data)
	csc_matrix := csr_matrix.ToCSC()
	idx := pipe.NewProductIndex(products, "10101012")
	mat1 := pipe.NewProductMatrix(csc_matrix, idx)
	idx2 := pipe.NewProductIndex(products, "3")
	mat2 := pipe.NewProductMatrix(csc_matrix, idx2)
	fmt.Println(pipe.ZhangMetric(mat1, mat2))
}
