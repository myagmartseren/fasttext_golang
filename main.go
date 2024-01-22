package main

/*
#cgo CFLAGS: -x c++ -std=c++11
#cgo CXXFLAGS: -march=native -Wall -I/usr/local/include/fastText
#cgo LDFLAGS: -L/usr/local/lib -lfasttext
#include "fasttext_wrapper.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// FastTextModel represents the Go wrapper for the C++ FastText model.
type FastTextModel struct {
	model *C.FastTextModel
}

// LoadModel loads the FastText model from the specified path.
func LoadModel(modelPath string) *FastTextModel {
	cModelPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cModelPath))

	model := C.fasttext_load_model(cModelPath)
	if model == nil {
		return nil
	}

	return &FastTextModel{model: model}
}

// GetDimension returns the dimension of the FastText model.
func (ft *FastTextModel) GetDimension() int {
	return int(C.fasttext_getDimension(ft.model))
}

// GetWordVector returns the word vector for the specified word.
func (ft *FastTextModel) GetWordVector(word string) []float32 {
	cWord := C.CString(word)
	defer C.free(unsafe.Pointer(cWord))

	dim := ft.GetDimension()
	vec := make([]float32, dim)

	C.fasttext_getWordVector(ft.model, cWord, (*C.float)(&vec[0]))

	return vec
}

func main() {
	// Load a pre-trained FastText model
	modelPath := "./model.bin"
	model := LoadModel(modelPath)
	if model == nil {
		panic("Failed to load FastText model")
	}
	defer C.free(unsafe.Pointer(model)) // Ensure model is freed

	// Get model dimension
	dimension := model.GetDimension()
	fmt.Printf("Model Dimension: %d\n", dimension)

	// Get word vector for a specific word
	word := "your_text_here"
	wordVector := model.GetWordVector(word)
	fmt.Printf("Word Vector for '%s': %v\n", word, wordVector)
}
