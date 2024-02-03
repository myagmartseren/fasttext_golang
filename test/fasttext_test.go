// main_test.go
package main

import (
	fasttext "fasttext"
	"testing"
)

func TestPrediction(t *testing.T) {
	// Create FastTextWrapper
	ft := fasttext.InitFastText()

	// Load pre-trained model
	ft.LoadModel("./model.bin")

	// Perform prediction
	result, err := ft.Predict("Some input text")
	if err != nil {
		t.Fatal(err)
	}

	// Add your assertion based on the expected result
	expectedResult := "Expected result"
	if result != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, result)
	}

	// Destroy resources
	// ft.Destroy()
}
