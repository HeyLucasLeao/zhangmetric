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

func zhangMetric(antecedent, consequent *[]float64) *float64 {
	supportA := mean(antecedent)
	supportC := mean(consequent)

	supportAC := logicalAndMean(antecedent, consequent)

	numerator := supportAC - (supportA * supportC)
	denominator := math.Max(supportAC*(1-supportA), supportA*(supportC-supportAC))

	if denominator == 0 {
		return nil
	}

	score := numerator / denominator

	return &score
}

func starMap[T any](iterable iterium.Iter[[]T], apply func(T, T) *float64) iterium.Iter[*float64] {
	iter := iterium.Instance[*float64](iterable.Count(), iterable.IsInfinite())

	go func() {
		defer iterium.IterRecover()
		defer iter.Close()

		for {
			next, err := iterable.Next()
			if err != nil {
				return
			}

			// Apply the function to the values from the slide.
			iter.Chan() <- apply(next[0], next[1])
		}
	}()

	return iter
}

func calculateMetric(csc_matrix *sparse.CSC, products *map[interface{}]int) func(antecedent string, consequent string) *float64 {
	return func(antecedent string, consequent string) *float64 {
		idxA := NewProductIndex(products, antecedent)
		matA := NewProductMatrix(csc_matrix, idxA)

		idxC := NewProductIndex(products, consequent)
		matC := NewProductMatrix(csc_matrix, idxC)

		return zhangMetric(matA, matC)
	}
}

func NewScore(csc_matrix *sparse.CSC, products *map[interface{}]int, arr []string) *map[string]*float64 {
	iterator := iterium.Combinations(arr, 2)
	combination, err := iterator.Slice()

	if err != nil {
		log.Fatalf("%s", err)
	}

	iterator = iterium.New(combination...)
	partial := calculateMetric(csc_matrix, products)
	starmap := starMap(iterator, partial)

	score, err := starmap.Slice()

	if err != nil {
		log.Fatalf("%s", err)
	}

	mat := make(map[string]*float64)
	for i, key := range combination {
		keyStr := key[0] + ",  " + key[1]
		mat[keyStr] = score[i]
	}

	return &mat
}
