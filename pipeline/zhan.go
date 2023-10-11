package pipe

import (
	"fmt"
	"log"
	"math"

	"github.com/james-bowman/sparse"
	"github.com/mowshon/iterium"
)

func mean(mat *[]float64) float64 {
	sum := 0.0
	total := len(*mat)
	for _, num := range *mat {
		sum += num
	}

	return sum / float64(total)
}

func logicalAndMean(a, b *[]float64) float64 {
	if len(*a) != len(*b) {
		log.Fatalf("Error: slices are not the same length")
	}

	count := 0.0
	for i := range *a {
		if (*a)[i]+(*b)[i] == 2 {
			count++
		}
	}

	return count / float64(len(*a))
}

func zhangMetric(antecedent, consequent *[]float64) string {
	supportA := mean(antecedent)
	supportC := mean(consequent)

	supportAC := logicalAndMean(antecedent, consequent)

	numerator := supportAC - (supportA * supportC)
	denominator := math.Max(supportAC*(1-supportA), supportA*(supportC-supportAC))

	if denominator == 0 {
		return "Itens sem registros juntos!"
	}

	return fmt.Sprintf("%f", numerator/denominator)
}

func calculateMetric(csc_matrix *sparse.CSC, products *map[interface{}]int) func(antecedent string, consequent string) string {
	return func(antecedent string, consequent string) string {
		idxA := NewProductIndex(products, antecedent)
		matA := NewProductMatrix(csc_matrix, idxA)

		idxC := NewProductIndex(products, consequent)
		matC := NewProductMatrix(csc_matrix, idxC)

		return zhangMetric(matA, matC)
	}
}

func NewScore(csc_matrix *sparse.CSC, products *map[interface{}]int, arr []string) *map[string]string {
	iterator := iterium.Combinations(arr, 2)
	combination, err := iterator.Slice()
	iterator = iterium.New(combination...)

	if err != nil {
		log.Fatalf("%s", err)
	}

	partial := calculateMetric(csc_matrix, products)

	starmap := iterium.StarMap(iterator, partial)

	score, err := starmap.Slice()

	if err != nil {
		log.Fatalf("%s", err)
	}

	mat := make(map[string]string)
	for i, key := range combination {
		keyStr := key[0] + ",  " + key[1]
		mat[keyStr] = score[i]
	}

	return &mat
}
