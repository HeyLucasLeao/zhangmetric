package pipe

import (
	"log"
	"math"
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

func ZhangMetric(antecedent, consequent *[]float64) float64 {
	supportA := mean(antecedent)
	supportC := mean(consequent)

	supportAC := logicalAndMean(antecedent, consequent)

	numerator := supportAC - (supportA * supportC)
	denominator := math.Max(supportAC*(1-supportA), supportA*(supportC-supportAC))

	return numerator / denominator
}
