package fasttext

/*
#cgo CFLAGS: -x c++ -std=c++17
#cgo CXXFLAGS: -std=c++17 -pthread -march=native
#cgo LDFLAGS: -pthread
#include <stdlib.h>
#include "fasttext_wrapper.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// FastTextWrapper is a wrapper for the FastText C++ library.
type FastText struct {
	fasttext *C.FastText
}

// NewFastTextWrapper creates a new instance of FastTextWrapper.
func InitFastText() *FastText {
	return &FastText{fasttext: C.createFastText()}
}

func (f *FastText) LoadModel(path string) error {
	modelPathC := C.CString(path)
	defer C.free(unsafe.Pointer(modelPathC))

	if err := C.LoadModel(f.fasttext, modelPathC); err != nil {
		return fmt.Errorf("loading model failed: %v", err)
	}

	return nil
}

func (fw *FastText) Predict(text string) (string, error) {
	textC := C.CString(text)
	defer C.free(unsafe.Pointer(textC))

	prediction := C.fasttextPredict(fw.fasttext, textC)
	defer C.free(unsafe.Pointer(prediction))

	return C.GoString(prediction), nil
}
