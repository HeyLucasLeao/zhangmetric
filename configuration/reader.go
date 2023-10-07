package config

import (
	"log"

	"github.com/nlpodyssey/gopickle/pickle"
	"github.com/sbinet/npyio/npz"
)

type CSRMatrix struct {
	Data    []bool
	Indices []int32
	Indptr  []int32
}

func NewCSRMatrix(f *npz.Reader) *CSRMatrix {
	var data []bool
	var indices []int32
	var indptr []int32

	err := f.Read("data.npy", &data)
	if err != nil {
		log.Fatalf("could not read value from npz file: %+v", err)
	}

	err = f.Read("indices.npy", &indices)
	if err != nil {
		log.Fatalf("could not read value from npz file: %+v", err)
	}

	err = f.Read("indptr.npy", &indptr)
	if err != nil {
		log.Fatalf("could not read value from npz file: %+v", err)
	}

	return &CSRMatrix{Data: data, Indices: indices, Indptr: indptr}
}

func SparseColumns() interface{} {
	f, err := pickle.Load("./data/one_hot_columns.pkl")

	if err != nil {
		log.Fatal(err)
	}

	return f

}

func SparseMatrix() *npz.Reader {

	f, err := npz.Open("./data/one_hot_values.npz")

	if err != nil {
		log.Fatal(err)
	}

	return f
}
