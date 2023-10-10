package pipe

import (
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

func zhangMetric(antecedent, consequent *[]float64) float64 {
	supportA := mean(antecedent)
	supportC := mean(consequent)

	supportAC := logicalAndMean(antecedent, consequent)

	numerator := supportAC - (supportA * supportC)
	denominator := math.Max(supportAC*(1-supportA), supportA*(supportC-supportAC))

	if denominator == 0 {
		return -1.0
	}

	return numerator / denominator
}

func calculateMetric(csc_matrix *sparse.CSC, products *map[interface{}]int, antecedent string, consequent string) float64 {

	idxA := NewProductIndex(products, antecedent)
	matA := NewProductMatrix(csc_matrix, idxA)

	idxC := NewProductIndex(products, consequent)
	matC := NewProductMatrix(csc_matrix, idxC)

	return zhangMetric(matA, matC)
}

func GenerateResponse(csc_matrix *sparse.CSC, products *map[interface{}]int, arr []string) *map[string]map[string]float64 {
	iterator := iterium.Combinations(arr, 2)
	comb, err := iterator.Slice()

	if err != nil {
		log.Fatalf("%s", err)
	}

	matrix := make(map[string]map[string]float64)

	for _, array := range comb {
		matrix[array[0]] = make(map[string]float64)
		matrix[array[0]][array[1]] = calculateMetric(csc_matrix, products, array[0], array[1])
	}

	return &matrix
}
