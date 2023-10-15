package main

import (
	"fmt"
	"os"
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

	var scorer pipe.ProductScorer

	scorer.Load()

	r := gin.Default()

	r.POST("/score", scorer.PostScore)

	r.Run()

}
