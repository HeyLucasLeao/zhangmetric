package main

import (
	"fmt"
	config "zhang/configuration"
)

func main() {
	//columns := config.SparseColumns()
	csr_matrix := config.SparseMatrix()
	matrix := config.NewCSRMatrix(csr_matrix)
	fmt.Println(matrix.Data)
}
