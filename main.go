// main.go

package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -L/usr/local/lib -lfasttext -lstdc++ -fPIC
#include "fasttext_wrapper.h"
*/
import "C"

func main() {
	// Create an instance of FastText
	model := C.fasttext_create_model()

	// Perform some operations using the model
	C.fasttext_predict(model, C.CString("your_text_here"), 42)

	// Destroy the model when done
	C.fasttext_destroy_model(model)
}
