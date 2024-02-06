package fasttext

/*
#cgo LDFLAGS: -lfasttext -pthread
#include <stdlib.h>
#include <math.h>
#include "wrapper.h"
*/
import "C"
import (
	"unsafe"
)

func init() {

}

// FastText is a wrapper for the FastText C++ library.
type FastText struct {
	fasttext *C.fasttext_t
}

// InitFastText creates a new instance of FastText.
func NewFasttext() *FastText {
	ft := new(FastText)
	ft.fasttext = C.fasttext_new()
	return ft
}

// func (ft *FastText) Delete() {
// 	C.fasttext_delete(unsafe.Fasttext(ft.fasttext))
// }

// LoadModel loads a pre-trained model.
func (f *FastText) LoadModel(path string) error {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))

	_, err := C.fasttext_load_model(f.fasttext, pathC)
	return err
}

func (f *FastText) Predict(text string) (string, error) {
	textC := C.CString(text)
	defer C.free(unsafe.Pointer(textC))

	resultC, err := C.fasttext_predict(f.fasttext, textC)
	if err != nil {
		return "", err
	}
	return C.GoString(resultC), nil
}
