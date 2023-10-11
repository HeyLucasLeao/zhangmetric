package main

import (
	"fmt"
	"net/http"
	"os"
	config "zhang/configuration"
	pipe "zhang/pipeline"

	"github.com/gin-gonic/gin"
	"github.com/james-bowman/sparse"
	"github.com/joho/godotenv"
)

func getProducts(ctx *gin.Context, csc_matrix *sparse.CSC, products *map[interface{}]int) {
	var request_products []string
	err := ctx.ShouldBindJSON(&request_products)

	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
	}

	mat := pipe.NewScore(csc_matrix, products, request_products)

	ctx.JSON(http.StatusOK, mat)
}

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

	r := gin.Default()

	r.GET("/score", func(ctx *gin.Context) { getProducts(ctx, csc_matrix, pickle_products) })

	r.Run()

}
