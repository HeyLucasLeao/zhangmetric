package config

import (
	"log"

	"github.com/james-bowman/sparse"
	"github.com/nlpodyssey/gopickle/pickle"
	"github.com/sbinet/npyio/npz"
)

func NewCSRMatrix(f *npz.Reader) *sparse.CSR {

	data := read_data_npy(f)
	indices := read_indices_npy(f)
	indptr := read_indptr_npy(f)
	shape := read_shape_npy(f)

	defer f.Close()

	return sparse.NewCSR(shape[0], shape[1], indices, indptr, data)
}

func NewNames() interface{} {
	f, err := pickle.Load("./data/one_hot_columns.pkl")

	if err != nil {
		log.Fatal(err)
	}

	return f
}

func read_data_npy(f *npz.Reader) []float64 {

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

func read_indices_npy(f *npz.Reader) []int {

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

func read_indptr_npy(f *npz.Reader) []int {

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

func read_shape_npy(f *npz.Reader) []int {

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

func SparseColumns() interface{} {
	f, err := pickle.Load("./data/one_hot_columns.pkl")

	if err != nil {
		log.Fatal(err)
	}

	return f

}

func NewSparseMatrix() *npz.Reader {

	f, err := npz.Open("./data/one_hot_values.npz")

	if err != nil {
		log.Fatal(err)
	}

	return f
}
