package main

import (
	"fmt"
	config "zhang/configuration"

	"github.com/james-bowman/sparse"
)

func main() {
	//columns := config.SparseColumns()
	f := config.SparseMatrix()
	csr_matrix := config.NewCSRMatrix(f)
	csc_matrix := csr_matrix.T().(*sparse.CSC)
	fmt.Println(csc_matrix)
}
