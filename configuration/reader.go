package config

import (
	"fmt"
	"log"
	"os"

	"github.com/nlpodyssey/gopickle/pickle"
	"github.com/nlpodyssey/gopickle/types"
	"github.com/sbinet/npyio/npz"
)

// Read NPZ file in ./data folder
func NewReadNpz() *npz.Reader {
	filepath := fmt.Sprintf("./data/%s", os.Getenv("NPZ_FILENAME"))

	f, err := npz.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	return f
}

// Read Npy files inside NPZ File
func NewReadNpy(f *npz.Reader) ([]int, []int, []int, []float64) {

	shape := readShapeNpy(f)
	indptr := readIntptrNpy(f)
	indices := readIndicesNpy(f)
	data := readDataNpy(f)

	defer f.Close()

	return shape, indptr, indices, data
}

// Read []string Python object pickle stored
func NewReadPickle() *map[interface{}]int {
	filepath := fmt.Sprintf("./data/%s", os.Getenv("PRODUCT_FILENAME"))
	f, err := pickle.Load(filepath)

	if err != nil {
		log.Fatal(err)
	}

	slice, ok := f.(*types.List)

	if !ok {
		log.Fatal("Could not convert file into *types.List")
	}

	mat := make(map[interface{}]int)
	for i := 0; i < slice.Len(); i++ {
		mat[slice.Get(i)] = i
	}

	return &mat
}

// Read data.npy from NPZ File and convert as []float64
func readDataNpy(f *npz.Reader) []float64 {

	var data []bool
	err := f.Read("data.npy", &data)
	if err != nil {
		log.Fatalf("could not read value from npz file: %+v", err)
	}
	m := map[bool]float64{true: 1, false: 0}

	var result []float64
	for _, b := range data {
		f := m[b]
		result = append(result, f)
	}

	return result
}

// Read indices.npy from NPZ File and convert as []int
func readIndicesNpy(f *npz.Reader) []int {

	var indices []int32
	err := f.Read("indices.npy", &indices)
	if err != nil {
		log.Fatalf("could not read value from npz file: %+v", err)
	}

	var result []int
	for _, v := range indices {
		result = append(result, int(v))
	}

	return result
}

// Read indptr.npy from NPZ File and return as []int
func readIntptrNpy(f *npz.Reader) []int {

	var indptr []int32
	err := f.Read("indptr.npy", &indptr)
	if err != nil {
		log.Fatalf("could not read value from npz file: %+v", err)
	}

	var result []int
	for _, v := range indptr {
		result = append(result, int(v))
	}

	return result
}

// Read shape.npy from NPZ File and return as []int
func readShapeNpy(f *npz.Reader) []int {

	var shape []int64
	err := f.Read("shape.npy", &shape)
	if err != nil {
		log.Fatalf("could not read value from npz file: %+v", err)
	}

	var result []int
	for _, v := range shape {
		result = append(result, int(v))
	}

	return result
}
