package pipe

import (
	"log"
	"math"
	"net/http"
	config "zhang/configuration"

	"github.com/gin-gonic/gin"
	"github.com/james-bowman/sparse"
	"github.com/mowshon/iterium"
)

type ProductScorer struct {
	Matrix   *sparse.CSC
	Products *map[interface{}]int
}

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

func (s ProductScorer) newProductIndex(product_name string) int {
	// Check if a key exists
	idx, ok := (*s.Products)[product_name]

	if !ok {
		return -1
	}

	return idx
}

func (s ProductScorer) newProductMatrix(idx int) *[]float64 {

	mat := s.Matrix.RawMatrix()
	arr := make([]float64, mat.J)

	if idx == -1 {
		return &arr
	}

	sl := mat.Ind[mat.Indptr[idx]:mat.Indptr[idx+1]]

	for _, j := range sl {
		arr[j] = 1
	}

	return &arr
}

func (s ProductScorer) calculateMetric(antecedent string, consequent string) *float64 {

	idxA := s.newProductIndex(antecedent)
	matA := s.newProductMatrix(idxA)

	idxC := s.newProductIndex(consequent)
	matC := s.newProductMatrix(idxC)

	return zhangMetric(matA, matC)
}

func (s ProductScorer) newScore(arr []string) *map[string]*float64 {
	iterator := iterium.Combinations(arr, 2)
	combination, err := iterator.Slice()

	if err != nil {
		log.Fatalf("%s", err)
	}

	iterator = iterium.New(combination...)
	starmap := starMap(iterator, s.calculateMetric)

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

func (s *ProductScorer) LoadFiles() {

	f := config.NewReadNpz()
	shape, indptr, indices, data := config.NewReadNpy(f)
	mat := sparse.NewCSR(shape[0], shape[1], indptr, indices, data).ToCSC()

	products := config.NewReadPickle()

	s.Matrix = mat
	s.Products = products
}

func (s ProductScorer) PostScore(ctx *gin.Context) {
	var request_products []string
	err := ctx.ShouldBindJSON(&request_products)

	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
	}

	mat := s.newScore(request_products)

	ctx.JSON(http.StatusOK, mat)
}
