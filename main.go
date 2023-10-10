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
	arr := []string{"10101012", "1900622023", "-1", "74403460", "31005497"}
	response := pipe.GenerateResponse(csc_matrix, products, arr)

	for key1, innerMap := range *response {
		for key2, value := range innerMap {
			fmt.Printf("Itens: %s & %s - Score %2f\n", key1, key2, value)
		}
	}

}
