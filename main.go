package main

import (
	"fmt"
	"os"
	config "zhang/configuration"
	pipe "zhang/pipeline"

	"github.com/gin-gonic/gin"
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
	pickle_products := config.NewReadPickle()
	csc_matrix := pipe.NewCSRMatrix(shape[0], shape[1], indptr, indices, data).ToCSC()
	scorer := pipe.ProductScorer{Matrix: csc_matrix, Products: pickle_products}

	r := gin.Default()

	r.POST("/score", scorer.PostScore)

	r.Run()

}
