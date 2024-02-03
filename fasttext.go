package main

/*
#cgo CFLAGS: -x c++ -std=c++17
#cgo CXXFLAGS: -std=c++17 -pthread -march=native
#cgo LDFLAGS: -L/path/to/fasttext -lfasttext -pthread
#include <stdlib.h>
#include "fasttext_wrapper.h"
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

// FastText is a wrapper for the FastText C++ library.
type FastText struct {
	fasttext *C.FastText
}

// InitFastText creates a new instance of FastText.
func InitFastText() *FastText {
	// return &FastText{fasttext: C.createFastText()}
	return C.createFastText()
}

// LoadModel loads a pre-trained model.
func (f *FastText) LoadModel(path string) error {
	modelPathC := C.CString(path)
	defer C.free(unsafe.Pointer(modelPathC))

	if C.LoadModel(f.fasttext, modelPathC) != 0 {
		return fmt.Errorf("loading model failed")
	}

	return nil
}

// Predict performs prediction on a text.
func (fw *FastText) Predict(text string) (string, error) {
	textC := C.CString(text)
	defer C.free(unsafe.Pointer(textC))

	prediction := C.fasttextPredict(fw.fasttext, textC)
	return C.GoString(prediction), nil
}

func main() {
	// Create FastText instance
	ft := InitFastText()

	// Load pre-trained model
	err := ft.LoadModel("./model.bin")
	if err != nil {
		log.Fatal(err)
	}

	// Perform prediction
	result, err := ft.Predict("Some input text")
	if err != nil {
		log.Fatal(err)
	}

	// Add your assertion based on the expected result
	expectedResult := "Expected result"
	if result != expectedResult {
		fmt.Errorf("Expected %s, but got %s", expectedResult, result)
	}

	// Destroy resources (not implemented in the provided C++ code)
	// ft.Destroy()
}
